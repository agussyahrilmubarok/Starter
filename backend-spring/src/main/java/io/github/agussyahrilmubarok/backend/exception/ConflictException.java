package io.github.agussyahrilmubarok.backend.exception;

import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.ResponseStatus;

@ResponseStatus(HttpStatus.CONFLICT)
public class ConflictException extends RuntimeException {

    private String field;

    public ConflictException() {
        super();
    }

    public ConflictException(final String message) {
        super(message);
    }

    public ConflictException(final String field, final String message) {
        super(message);
        this.field = field;
    }

    public String getField() {
        return field;
    }
}