package io.github.agussyahrilmubarok.web.model.user;

import jakarta.validation.constraints.Email;
import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class UpdateUserRequest {

    private String name;

    @Email(message = "Email format is invalid")
    private String email;

    private String password;
}
