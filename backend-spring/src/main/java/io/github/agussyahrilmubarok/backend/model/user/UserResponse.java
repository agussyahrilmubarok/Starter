package io.github.agussyahrilmubarok.backend.model.user;

import java.time.LocalDateTime;

import com.fasterxml.jackson.annotation.JsonProperty;

public record UserResponse(
        @JsonProperty("id") String id,
        @JsonProperty("name") String name,
        @JsonProperty("email") String email,
        @JsonProperty("created_at") LocalDateTime createdAt,
        @JsonProperty("udpated_at") LocalDateTime updatedAt) {
}
