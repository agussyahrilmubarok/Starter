package io.github.agussyahrilmubarok.backend.common.exception;

import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.ResponseStatus;

@ResponseStatus(HttpStatus.UNAUTHORIZED)
public class JwtInvalidTokenException extends RuntimeException {
    public JwtInvalidTokenException() {
        super("Invalid token.");
    }
}
