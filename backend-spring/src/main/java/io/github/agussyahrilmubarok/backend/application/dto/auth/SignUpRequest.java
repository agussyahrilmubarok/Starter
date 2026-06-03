package io.github.agussyahrilmubarok.backend.application.dto.auth;

import com.fasterxml.jackson.annotation.JsonProperty;
import jakarta.validation.constraints.Email;
import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.Size;

public record SignUpRequest(
        @JsonProperty("name")
        @NotBlank(message = "Name is required")
        @Size(min = 2, max = 100, message = "Name must be between 2 and 100 characters")
        String name,

        @JsonProperty("email")
        @NotBlank(message = "Email is required")
        @Email(message = "Email is not valid")
        @Size(max = 150, message = "Email must not exceed 150 characters")
        String email,

        @JsonProperty("password")
        @NotBlank(message = "Password is required")
        @Size(min = 8, max = 72, message = "Password must be between 8 and 72 characters")
        String password
) {
}