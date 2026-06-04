package io.github.agussyahrilmubarok.web.application.dto.user;

import java.time.LocalDateTime;
import java.util.UUID;

public record UserResponse(
        UUID id,
        String name,
        String email,
        LocalDateTime createdAt,
        LocalDateTime updatedAt
) {
}
