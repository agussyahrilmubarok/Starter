package io.github.agussyahrilmubarok.backend.exception;

import io.github.agussyahrilmubarok.backend.model.ApiResponse;
import jakarta.servlet.http.HttpServletRequest;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.http.converter.HttpMessageNotReadableException;
import org.springframework.web.bind.MethodArgumentNotValidException;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.RestControllerAdvice;

import java.util.HashMap;
import java.util.Map;

@RestControllerAdvice
public class GlobalExceptionHandler {

    @ExceptionHandler(MethodArgumentNotValidException.class)
    public ResponseEntity<ApiResponse<Void>> handleValidationException(
            MethodArgumentNotValidException ex) {

        Map<String, String> errors = new HashMap<>();
        ex.getBindingResult().getFieldErrors().forEach(fieldError -> {
            String field = capitalize(fieldError.getField());
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
                .status(HttpStatus.UNPROCESSABLE_CONTENT)
                .body(new ApiResponse<>("Validation Failed", errors));
    }

    @ExceptionHandler(NotFoundException.class)
    public ResponseEntity<ApiResponse<Void>> handleNotFoundException(
            NotFoundException ex,
            HttpServletRequest request) {

        Map<String, String> errors = new HashMap<>();

        if (ex.getField() != null) {
            errors.put(ex.getField(), ex.getMessage());
        } else {
            errors.put("error", ex.getMessage());
        }

        return ResponseEntity
                .status(HttpStatus.NOT_FOUND)
                .body(new ApiResponse<>("Not Found Error", errors));
    }

    @ExceptionHandler(ConflictException.class)
    public ResponseEntity<ApiResponse<Void>> handleConflictException(
            ConflictException ex,
            HttpServletRequest request) {

        Map<String, String> errors = new HashMap<>();

        if (ex.getField() != null) {
            errors.put(ex.getField(), ex.getMessage());
        } else {
            errors.put("error", ex.getMessage());
        }

        return ResponseEntity
                .status(HttpStatus.CONFLICT)
                .body(new ApiResponse<>("Conflict Error", errors));
    }

    @ExceptionHandler(UnauthorizedException.class)
    public ResponseEntity<ApiResponse<Void>> handleUnauthorizedException(
            UnauthorizedException ex,
            HttpServletRequest request) {

        Map<String, String> errors = new HashMap<>();

        if (ex.getField() != null) {
            errors.put(ex.getField(), ex.getMessage());
        } else {
            errors.put("error", ex.getMessage());
        }

        return ResponseEntity
                .status(HttpStatus.UNAUTHORIZED)
                .body(new ApiResponse<>("Unauthorized Error", errors));
    }

    @ExceptionHandler(HttpMessageNotReadableException.class)
    public ResponseEntity<ApiResponse<Void>> handleNotReadableException(
            HttpMessageNotReadableException ex) {

        Map<String, String> errors = new HashMap<>();
        errors.put("error", "Malformed JSON request body");

        return ResponseEntity
                .status(HttpStatus.BAD_REQUEST)
                .body(new ApiResponse<>("Bad Request Error", errors));
    }

    @ExceptionHandler(Exception.class)
    public ResponseEntity<ApiResponse<Void>> handleGenericException(
            Exception ex,
            HttpServletRequest request) {

        Map<String, String> errors = new HashMap<>();
        errors.put("error", "Internal server error");

        return ResponseEntity
                .status(HttpStatus.INTERNAL_SERVER_ERROR)
                .body(new ApiResponse<>("Internal Server Error", errors));
    }

    private String capitalize(String value) {
    if (value == null || value.isEmpty()) {
        return value;
    }

    return value.substring(0, 1).toUpperCase() + value.substring(1);
}
}