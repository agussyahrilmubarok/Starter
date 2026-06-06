package io.github.agussyahrilmubarok.backend.application.dto.user;

import io.github.agussyahrilmubarok.backend.domain.User;
import org.springframework.stereotype.Component;

@Component
public class UserMapper {

    public UserResponse toResponse(User user) {
        return new UserResponse(
                user.getId(), user.getName(), user.getEmail(), user.getCreatedAt(), user.getUpdatedAt());
    }
}
