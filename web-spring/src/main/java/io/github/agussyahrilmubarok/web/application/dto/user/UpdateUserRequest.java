package io.github.agussyahrilmubarok.web.application.dto.user;

import jakarta.validation.constraints.Email;

public record UpdateUserRequest(
        String name,

        @Email(message = "Email is not valid")
        String email,

        String password
) {
}
