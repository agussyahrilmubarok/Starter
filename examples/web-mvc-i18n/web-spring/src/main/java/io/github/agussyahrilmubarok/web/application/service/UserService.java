package io.github.agussyahrilmubarok.web.application.service;

import io.github.agussyahrilmubarok.web.application.dto.user.CreateUserRequest;
import io.github.agussyahrilmubarok.web.application.dto.user.UpdateUserRequest;
import io.github.agussyahrilmubarok.web.application.dto.user.UserResponse;
import java.util.List;
import java.util.UUID;

public interface UserService {
    List<UserResponse> getAll();

    UserResponse getById(final UUID id);

    UserResponse create(final CreateUserRequest request);

    UserResponse update(final UUID id, final UpdateUserRequest request);

    void delete(final UUID id);
}
