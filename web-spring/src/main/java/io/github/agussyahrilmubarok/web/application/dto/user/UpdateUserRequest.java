package io.github.agussyahrilmubarok.web.application.dto.user;

import jakarta.validation.constraints.Email;
import jakarta.validation.constraints.Size;

import java.util.UUID;

public record UpdateUserRequest(
        UUID id,

        @Size(max = 100)
        String name,

        @Email
        @Size(max = 150)
        String email,

        @Size(min = 8, max = 72)
        String password
) {
}