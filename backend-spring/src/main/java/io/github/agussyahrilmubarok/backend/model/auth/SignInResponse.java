package io.github.agussyahrilmubarok.backend.model.auth;

import com.fasterxml.jackson.annotation.JsonProperty;
import io.github.agussyahrilmubarok.backend.model.user.UserResponse;

public record SignInResponse(
        @JsonProperty("token") String token,
        @JsonProperty("user") UserResponse userResponse) {
}
