package io.github.agussyahrilmubarok.web.application.dto.user;

import jakarta.validation.constraints.Email;
import jakarta.validation.constraints.Size;
import java.util.UUID;

public record UpdateUserRequest(
        UUID id,

        @Size(max = 100, message = "{user.name.size}") 
        String name,

        @Email(message = "{user.email.invalid}") 
        @Size(max = 150, message = "{user.email.size}") 
        String email,

        @Size(min = 8, max = 72, message = "{user.password.size}") 
        String password) {}
