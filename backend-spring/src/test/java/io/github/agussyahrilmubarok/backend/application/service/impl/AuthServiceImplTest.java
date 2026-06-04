package io.github.agussyahrilmubarok.backend.application.service.impl;

import io.github.agussyahrilmubarok.backend.application.dto.auth.AuthResponse;
import io.github.agussyahrilmubarok.backend.application.dto.auth.SignInRequest;
import io.github.agussyahrilmubarok.backend.application.dto.auth.SignUpRequest;
import io.github.agussyahrilmubarok.backend.application.dto.user.UserMapper;
import io.github.agussyahrilmubarok.backend.application.dto.user.UserResponse;
import io.github.agussyahrilmubarok.backend.common.exception.EmailAlreadyExistsException;
import io.github.agussyahrilmubarok.backend.common.exception.EmailNotRegisteredException;
import io.github.agussyahrilmubarok.backend.common.exception.WrongPasswordException;
import io.github.agussyahrilmubarok.backend.domain.User;
import io.github.agussyahrilmubarok.backend.infrastructure.persistence.repository.UserRepository;
import io.github.agussyahrilmubarok.backend.infrastructure.security.JwtManager;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.ArgumentCaptor;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;
import org.springframework.security.crypto.password.PasswordEncoder;

import java.time.LocalDateTime;
import java.time.OffsetDateTime;
import java.util.Optional;
import java.util.UUID;

import static org.assertj.core.api.Assertions.assertThat;
import static org.assertj.core.api.Assertions.assertThatThrownBy;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.ArgumentMatchers.anyString;
import static org.mockito.Mockito.*;

@ExtendWith(MockitoExtension.class)
class AuthServiceImplTest {

    @Mock
    private UserRepository userRepository;
    @Mock
    private UserMapper userMapper;
    @Mock
    private PasswordEncoder passwordEncoder;
    @Mock
    private JwtManager jwtManager;

    @InjectMocks
    private AuthServiceImpl authService;

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
    @DisplayName("signUp()")
    class SignUp {

        private SignUpRequest request;

        @BeforeEach
        void setUp() {
            request = new SignUpRequest("Alice", "alice@example.com", "password123");
        }

        @Test
        @DisplayName("should save user and return token + user response")
        void shouldSaveUserAndReturnAuthResponse() {
            UUID aliceId = UUID.randomUUID();

            when(userRepository.existsByEmail("alice@example.com")).thenReturn(false);
            when(passwordEncoder.encode("password123")).thenReturn("hashed_password");
            when(userRepository.save(any(User.class))).thenAnswer(inv -> {
                User u = inv.getArgument(0);
                u.setId(aliceId);
                return u;
            });
            when(jwtManager.generateToken(aliceId.toString())).thenReturn("alice-token");
            when(userMapper.toResponse(any(User.class))).thenAnswer(inv ->
                    buildUserResponse(inv.getArgument(0)));

            AuthResponse response = authService.signUp(request);

            assertThat(response.token()).isEqualTo("alice-token");
            assertThat(response.user()).isNotNull();
        }

        @Test
        @DisplayName("should store email in lowercase")
        void shouldStoreEmailInLowercase() {
            when(userRepository.existsByEmail(anyString())).thenReturn(false);
            when(passwordEncoder.encode(anyString())).thenReturn("hashed_password");
            when(userRepository.save(any(User.class))).thenAnswer(inv -> {
                User u = inv.getArgument(0);
                u.setId(UUID.randomUUID());
                return u;
            });
            when(jwtManager.generateToken(anyString())).thenReturn("alice-token");
            when(userMapper.toResponse(any(User.class))).thenAnswer(inv ->
                    buildUserResponse(inv.getArgument(0)));

            authService.signUp(request);

            ArgumentCaptor<User> captor = ArgumentCaptor.forClass(User.class);
            verify(userRepository).save(captor.capture());
            assertThat(captor.getValue().getEmail()).isEqualTo("alice@example.com");
        }

        @Test
        @DisplayName("should encode password before saving")
        void shouldEncodePasswordBeforeSaving() {
            when(userRepository.existsByEmail(anyString())).thenReturn(false);
            when(passwordEncoder.encode("password123")).thenReturn("hashed_password");
            when(userRepository.save(any(User.class))).thenAnswer(inv -> {
                User u = inv.getArgument(0);
                u.setId(UUID.randomUUID());
                return u;
            });
            when(jwtManager.generateToken(anyString())).thenReturn("alice-token");
            when(userMapper.toResponse(any(User.class))).thenAnswer(inv ->
                    buildUserResponse(inv.getArgument(0)));

            authService.signUp(request);

            ArgumentCaptor<User> captor = ArgumentCaptor.forClass(User.class);
            verify(userRepository).save(captor.capture());
            assertThat(captor.getValue().getPassword()).isEqualTo("hashed_password");
            assertThat(captor.getValue().getPassword()).isNotEqualTo("password123");
        }

        @Test
        @DisplayName("should throw EmailAlreadyExistsException when email is taken")
        void shouldThrowWhenEmailAlreadyExists() {
            when(userRepository.existsByEmail("alice@example.com")).thenReturn(true);

            assertThatThrownBy(() -> authService.signUp(request))
                    .isInstanceOf(EmailAlreadyExistsException.class)
                    .hasMessage("The email has already been taken");

            verify(userRepository, never()).save(any());
            verifyNoInteractions(passwordEncoder);
            verifyNoInteractions(jwtManager);
        }
    }

    @Nested
    @DisplayName("signIn()")
    class SignIn {

        private SignInRequest request;
        private User alice;

        @BeforeEach
        void setUp() {
            alice = buildUser(UUID.randomUUID(), "Alice", "alice@example.com");
            request = new SignInRequest("alice@example.com", "password123");
        }

        @Test
        @DisplayName("should return token + user response on valid credentials")
        void shouldReturnAuthResponseOnValidCredentials() {
            UserResponse response = buildUserResponse(alice);

            when(userRepository.findByEmail("alice@example.com")).thenReturn(Optional.of(alice));
            when(passwordEncoder.matches("password123", "hashed_password")).thenReturn(true);
            when(jwtManager.generateToken(alice.getId().toString())).thenReturn("alice-token");
            when(userMapper.toResponse(alice)).thenReturn(response);

            AuthResponse result = authService.signIn(request);

            assertThat(result.token()).isEqualTo("alice-token");
            assertThat(result.user()).isEqualTo(response);
        }

        @Test
        @DisplayName("should generate token using Alice's id from DB")
        void shouldGenerateTokenUsingAliceId() {
            when(userRepository.findByEmail("alice@example.com")).thenReturn(Optional.of(alice));
            when(passwordEncoder.matches(anyString(), anyString())).thenReturn(true);
            when(jwtManager.generateToken(alice.getId().toString())).thenReturn("alice-token");
            when(userMapper.toResponse(alice)).thenReturn(buildUserResponse(alice));

            authService.signIn(request);

            verify(jwtManager).generateToken(alice.getId().toString());
        }

        @Test
        @DisplayName("should throw EmailNotRegisteredException when email not found")
        void shouldThrowWhenEmailNotFound() {
            when(userRepository.findByEmail("alice@example.com")).thenReturn(Optional.empty());

            assertThatThrownBy(() -> authService.signIn(request))
                    .isInstanceOf(EmailNotRegisteredException.class)
                    .hasMessage("The email address is not registered");

            verifyNoInteractions(passwordEncoder);
            verifyNoInteractions(jwtManager);
        }

        @Test
        @DisplayName("should throw WrongPasswordException when password does not match")
        void shouldThrowWhenPasswordDoesNotMatch() {
            when(userRepository.findByEmail("alice@example.com")).thenReturn(Optional.of(alice));
            when(passwordEncoder.matches("password123", "hashed_password")).thenReturn(false);

            assertThatThrownBy(() -> authService.signIn(request))
                    .isInstanceOf(WrongPasswordException.class)
                    .hasMessage("The password is incorrect");

            verifyNoInteractions(jwtManager);
        }

        @Test
        @DisplayName("should check raw password against encoded password from DB")
        void shouldMatchRawPasswordAgainstEncodedPassword() {
            when(userRepository.findByEmail("alice@example.com")).thenReturn(Optional.of(alice));
            when(passwordEncoder.matches("password123", "hashed_password")).thenReturn(true);
            when(jwtManager.generateToken(anyString())).thenReturn("alice-token");
            when(userMapper.toResponse(alice)).thenReturn(buildUserResponse(alice));

            authService.signIn(request);

            verify(passwordEncoder).matches("password123", "hashed_password");
        }
    }
}