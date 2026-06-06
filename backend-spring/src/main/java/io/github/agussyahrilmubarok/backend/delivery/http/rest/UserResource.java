package io.github.agussyahrilmubarok.backend.delivery.http.rest;

import io.github.agussyahrilmubarok.backend.application.dto.user.CreateUserRequest;
import io.github.agussyahrilmubarok.backend.application.dto.user.UpdateUserRequest;
import io.github.agussyahrilmubarok.backend.application.dto.user.UserPageRequest;
import io.github.agussyahrilmubarok.backend.application.dto.user.UserResponse;
import io.github.agussyahrilmubarok.backend.application.service.UserService;
import io.github.agussyahrilmubarok.backend.delivery.http.model.ApiResponse;
import io.github.agussyahrilmubarok.backend.delivery.http.model.PagedApiResponse;
import io.github.agussyahrilmubarok.backend.delivery.http.model.PaginationResponse;
import jakarta.validation.Valid;
import jakarta.validation.constraints.Max;
import jakarta.validation.constraints.Min;
import java.util.List;
import java.util.UUID;
import lombok.RequiredArgsConstructor;
import org.springframework.data.domain.Page;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping(value = "/api/v2/users", produces = MediaType.APPLICATION_JSON_VALUE)
@RequiredArgsConstructor
public class UserResource {

    private final UserService userService;

    @GetMapping
    public ResponseEntity<PagedApiResponse<List<UserResponse>>> getAll(
            @RequestParam(defaultValue = "1") @Min(value = 1, message = "Must be a positive integer") int page,
            @RequestParam(defaultValue = "10")
                    @Min(value = 1, message = "Must be between 1 and 100")
                    @Max(value = 100, message = "Must be between 1 and 100")
                    int size,
            @RequestParam(defaultValue = "created_at,desc") String sort,
            @RequestParam(required = false) String search) {
        UserPageRequest request = new UserPageRequest(page, size, sort, search);
        Page<UserResponse> result = userService.getAll(request);
        PaginationResponse pagination = new PaginationResponse(
                result.getNumber() + 1,
                result.getSize(),
                result.getTotalElements(),
                result.getTotalPages(),
                result.hasNext(),
                result.hasPrevious());
        return ResponseEntity.ok(
                new PagedApiResponse<>("Users retrieved successfully", result.getContent(), pagination));
    }

    @GetMapping("/{userId}")
    public ResponseEntity<ApiResponse<UserResponse>> getById(@PathVariable UUID userId) {
        UserResponse response = userService.getById(userId);
        return ResponseEntity.ok(new ApiResponse<>("User retrieved successfully", response));
    }

    @PostMapping
    public ResponseEntity<ApiResponse<UserResponse>> create(@Valid @RequestBody CreateUserRequest request) {
        UserResponse response = userService.create(request);
        return ResponseEntity.status(HttpStatus.CREATED).body(new ApiResponse<>("User created successfully", response));
    }

    @PutMapping("/{userId}")
    public ResponseEntity<ApiResponse<UserResponse>> updateById(
            @PathVariable UUID userId, @Valid @RequestBody UpdateUserRequest request) {
        UserResponse response = userService.update(userId, request);
        return ResponseEntity.ok(new ApiResponse<>("User updated successfully", response));
    }

    @DeleteMapping("/{userId}")
    public ResponseEntity<ApiResponse<Void>> deleteById(@PathVariable UUID userId) {
        userService.delete(userId);
        return ResponseEntity.ok(new ApiResponse<>("User deleted successfully", null));
    }
}
