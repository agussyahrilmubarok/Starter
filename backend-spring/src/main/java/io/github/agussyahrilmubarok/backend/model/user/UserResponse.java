package io.github.agussyahrilmubarok.backend.model.user;

import com.fasterxml.jackson.annotation.JsonProperty;

import java.time.LocalDateTime;

public record UserResponse(
        @JsonProperty("id") String id,
        @JsonProperty("name") String name,
        @JsonProperty("email") String email,
        @JsonProperty("created_at") LocalDateTime createdAt,
        @JsonProperty("udpated_at") LocalDateTime updatedAt) {
}
