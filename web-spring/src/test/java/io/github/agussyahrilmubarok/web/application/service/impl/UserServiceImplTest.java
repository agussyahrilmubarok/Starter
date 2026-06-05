package io.github.agussyahrilmubarok.web.application.service.impl;

import io.github.agussyahrilmubarok.web.application.dto.user.CreateUserRequest;
import io.github.agussyahrilmubarok.web.application.dto.user.UpdateUserRequest;
import io.github.agussyahrilmubarok.web.application.dto.user.UserMapper;
import io.github.agussyahrilmubarok.web.application.dto.user.UserResponse;
import io.github.agussyahrilmubarok.web.common.exception.EmailAlreadyExistsException;
import io.github.agussyahrilmubarok.web.common.exception.NotFoundException;
import io.github.agussyahrilmubarok.web.domain.User;
import io.github.agussyahrilmubarok.web.infrastructure.persistence.repository.UserRepository;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;
import org.springframework.security.crypto.password.PasswordEncoder;

import java.time.OffsetDateTime;
import java.util.List;
import java.util.Optional;
import java.util.UUID;

import static org.assertj.core.api.Assertions.assertThat;
import static org.assertj.core.api.Assertions.assertThatThrownBy;
import static org.mockito.ArgumentMatchers.*;
import static org.mockito.Mockito.*;

@ExtendWith(MockitoExtension.class)
class UserServiceImplTest {

    @Mock
    private UserRepository userRepository;

    @Mock
    private UserMapper userMapper;

    @Mock
    private PasswordEncoder passwordEncoder;

    @InjectMocks
    private UserServiceImpl userService;

    private User buildUser(UUID id, String name, String email) {
        User user = new User();
        user.setId(id);
        user.setName(name);
        user.setEmail(email);
        user.setPassword("hashed_password");
        user.setCreatedAt(OffsetDateTime.now());
        user.setUpdatedAt(OffsetDateTime.now());
        return user;
    }

    private UserResponse buildUserResponse(User user) {
        return new UserResponse(
                user.getId(),
                user.getName(),
                user.getEmail(),
                OffsetDateTime.now(),
                OffsetDateTime.now()
        );
    }

    @Nested
    @DisplayName("getAll()")
    class GetAll {

        @Test
        @DisplayName("should return list of all users")
        void shouldReturnAllUsers() {
            UUID id1 = UUID.randomUUID();
            UUID id2 = UUID.randomUUID();
            User user1 = buildUser(id1, "Alice", "alice@example.com");
            User user2 = buildUser(id2, "Bob", "bob@example.com");
            UserResponse response1 = buildUserResponse(user1);
            UserResponse response2 = buildUserResponse(user2);

            when(userRepository.findAll()).thenReturn(List.of(user1, user2));
            when(userMapper.toResponse(user1)).thenReturn(response1);
            when(userMapper.toResponse(user2)).thenReturn(response2);

            List<UserResponse> result = userService.getAll();

            assertThat(result).hasSize(2);
            assertThat(result).containsExactly(response1, response2);
            verify(userRepository).findAll();
        }

        @Test
        @DisplayName("should return empty list when no users exist")
        void shouldReturnEmptyList() {
            when(userRepository.findAll()).thenReturn(List.of());

            List<UserResponse> result = userService.getAll();

            assertThat(result).isEmpty();
        }
    }

    @Nested
    @DisplayName("getById()")
    class GetById {

        @Test
        @DisplayName("should return user when found")
        void shouldReturnUserWhenFound() {
            UUID id = UUID.randomUUID();
            User user = buildUser(id, "Alice", "alice@example.com");
            UserResponse response = buildUserResponse(user);

            when(userRepository.findById(id)).thenReturn(Optional.of(user));
            when(userMapper.toResponse(user)).thenReturn(response);

            UserResponse result = userService.getById(id);

            assertThat(result).isEqualTo(response);
            verify(userRepository).findById(id);
        }

        @Test
        @DisplayName("should throw NotFoundException when user not found")
        void shouldThrowNotFoundWhenUserNotFound() {
            UUID id = UUID.randomUUID();

            when(userRepository.findById(id)).thenReturn(Optional.empty());

            assertThatThrownBy(() -> userService.getById(id))
                    .isInstanceOf(NotFoundException.class)
                    .hasMessage("User not found");
        }
    }

    @Nested
    @DisplayName("create()")
    class Create {

        @Test
        @DisplayName("should create and return user when email is not taken")
        void shouldCreateUserSuccessfully() {
            CreateUserRequest request = new CreateUserRequest("Alice", "alice@example.com", "password123");
            User savedUser = buildUser(UUID.randomUUID(), "Alice", "alice@example.com");
            UserResponse response = buildUserResponse(savedUser);

            when(userRepository.existsByEmail("alice@example.com")).thenReturn(false);
            when(passwordEncoder.encode(request.password())).thenReturn("hashed_password");
            when(userRepository.save(any(User.class))).thenReturn(savedUser);
            when(userMapper.toResponse(any(User.class))).thenReturn(response);

            UserResponse result = userService.create(request);

            assertThat(result).isEqualTo(response);
            verify(userRepository).existsByEmail("alice@example.com");
            verify(passwordEncoder).encode(request.password());
            verify(userRepository).save(any(User.class));
        }

        @Test
        @DisplayName("should save email in lowercase")
        void shouldSaveEmailLowercase() {
            CreateUserRequest request = new CreateUserRequest("Alice", "ALICE@EXAMPLE.COM", "password123");
            User savedUser = buildUser(UUID.randomUUID(), "Alice", "alice@example.com");

            when(userRepository.existsByEmail("alice@example.com")).thenReturn(false);
            when(passwordEncoder.encode(anyString())).thenReturn("hashed");
            when(userRepository.save(any(User.class))).thenReturn(savedUser);
            when(userMapper.toResponse(any(User.class))).thenReturn(buildUserResponse(savedUser));

            userService.create(request);

            verify(userRepository).existsByEmail("alice@example.com");
            verify(userRepository).save(argThat(u -> u.getEmail().equals("alice@example.com")));
        }

        @Test
        @DisplayName("should throw EmailAlreadyExistsException when email is taken")
        void shouldThrowWhenEmailAlreadyExists() {
            CreateUserRequest request = new CreateUserRequest("Alice", "alice@example.com", "password123");

            when(userRepository.existsByEmail("alice@example.com")).thenReturn(true);

            assertThatThrownBy(() -> userService.create(request))
                    .isInstanceOf(EmailAlreadyExistsException.class);

            verify(userRepository, never()).save(any());
        }
    }

    @Nested
    @DisplayName("update()")
    class Update {

        private UUID userId;
        private User existingUser;

        @BeforeEach
        void setUp() {
            userId = UUID.randomUUID();
            existingUser = buildUser(userId, "Alice", "alice@example.com");
        }

        @Test
        @DisplayName("should update name only")
        void shouldUpdateNameOnly() {
            UpdateUserRequest request = new UpdateUserRequest(userId, "New Name", null, null);
            UserResponse response = buildUserResponse(existingUser);

            when(userRepository.findById(userId)).thenReturn(Optional.of(existingUser));
            when(userRepository.save(any(User.class))).thenReturn(existingUser);
            when(userMapper.toResponse(any(User.class))).thenReturn(response);

            userService.update(userId, request);

            assertThat(existingUser.getName()).isEqualTo("New Name");
            verify(passwordEncoder, never()).encode(anyString());
        }

        @Test
        @DisplayName("should update email to lowercase when new email provided")
        void shouldUpdateEmailLowercase() {
            UpdateUserRequest request = new UpdateUserRequest(userId, null, "NEW@EXAMPLE.COM", null);

            when(userRepository.findById(userId)).thenReturn(Optional.of(existingUser));
            when(userRepository.existsByEmail("new@example.com")).thenReturn(false);
            when(userRepository.save(any(User.class))).thenReturn(existingUser);
            when(userMapper.toResponse(any(User.class))).thenReturn(buildUserResponse(existingUser));

            userService.update(userId, request);

            assertThat(existingUser.getEmail()).isEqualTo("new@example.com");
        }

        @Test
        @DisplayName("should encode and update password when provided")
        void shouldUpdatePassword() {
            UpdateUserRequest request = new UpdateUserRequest(userId, null, null, "newpassword123");

            when(userRepository.findById(userId)).thenReturn(Optional.of(existingUser));
            when(passwordEncoder.encode("newpassword123")).thenReturn("new_hashed");
            when(userRepository.save(any(User.class))).thenReturn(existingUser);
            when(userMapper.toResponse(any(User.class))).thenReturn(buildUserResponse(existingUser));

            userService.update(userId, request);

            assertThat(existingUser.getPassword()).isEqualTo("new_hashed");
        }

        @Test
        @DisplayName("should not update email when same as current")
        void shouldNotUpdateEmailWhenSame() {
            UpdateUserRequest request = new UpdateUserRequest(userId, null, "alice@example.com", null);

            when(userRepository.findById(userId)).thenReturn(Optional.of(existingUser));
            when(userRepository.save(any(User.class))).thenReturn(existingUser);
            when(userMapper.toResponse(any(User.class))).thenReturn(buildUserResponse(existingUser));

            userService.update(userId, request);

            verify(userRepository, never()).existsByEmail(anyString());
        }

        @Test
        @DisplayName("should throw EmailAlreadyExistsException when new email is taken")
        void shouldThrowWhenNewEmailTaken() {
            UpdateUserRequest request = new UpdateUserRequest(userId, null, "taken@example.com", null);

            when(userRepository.findById(userId)).thenReturn(Optional.of(existingUser));
            when(userRepository.existsByEmail("taken@example.com")).thenReturn(true);

            assertThatThrownBy(() -> userService.update(userId, request))
                    .isInstanceOf(EmailAlreadyExistsException.class);

            verify(userRepository, never()).save(any());
        }

        @Test
        @DisplayName("should throw NotFoundException when user not found")
        void shouldThrowNotFoundOnUpdate() {
            UpdateUserRequest request = new UpdateUserRequest(userId, "Name", null, null);

            when(userRepository.findById(userId)).thenReturn(Optional.empty());

            assertThatThrownBy(() -> userService.update(userId, request))
                    .isInstanceOf(NotFoundException.class)
                    .hasMessage("User not found");
        }
    }

    @Nested
    @DisplayName("delete()")
    class Delete {

        @Test
        @DisplayName("should delete user when found")
        void shouldDeleteUserSuccessfully() {
            UUID id = UUID.randomUUID();
            User user = buildUser(id, "Alice", "alice@example.com");

            when(userRepository.findById(id)).thenReturn(Optional.of(user));

            userService.delete(id);

            verify(userRepository).deleteById(id);
        }

        @Test
        @DisplayName("should throw NotFoundException when user not found")
        void shouldThrowNotFoundOnDelete() {
            UUID id = UUID.randomUUID();

            when(userRepository.findById(id)).thenReturn(Optional.empty());

            assertThatThrownBy(() -> userService.delete(id))
                    .isInstanceOf(NotFoundException.class)
                    .hasMessage("User not found");

            verify(userRepository, never()).deleteById(any());
        }
    }
}