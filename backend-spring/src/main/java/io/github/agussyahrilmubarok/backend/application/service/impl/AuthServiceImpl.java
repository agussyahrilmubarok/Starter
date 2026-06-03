package io.github.agussyahrilmubarok.backend.application.service.impl;

import io.github.agussyahrilmubarok.backend.application.dto.auth.AuthResponse;
import io.github.agussyahrilmubarok.backend.application.dto.auth.SignInRequest;
import io.github.agussyahrilmubarok.backend.application.dto.auth.SignUpRequest;
import io.github.agussyahrilmubarok.backend.application.dto.user.UserMapper;
import io.github.agussyahrilmubarok.backend.application.service.AuthService;
import io.github.agussyahrilmubarok.backend.common.exception.EmailAlreadyExistsException;
import io.github.agussyahrilmubarok.backend.common.exception.EmailNotRegisteredException;
import io.github.agussyahrilmubarok.backend.common.exception.WrongPasswordException;
import io.github.agussyahrilmubarok.backend.domain.User;
import io.github.agussyahrilmubarok.backend.infrastructure.persistence.repository.UserRepository;
import io.github.agussyahrilmubarok.backend.infrastructure.security.JwtManager;
import lombok.RequiredArgsConstructor;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
@RequiredArgsConstructor
public class AuthServiceImpl implements AuthService {

    private final UserRepository userRepository;
    private final UserMapper userMapper;
    private final PasswordEncoder passwordEncoder;
    private final JwtManager jwtManager;

    @Override
    @Transactional
    public AuthResponse signUp(SignUpRequest request) {
        if (userRepository.existsByEmail(request.email())) {
            throw new EmailAlreadyExistsException();
        }

        User user = new User();
        user.setName(request.name());
        user.setEmail(request.email().toLowerCase());
        user.setPassword(passwordEncoder.encode(request.password()));
        userRepository.save(user);

        String token = jwtManager.generateToken(user.getId().toString());

        return toResponse(token, user);
    }

    @Override
    @Transactional(readOnly = true)
    public AuthResponse signIn(SignInRequest request) {
        User user = userRepository.findByEmail(request.email())
                .orElseThrow(EmailNotRegisteredException::new);

        if (!passwordEncoder.matches(request.password(), user.getPassword())) {
            throw new WrongPasswordException();
        }

        String token = jwtManager.generateToken(user.getId().toString());

        return toResponse(token, user);
    }

    private AuthResponse toResponse(String token, User user) {
        return new AuthResponse(token, userMapper.toResponse(user));
    }
}
