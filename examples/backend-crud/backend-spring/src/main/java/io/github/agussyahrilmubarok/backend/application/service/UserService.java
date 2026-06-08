package io.github.agussyahrilmubarok.backend.application.service;

import io.github.agussyahrilmubarok.backend.application.dto.user.CreateUserRequest;
import io.github.agussyahrilmubarok.backend.application.dto.user.UpdateUserRequest;
import io.github.agussyahrilmubarok.backend.application.dto.user.UserResponse;
import java.util.List;
import java.util.UUID;

public interface UserService {
    List<UserResponse> getAll();

    UserResponse getById(final UUID id);

    UserResponse create(final CreateUserRequest request);

    UserResponse update(final UUID id, final UpdateUserRequest request);

    void delete(final UUID id);
}
