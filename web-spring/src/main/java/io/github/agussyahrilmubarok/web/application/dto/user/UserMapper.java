package io.github.agussyahrilmubarok.web.application.dto.user;

import io.github.agussyahrilmubarok.web.domain.User;
import org.springframework.stereotype.Component;

@Component
public class UserMapper {

    public UserResponse toResponse(User user) {
        return new UserResponse(
                user.getId(),
                user.getName(),
                user.getEmail(),
                "",
                user.getCreatedAt() != null ? user.getCreatedAt().toLocalDateTime() : null,
                user.getUpdatedAt() != null ? user.getUpdatedAt().toLocalDateTime() : null);
    }
}