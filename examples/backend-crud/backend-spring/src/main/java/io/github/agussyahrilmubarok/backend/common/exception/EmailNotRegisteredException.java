package io.github.agussyahrilmubarok.backend.common.exception;

import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.ResponseStatus;

@ResponseStatus(HttpStatus.UNAUTHORIZED)
public class EmailNotRegisteredException extends RuntimeException {
    public EmailNotRegisteredException() {
        super("Email is not registered");
    }
}
