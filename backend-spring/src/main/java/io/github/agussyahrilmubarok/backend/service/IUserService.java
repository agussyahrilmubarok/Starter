package io.github.agussyahrilmubarok.backend.service;

import io.github.agussyahrilmubarok.backend.model.user.CreateUserRequest;
import io.github.agussyahrilmubarok.backend.model.user.UpdateUserRequest;
import io.github.agussyahrilmubarok.backend.model.user.UserResponse;

import java.util.List;

public interface IUserService {
    List<UserResponse> findAll();

    UserResponse findById(String userId);

    UserResponse create(CreateUserRequest request);

    UserResponse updateById(String userId, UpdateUserRequest request);

    void deleteById(String userId);
}
