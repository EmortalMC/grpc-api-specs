package dev.emortal.api.service.gameplayerdata;

import com.google.protobuf.Message;
import dev.emortal.api.model.gamedata.GameDataGameMode;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.util.List;
import java.util.Map;
import java.util.UUID;

public interface GamePlayerDataService {

    <T extends Message> @Nullable T getGameData(@NotNull GameDataGameMode gameMode, @NotNull Class<T> clazz, @NotNull UUID playerId);

    <T extends Message> @NotNull Map<UUID, @Nullable T> getGameData(@NotNull GameDataGameMode gameMode, @NotNull Class<T> clazz, @NotNull Iterable<@NotNull UUID> playerIds);
}
