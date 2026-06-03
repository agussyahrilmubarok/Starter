package io.github.agussyahrilmubarok.backend.application.service.impl;

import io.github.agussyahrilmubarok.backend.application.dto.user.*;
import io.github.agussyahrilmubarok.backend.common.exception.EmailAlreadyExistsException;
import io.github.agussyahrilmubarok.backend.common.exception.NotFoundException;
import io.github.agussyahrilmubarok.backend.domain.User;
import io.github.agussyahrilmubarok.backend.infrastructure.persistence.repository.UserRepository;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;
import org.springframework.data.domain.*;
import org.springframework.data.jpa.domain.Specification;
import org.springframework.security.crypto.password.PasswordEncoder;

import java.time.LocalDateTime;
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
                LocalDateTime.now(),
                LocalDateTime.now()
        );
    }

    @Nested
    @DisplayName("getAll()")
    class GetAll {

        @SuppressWarnings("unchecked")
        @Test
        @DisplayName("should return paginated users with default params")
        void shouldReturnPaginatedUsers() {
            UUID id1 = UUID.randomUUID();
            UUID id2 = UUID.randomUUID();
            User user1 = buildUser(id1, "Alice", "alice@example.com");
            User user2 = buildUser(id2, "Bob", "bob@example.com");
            UserResponse response1 = buildUserResponse(user1);
            UserResponse response2 = buildUserResponse(user2);

            UserPageRequest request = new UserPageRequest(1, 10, "created_at,desc", null);

            Page<User> userPage = new PageImpl<>(
                    List.of(user1, user2),
                    PageRequest.of(0, 10, Sort.by(Sort.Direction.DESC, "createdAt")),
                    2
            );

            when(userRepository.findAll(any(Specification.class), any(Pageable.class))).thenReturn(userPage);
            when(userMapper.toResponse(user1)).thenReturn(response1);
            when(userMapper.toResponse(user2)).thenReturn(response2);

            Page<UserResponse> result = userService.getAll(request);

            assertThat(result.getContent()).hasSize(2);
            assertThat(result.getContent()).containsExactly(response1, response2);
            assertThat(result.getTotalElements()).isEqualTo(2);
            assertThat(result.getNumber()).isEqualTo(0);
            verify(userRepository).findAll(any(Specification.class), any(Pageable.class));
        }

        @SuppressWarnings("unchecked")
        @Test
        @DisplayName("should return empty page when no users exist")
        void shouldReturnEmptyPage() {
            UserPageRequest request = new UserPageRequest(1, 10, "created_at,desc", null);

            Page<User> emptyPage = new PageImpl<>(
                    List.of(),
                    PageRequest.of(0, 10, Sort.by(Sort.Direction.DESC, "createdAt")),
                    0
            );

            when(userRepository.findAll(any(Specification.class), any(Pageable.class))).thenReturn(emptyPage);

            Page<UserResponse> result = userService.getAll(request);

            assertThat(result.getContent()).isEmpty();
            assertThat(result.getTotalElements()).isEqualTo(0);
        }

        @SuppressWarnings("unchecked")
        @Test
        @DisplayName("should pass search keyword to repository via specification")
        void shouldPassSearchKeyword() {
            UserPageRequest request = new UserPageRequest(1, 10, "name,asc", "alice");

            Page<User> emptyPage = new PageImpl<>(
                    List.of(),
                    PageRequest.of(0, 10, Sort.by(Sort.Direction.ASC, "name")),
                    0
            );

            when(userRepository.findAll(any(Specification.class), any(Pageable.class))).thenReturn(emptyPage);

            userService.getAll(request);

            verify(userRepository).findAll(any(Specification.class), any(Pageable.class));
        }

        @SuppressWarnings("unchecked")
        @Test
        @DisplayName("should treat blank search as null specification")
        void shouldTreatBlankSearchAsNull() {
            UserPageRequest request = new UserPageRequest(1, 10, "created_at,desc", "   ");

            Page<User> emptyPage = new PageImpl<>(
                    List.of(),
                    PageRequest.of(0, 10, Sort.by(Sort.Direction.DESC, "createdAt")),
                    0
            );

            when(userRepository.findAll(any(Specification.class), any(Pageable.class))).thenReturn(emptyPage);

            userService.getAll(request);

            verify(userRepository).findAll(any(Specification.class), any(Pageable.class));
        }

        @SuppressWarnings("unchecked")
        @Test
        @DisplayName("should return correct pagination metadata")
        void shouldReturnCorrectPaginationMetadata() {
            UserPageRequest request = new UserPageRequest(2, 5, "created_at,desc", null);

            List<User> users = List.of(
                    buildUser(UUID.randomUUID(), "Alice", "alice@example.com"),
                    buildUser(UUID.randomUUID(), "Bob", "bob@example.com")
            );

            Page<User> page = new PageImpl<>(
                    users,
                    PageRequest.of(1, 5, Sort.by(Sort.Direction.DESC, "createdAt")),
                    12
            );

            when(userRepository.findAll(any(Specification.class), any(Pageable.class))).thenReturn(page);
            when(userMapper.toResponse(any(User.class))).thenReturn(buildUserResponse(users.getFirst()));

            Page<UserResponse> result = userService.getAll(request);

            assertThat(result.getTotalElements()).isEqualTo(12);
            assertThat(result.getTotalPages()).isEqualTo(3);
            assertThat(result.getNumber()).isEqualTo(1);
            assertThat(result.hasNext()).isTrue();
            assertThat(result.hasPrevious()).isTrue();
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
                    .hasMessageContaining(id.toString());
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

            when(userRepository.existsByEmail(request.email())).thenReturn(false);
            when(passwordEncoder.encode(request.password())).thenReturn("hashed_password");
            when(userRepository.save(any(User.class))).thenReturn(savedUser);
            when(userMapper.toResponse(any(User.class))).thenReturn(response);

            UserResponse result = userService.create(request);

            assertThat(result).isEqualTo(response);
            verify(userRepository).existsByEmail(request.email());
            verify(passwordEncoder).encode(request.password());
            verify(userRepository).save(any(User.class));
        }

        @Test
        @DisplayName("should save email in lowercase")
        void shouldSaveEmailLowercase() {
            CreateUserRequest request = new CreateUserRequest("Alice", "ALICE@EXAMPLE.COM", "password123");
            User savedUser = buildUser(UUID.randomUUID(), "Alice", "alice@example.com");

            when(userRepository.existsByEmail(anyString())).thenReturn(false);
            when(passwordEncoder.encode(anyString())).thenReturn("hashed");
            when(userRepository.save(any(User.class))).thenReturn(savedUser);
            when(userMapper.toResponse(any(User.class))).thenReturn(buildUserResponse(savedUser));

            userService.create(request);

            verify(userRepository).save(argThat(u -> u.getEmail().equals("alice@example.com")));
        }

        @Test
        @DisplayName("should throw EmailAlreadyExistsException when email is taken")
        void shouldThrowWhenEmailAlreadyExists() {
            CreateUserRequest request = new CreateUserRequest("Alice", "alice@example.com", "password123");

            when(userRepository.existsByEmail(request.email())).thenReturn(true);

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
            UpdateUserRequest request = new UpdateUserRequest("New Name", null, null);
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
            UpdateUserRequest request = new UpdateUserRequest(null, "NEW@EXAMPLE.COM", null);

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
            UpdateUserRequest request = new UpdateUserRequest(null, null, "newpassword123");

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
            UpdateUserRequest request = new UpdateUserRequest(null, "alice@example.com", null);

            when(userRepository.findById(userId)).thenReturn(Optional.of(existingUser));
            when(userRepository.save(any(User.class))).thenReturn(existingUser);
            when(userMapper.toResponse(any(User.class))).thenReturn(buildUserResponse(existingUser));

            userService.update(userId, request);

            verify(userRepository, never()).existsByEmail(anyString());
        }

        @Test
        @DisplayName("should throw EmailAlreadyExistsException when new email is taken")
        void shouldThrowWhenNewEmailTaken() {
            UpdateUserRequest request = new UpdateUserRequest(null, "taken@example.com", null);

            when(userRepository.findById(userId)).thenReturn(Optional.of(existingUser));
            when(userRepository.existsByEmail("taken@example.com")).thenReturn(true);

            assertThatThrownBy(() -> userService.update(userId, request))
                    .isInstanceOf(EmailAlreadyExistsException.class);

            verify(userRepository, never()).save(any());
        }

        @Test
        @DisplayName("should throw NotFoundException when user not found")
        void shouldThrowNotFoundOnUpdate() {
            UpdateUserRequest request = new UpdateUserRequest("Name", null, null);

            when(userRepository.findById(userId)).thenReturn(Optional.empty());

            assertThatThrownBy(() -> userService.update(userId, request))
                    .isInstanceOf(NotFoundException.class)
                    .hasMessageContaining(userId.toString());
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
                    .hasMessageContaining(id.toString());

            verify(userRepository, never()).deleteById(any());
        }
    }
}