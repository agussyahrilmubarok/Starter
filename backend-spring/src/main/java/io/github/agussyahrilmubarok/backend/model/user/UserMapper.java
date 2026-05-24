package io.github.agussyahrilmubarok.backend.model.user;

import io.github.agussyahrilmubarok.backend.domain.User;
import org.springframework.stereotype.Component;

@Component
public class UserMapper {

    public UserResponse toResponse(User user) {
        return new UserResponse(
                user.getId().toString(),
                user.getName(),
                user.getEmail(),
                user.getCreatedAt() != null ? user.getCreatedAt().toLocalDateTime() : null,
                user.getUpdatedAt() != null ? user.getUpdatedAt().toLocalDateTime() : null);
    }
}
