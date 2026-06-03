package io.github.agussyahrilmubarok.backend.application.dto.auth;

import com.fasterxml.jackson.annotation.JsonProperty;
import io.github.agussyahrilmubarok.backend.application.dto.user.UserResponse;

public record AuthResponse(
        @JsonProperty("token")
        String token,

        @JsonProperty("user")
        UserResponse user
) {
}