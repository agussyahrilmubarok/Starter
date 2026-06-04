package io.github.agussyahrilmubarok.web.application.dto.user;

import jakarta.validation.constraints.Email;

import java.util.UUID;

public record UpdateUserRequest(
        UUID id,
        
        String name,

        @Email(message = "Email is not valid")
        String email,

        String password
) {
}
