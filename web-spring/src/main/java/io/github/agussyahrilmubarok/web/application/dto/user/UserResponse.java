package io.github.agussyahrilmubarok.web.application.dto.user;

import java.time.OffsetDateTime;
import java.util.UUID;

public record UserResponse(
        UUID id,
        String name,
        String email,
        OffsetDateTime createdAt,
        OffsetDateTime updatedAt
) {
}