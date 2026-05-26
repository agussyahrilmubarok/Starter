package io.github.agussyahrilmubarok.backend.model.user;

import jakarta.validation.constraints.Email;
import jakarta.validation.constraints.Size;

public record UpdateUserRequest(

        String name,

        @Email(message = "Email format is invalid")
        String email,

        String password
) {
}