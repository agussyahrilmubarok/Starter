package io.github.agussyahrilmubarok.web.application.dto.user;

import jakarta.validation.constraints.Email;
import jakarta.validation.constraints.Size;
import java.util.UUID;

public record UpdateUserRequest(
        UUID id,

        @Size(max = 100, message = "Name must be between 2 and 100 characters")
        String name,

        @Email(message = "Email format is invalid") @Size(max = 150, message = "Email must not exceed 150 characters")
        String email,

        @Size(min = 8, max = 72, message = "Password must be between 8 and 72 characters")
        String password) {}
