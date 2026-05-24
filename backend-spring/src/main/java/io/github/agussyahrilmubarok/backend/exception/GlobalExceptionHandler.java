package io.github.agussyahrilmubarok.backend.exception;

import io.github.agussyahrilmubarok.backend.model.ApiResponse;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.MethodArgumentNotValidException;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.RestControllerAdvice;

import java.util.HashMap;
import java.util.Map;

@RestControllerAdvice
public class GlobalExceptionHandler {

    @ExceptionHandler(MethodArgumentNotValidException.class)
    public ResponseEntity<ApiResponse<Void>> handleValidationException(MethodArgumentNotValidException ex) {
        Map<String, String> errors = new HashMap<>();

        ex.getBindingResult().getFieldErrors().forEach(fieldError -> {
            String field = fieldError.getField();
            String message = switch (fieldError.getCode()) {
                case "NotBlank", "NotNull", "NotEmpty" -> field + " is required";
                case "Email" -> "Invalid email format";
                case "Size" -> {
                    Object[] args = fieldError.getArguments();
                    int max = (int) args[1];
                    int min = (int) args[2];
                    if (min > 0 && max < Integer.MAX_VALUE)
                        yield field + " must be between " + min + " and " + max + " characters";
                    else if (min > 0)
                        yield field + " must be at least " + min + " characters";
                    else
                        yield field + " must be at most " + max + " characters";
                }
                case "Min" -> field + " must be at least " + fieldError.getArguments()[1];
                case "Max" -> field + " must be at most " + fieldError.getArguments()[1];
                case "Pattern" -> "Invalid " + field + " format";
                default -> "Invalid value";
            };
            errors.put(field, message);
        });

        return ResponseEntity
                .status(HttpStatus.BAD_REQUEST)
                .body(new ApiResponse<>("validation failed", errors));
    }
}
