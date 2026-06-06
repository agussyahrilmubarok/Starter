package io.github.agussyahrilmubarok.backend.application.service;

import io.github.agussyahrilmubarok.backend.application.dto.user.CreateUserRequest;
import io.github.agussyahrilmubarok.backend.application.dto.user.UpdateUserRequest;
import io.github.agussyahrilmubarok.backend.application.dto.user.UserPageRequest;
import io.github.agussyahrilmubarok.backend.application.dto.user.UserResponse;
import java.util.UUID;
import org.springframework.data.domain.Page;

public interface UserService {
    Page<UserResponse> getAll(final UserPageRequest request);

    UserResponse getById(final UUID id);

    UserResponse create(final CreateUserRequest request);

    UserResponse update(final UUID id, final UpdateUserRequest request);

    void delete(final UUID id);
}
