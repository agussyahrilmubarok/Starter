package io.github.agussyahrilmubarok.web.service;

import io.github.agussyahrilmubarok.web.model.user.CreateUserRequest;
import io.github.agussyahrilmubarok.web.model.user.UpdateUserRequest;
import io.github.agussyahrilmubarok.web.model.user.UserResponse;

import java.util.List;
import java.util.UUID;

public interface IUserService {

    List<UserResponse> findAll();

    UserResponse findById(UUID userId);

    UserResponse create(CreateUserRequest request);

    UserResponse updateById(UUID userId, UpdateUserRequest request);

    void deleteById(UUID userId);
}
