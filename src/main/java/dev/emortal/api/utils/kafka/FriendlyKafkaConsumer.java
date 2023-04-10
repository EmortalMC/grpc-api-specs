package dev.emortal.api.utils.kafka;

import com.google.common.util.concurrent.ThreadFactoryBuilder;
import com.google.protobuf.AbstractMessage;
import com.google.protobuf.InvalidProtocolBufferException;
import dev.emortal.api.utils.parser.MessageProtoConfig;
import dev.emortal.api.utils.parser.MessagingService;
import dev.emortal.api.utils.parser.ProtoParserRegistry;
import org.apache.kafka.clients.consumer.ConsumerRecord;
import org.apache.kafka.clients.consumer.ConsumerRecords;
import org.apache.kafka.clients.consumer.KafkaConsumer;
import org.apache.kafka.common.header.Header;
import org.apache.kafka.common.header.Headers;
import org.apache.kafka.common.serialization.ByteArrayDeserializer;
import org.apache.kafka.common.serialization.StringDeserializer;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.time.Duration;
import java.util.Collections;
import java.util.HashSet;
import java.util.Map;
import java.util.Set;
import java.util.UUID;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.Executors;
import java.util.concurrent.ScheduledExecutorService;
import java.util.concurrent.TimeUnit;
import java.util.function.Consumer;

public class FriendlyKafkaConsumer {
    private static final Logger LOGGER = LoggerFactory.getLogger(FriendlyKafkaConsumer.class);

    private final @NotNull KafkaConsumer<String, byte[]> consumer;

    private final Map<Class<?>, Set<Consumer<AbstractMessage>>> protoListeners = new ConcurrentHashMap<>();
    private final @NotNull Set<String> consumedTopics = Collections.synchronizedSet(new HashSet<>());

    private final ScheduledExecutorService scheduler = Executors.newSingleThreadScheduledExecutor(
            new ThreadFactoryBuilder().setNameFormat("kafka-consumer-scheduler")
                    .setUncaughtExceptionHandler((t, e) -> LOGGER.error("Error in Kafka consumer: ", e))
                    .build());

    private final boolean autoCommit;

    public FriendlyKafkaConsumer(@NotNull KafkaSettings settings) {
        // Nullable options
        this.autoCommit = settings.isAutoCommit();

        Map<String, Object> settingsMap = settings.getSettings();
        if (settingsMap.get("group.id") == null) {
            settingsMap.put("group.id", "kafka-core-" + UUID.randomUUID());
        }

        this.consumer = new KafkaConsumer<>(settingsMap, new StringDeserializer(), new ByteArrayDeserializer());

        LOGGER.info("Starting Kafka consumer thread");
        this.scheduler.scheduleAtFixedRate(this::consume, 0, 100, TimeUnit.MILLISECONDS);
    }

    private void consume() {
        synchronized (this.consumedTopics) {
            if (this.consumedTopics.isEmpty()) return;

            ConsumerRecords<String, byte[]> records = this.consumer.poll(Duration.ofMillis(100));

            for (ConsumerRecord<String, byte[]> record : records) {
                String protoType = this.getProtoType(record.headers());
                if (protoType == null) {
                    LOGGER.warn("Received message without X-Proto-Type header");
                    continue;
                }

                AbstractMessage message;
                try {
                    message = ProtoParserRegistry.parse(protoType, record.value());
                } catch (InvalidProtocolBufferException e) {
                    LOGGER.warn("Failed to parse message", e);
                    continue;
                }

                try {
                    Set<Consumer<AbstractMessage>> consumers = this.protoListeners.get(message.getClass());
                    if (consumers != null) {
                        for (Consumer<AbstractMessage> consumer : consumers) {
                            consumer.accept(message);
                        }
                    }
                } catch (Exception e) {
                    LOGGER.error("Failed to handle message (topic: {}, type: {}): {}", record.topic(), protoType, e);
                }
            }

            if (this.autoCommit) {
                this.consumer.commitSync();
            }
        }
    }

    private @Nullable String getProtoType(@NotNull Headers headers) {
        for (Header header : headers) {
            if (header.key().equals("X-Proto-Type")) {
                return new String(header.value());
            }
        }
        return null;
    }

    @SuppressWarnings("unchecked")
    public <T extends AbstractMessage> void addListener(@NotNull Class<T> messageType, @NotNull Consumer<T> listener) {
        synchronized (this.consumedTopics) {
            MessageProtoConfig<T> protoConfig = ProtoParserRegistry.getParser(messageType);
            if (protoConfig == null) {
                throw new IllegalArgumentException("No parser found for " + messageType.getName());
            }

            if (protoConfig.service() != MessagingService.KAFKA) {
                throw new IllegalArgumentException("Parser for " + messageType.getName() + " is not for Kafka");
            }

            if (!this.consumedTopics.contains(protoConfig.topic())) {
                LOGGER.debug("Subscribing to topic {}", protoConfig.topic());
                this.consumedTopics.add(protoConfig.topic());
                this.consumer.subscribe(this.consumedTopics);
            }

            this.protoListeners.computeIfAbsent(messageType, k -> new HashSet<>()).add((Consumer<AbstractMessage>) listener);
        }
    }

    public void close() {
        synchronized (this.consumedTopics) {
            this.scheduler.shutdown();
            this.consumer.close();
        }
    }
}