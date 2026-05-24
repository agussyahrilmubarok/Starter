package io.github.agussyahrilmubarok.backend.rest;

import io.github.agussyahrilmubarok.backend.model.ApiResponse;
import io.github.agussyahrilmubarok.backend.model.user.CreateUserRequest;
import io.github.agussyahrilmubarok.backend.model.user.UpdateUserRequest;
import io.github.agussyahrilmubarok.backend.model.user.UserResponse;
import io.github.agussyahrilmubarok.backend.service.IUserService;
import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping(value = "/api/v1/users", produces = MediaType.APPLICATION_JSON_VALUE)
@RequiredArgsConstructor
public class UserResource {

    private final IUserService userService;

    @GetMapping
    public ResponseEntity<ApiResponse<List<UserResponse>>> findAll() {
        List<UserResponse> response = userService.findAll();
        return ResponseEntity.ok(new ApiResponse<>("Users retrieved successfully", response));
    }

    @GetMapping("/{userId}")
    public ResponseEntity<ApiResponse<UserResponse>> findById(@PathVariable String userId) {
        UserResponse response = userService.findById(userId);
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
            @PathVariable String userId,
            @Valid @RequestBody UpdateUserRequest request) {
        UserResponse response = userService.updateById(userId, request);
        return ResponseEntity.ok(new ApiResponse<>("User updated successfully", response));
    }

    @DeleteMapping("/{userId}")
    public ResponseEntity<ApiResponse<Void>> deleteById(@PathVariable String userId) {
        userService.deleteById(userId);
        return ResponseEntity.ok(new ApiResponse<>("User deleted successfully", null));
    }
}