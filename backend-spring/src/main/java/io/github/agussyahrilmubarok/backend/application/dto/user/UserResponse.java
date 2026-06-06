package io.github.agussyahrilmubarok.backend.application.dto.user;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.time.OffsetDateTime;
import java.util.UUID;

public record UserResponse(
        @JsonProperty("id") UUID id,

        @JsonProperty("name") String name,

        @JsonProperty("email") String email,

        @JsonProperty("created_at") OffsetDateTime createdAt,

        @JsonProperty("updated_at") OffsetDateTime updatedAt) {}
