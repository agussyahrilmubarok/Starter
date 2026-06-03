package io.github.agussyahrilmubarok.backend.application.dto.user;

import io.github.agussyahrilmubarok.backend.domain.User;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import java.time.OffsetDateTime;
import java.time.ZoneOffset;
import java.util.UUID;

import static org.assertj.core.api.Assertions.assertThat;

class UserMapperTest {

    private UserMapper userMapper;

    private User buildUser(UUID id, String name, String email) {
        User user = new User();
        user.setId(id);
        user.setName(name);
        user.setEmail(email);
        user.setPassword("hashed_password");
        return user;
    }

    @BeforeEach
    void setUp() {
        userMapper = new UserMapper();
    }

    @Test
    @DisplayName("should map all fields correctly")
    void shouldMapAllFields() {
        UUID id = UUID.randomUUID();
        OffsetDateTime now = OffsetDateTime.now(ZoneOffset.UTC);

        User alice = buildUser(id, "Alice", "alice@example.com");
        alice.setCreatedAt(now);
        alice.setUpdatedAt(now);

        UserResponse response = userMapper.toResponse(alice);

        assertThat(response.id()).isEqualTo(id);
        assertThat(response.name()).isEqualTo("Alice");
        assertThat(response.email()).isEqualTo("alice@example.com");
        assertThat(response.createdAt()).isEqualTo(now.toLocalDateTime());
        assertThat(response.updatedAt()).isEqualTo(now.toLocalDateTime());
    }

    @Test
    @DisplayName("should return null createdAt when createdAt is null")
    void shouldReturnNullCreatedAt() {
        User alice = buildUser(UUID.randomUUID(), "Alice", "alice@example.com");
        alice.setCreatedAt(null);
        alice.setUpdatedAt(OffsetDateTime.now());

        UserResponse response = userMapper.toResponse(alice);

        assertThat(response.createdAt()).isNull();
    }

    @Test
    @DisplayName("should return null updatedAt when updatedAt is null")
    void shouldReturnNullUpdatedAt() {
        User alice = buildUser(UUID.randomUUID(), "Alice", "alice@example.com");
        alice.setCreatedAt(OffsetDateTime.now());
        alice.setUpdatedAt(null);

        UserResponse response = userMapper.toResponse(alice);

        assertThat(response.updatedAt()).isNull();
    }

    @Test
    @DisplayName("should return null for both timestamps when both are null")
    void shouldReturnNullWhenBothTimestampsNull() {
        User alice = buildUser(UUID.randomUUID(), "Alice", "alice@example.com");
        alice.setCreatedAt(null);
        alice.setUpdatedAt(null);

        UserResponse response = userMapper.toResponse(alice);

        assertThat(response.createdAt()).isNull();
        assertThat(response.updatedAt()).isNull();
    }
}