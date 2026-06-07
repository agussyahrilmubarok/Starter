package io.github.agussyahrilmubarok.web.application.dto.user;

import jakarta.validation.constraints.Email;
import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.Size;

public record CreateUserRequest(
        @NotBlank(message = "{user.name.notBlank}")
        @Size(min = 2, max = 100, message = "{user.name.size}")
        String name,

        @NotBlank(message = "{user.email.notBlank}")
        @Email(message = "{user.email.invalid}")
        @Size(max = 150, message = "{user.email.size}")
        String email,

        @NotBlank(message = "{user.password.notBlank}")
        @Size(min = 8, max = 72, message = "{user.password.size}")
        String password) {}
