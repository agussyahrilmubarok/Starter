package io.github.agussyahrilmubarok.backend.application.dto.user;

import com.fasterxml.jackson.annotation.JsonProperty;
import jakarta.validation.constraints.Email;

public record UpdateUserRequest(
        @JsonProperty("name") String name,

        @JsonProperty("email") @Email(message = "Email is not valid")
        String email,

        @JsonProperty("password") String password) {}
