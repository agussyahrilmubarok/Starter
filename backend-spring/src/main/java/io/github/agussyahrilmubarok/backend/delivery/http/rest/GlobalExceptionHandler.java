package io.github.agussyahrilmubarok.backend.delivery.http.rest;

import io.github.agussyahrilmubarok.backend.common.exception.*;
import io.github.agussyahrilmubarok.backend.delivery.http.model.ApiResponse;
import java.util.Map;
import java.util.stream.Collectors;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.validation.FieldError;
import org.springframework.web.bind.MethodArgumentNotValidException;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.RestControllerAdvice;

@RestControllerAdvice
public class GlobalExceptionHandler {

    @ExceptionHandler(EmailAlreadyInUseException.class)
    public ResponseEntity<ApiResponse<Object>> handleEmailAlreadyExistsException(EmailAlreadyInUseException ex) {
        return ResponseEntity.status(HttpStatus.CONFLICT)
                .body(new ApiResponse<>("Conflict", Map.of("email", ex.getMessage())));
    }

    @ExceptionHandler(EmailNotRegisteredException.class)
    public ResponseEntity<ApiResponse<Object>> handleEmailNotRegisteredException(EmailNotRegisteredException ex) {
        return ResponseEntity.status(HttpStatus.UNAUTHORIZED)
                .body(new ApiResponse<>("Unauthorized", Map.of("email", ex.getMessage())));
    }

    @ExceptionHandler(JwtExpiredTokenException.class)
    public ResponseEntity<ApiResponse<Object>> handleJwtExpiredTokenException(JwtExpiredTokenException ex) {
        return ResponseEntity.status(HttpStatus.UNAUTHORIZED)
                .body(new ApiResponse<>("Unauthorized", Map.of("error", ex.getMessage())));
    }

    @ExceptionHandler(JwtInvalidTokenException.class)
    public ResponseEntity<ApiResponse<Object>> handleJwtInvalidTokenException(JwtInvalidTokenException ex) {
        return ResponseEntity.status(HttpStatus.UNAUTHORIZED)
                .body(new ApiResponse<>("Unauthorized", Map.of("error", ex.getMessage())));
    }

    @ExceptionHandler(NotFoundException.class)
    public ResponseEntity<ApiResponse<Object>> handleNotFoundException(NotFoundException ex) {
        return ResponseEntity.status(HttpStatus.NOT_FOUND)
                .body(new ApiResponse<>("Not found", Map.of("error", ex.getMessage())));
    }

    @ExceptionHandler(PasswordMismatchException.class)
    public ResponseEntity<ApiResponse<Object>> handleWrongPasswordException(PasswordMismatchException ex) {
        return ResponseEntity.status(HttpStatus.UNAUTHORIZED)
                .body(new ApiResponse<>("Unauthorized", Map.of("password", ex.getMessage())));
    }

    @ExceptionHandler(MethodArgumentNotValidException.class)
    public ResponseEntity<ApiResponse<Object>> handleValidationException(MethodArgumentNotValidException ex) {
        Map<String, String> errors = ex.getBindingResult().getFieldErrors().stream()
                .collect(Collectors.toMap(
                        FieldError::getField,
                        field -> field.getDefaultMessage() != null ? field.getDefaultMessage() : "Invalid value",
                        (existing, duplicate) -> existing));
        return ResponseEntity.status(HttpStatus.BAD_REQUEST).body(new ApiResponse<>("Validation failed", errors));
    }

    @ExceptionHandler(Exception.class)
    public ResponseEntity<ApiResponse<Object>> handleException(Exception ex) {
        return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR)
                .body(new ApiResponse<>("Internal server error", Map.of("error", "Something went wrong")));
    }
}
