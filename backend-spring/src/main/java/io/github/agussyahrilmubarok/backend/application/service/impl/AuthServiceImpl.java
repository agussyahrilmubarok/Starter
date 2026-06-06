package io.github.agussyahrilmubarok.backend.application.service.impl;

import io.github.agussyahrilmubarok.backend.application.dto.auth.AuthResponse;
import io.github.agussyahrilmubarok.backend.application.dto.auth.SignInRequest;
import io.github.agussyahrilmubarok.backend.application.dto.auth.SignUpRequest;
import io.github.agussyahrilmubarok.backend.application.dto.user.UserMapper;
import io.github.agussyahrilmubarok.backend.application.service.AuthService;
import io.github.agussyahrilmubarok.backend.common.exception.EmailAlreadyInUseException;
import io.github.agussyahrilmubarok.backend.common.exception.EmailNotRegisteredException;
import io.github.agussyahrilmubarok.backend.common.exception.PasswordMismatchException;
import io.github.agussyahrilmubarok.backend.domain.User;
import io.github.agussyahrilmubarok.backend.infrastructure.persistence.repository.UserRepository;
import io.github.agussyahrilmubarok.backend.infrastructure.security.JwtManager;
import java.util.Locale;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
@Slf4j
@RequiredArgsConstructor
public class AuthServiceImpl implements AuthService {

    private final UserRepository userRepository;
    private final UserMapper userMapper;
    private final PasswordEncoder passwordEncoder;
    private final JwtManager jwtManager;

    @Override
    @Transactional
    public AuthResponse signUp(SignUpRequest request) {
        String email = request.email().trim().toLowerCase(Locale.ROOT);
        if (userRepository.existsByEmail(email)) {
            log.warn("Email is already in use");
            throw new EmailAlreadyInUseException();
        }

        User user = new User();
        user.setName(request.name());
        user.setEmail(email);
        user.setPassword(passwordEncoder.encode(request.password()));
        userRepository.save(user);

        String token = jwtManager.generateToken(user.getId().toString());

        log.info("User signed up {}", user.getId());
        return new AuthResponse(token, userMapper.toResponse(user));
    }

    @Override
    @Transactional(readOnly = true)
    public AuthResponse signIn(SignInRequest request) {
        String email = request.email().trim().toLowerCase(Locale.ROOT);

        User user = userRepository.findByEmail(email).orElseThrow(EmailNotRegisteredException::new);

        if (!passwordEncoder.matches(request.password(), user.getPassword())) {
            log.warn("Invalid password attempt for user {}", user.getId());
            throw new PasswordMismatchException();
        }

        String token = jwtManager.generateToken(user.getId().toString());

        log.info("User signed in {}", user.getId());
        return new AuthResponse(token, userMapper.toResponse(user));
    }
}
