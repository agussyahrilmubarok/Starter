package io.github.agussyahrilmubarok.backend.delivery.http.rest;

import io.github.agussyahrilmubarok.backend.application.dto.user.CreateUserRequest;
import io.github.agussyahrilmubarok.backend.application.dto.user.UpdateUserRequest;
import io.github.agussyahrilmubarok.backend.application.dto.user.UserResponse;
import io.github.agussyahrilmubarok.backend.application.service.UserService;
import io.github.agussyahrilmubarok.backend.delivery.http.payload.ApiResponse;
import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.UUID;

@RestController
@RequestMapping(value = "/api/v1/users", produces = MediaType.APPLICATION_JSON_VALUE)
@RequiredArgsConstructor
public class UserResource {

    private final UserService userService;

    @GetMapping
    public ResponseEntity<ApiResponse<List<UserResponse>>> getAll() {
        List<UserResponse> response = userService.getAll();
        return ResponseEntity.ok(new ApiResponse<>("Users retrieved successfully", response));
    }

    @GetMapping("/{userId}")
    public ResponseEntity<ApiResponse<UserResponse>> getById(@PathVariable UUID userId) {
        UserResponse response = userService.getById(userId);
        return ResponseEntity.ok(new ApiResponse<>("User retrieved successfully", response));
    }

    @PostMapping
    public ResponseEntity<ApiResponse<UserResponse>> create(
            @Valid @RequestBody CreateUserRequest request) {
        UserResponse response = userService.create(request);
        return ResponseEntity.status(HttpStatus.CREATED)
                .body(new ApiResponse<>("User created successfully", response));
    }

    @PutMapping("/{userId}")
    public ResponseEntity<ApiResponse<UserResponse>> updateById(
            @PathVariable UUID userId,
            @Valid @RequestBody UpdateUserRequest request) {
        UserResponse response = userService.update(userId, request);
        return ResponseEntity.ok(new ApiResponse<>("User updated successfully", response));
    }

    @DeleteMapping("/{userId}")
    public ResponseEntity<ApiResponse<Void>> deleteById(@PathVariable UUID userId) {
        userService.delete(userId);
        return ResponseEntity.ok(new ApiResponse<>("User deleted successfully", null));
    }
}