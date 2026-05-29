package io.github.agussyahrilmubarok.web.util;

import lombok.Getter;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.ResponseStatus;

@Getter
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
}
