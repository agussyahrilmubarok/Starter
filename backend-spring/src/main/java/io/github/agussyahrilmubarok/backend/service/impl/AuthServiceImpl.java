package io.github.agussyahrilmubarok.backend.service.impl;

import io.github.agussyahrilmubarok.backend.domain.User;
import io.github.agussyahrilmubarok.backend.exception.ConflictException;
import io.github.agussyahrilmubarok.backend.exception.UnauthorizedException;
import io.github.agussyahrilmubarok.backend.model.auth.SignInRequest;
import io.github.agussyahrilmubarok.backend.model.auth.SignInResponse;
import io.github.agussyahrilmubarok.backend.model.auth.SignUpRequest;
import io.github.agussyahrilmubarok.backend.model.auth.SignUpResponse;
import io.github.agussyahrilmubarok.backend.model.user.UserMapper;
import io.github.agussyahrilmubarok.backend.repository.UserRepository;
import io.github.agussyahrilmubarok.backend.security.JwtProvider;
import io.github.agussyahrilmubarok.backend.service.IAuthService;
import lombok.RequiredArgsConstructor;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
@RequiredArgsConstructor
public class AuthServiceImpl implements IAuthService {

    private static final Logger log = LoggerFactory.getLogger(AuthServiceImpl.class);

    private final UserRepository userRepository;
    private final PasswordEncoder passwordEncoder;
    private final JwtProvider jwtProvider;
    private final UserMapper userMapper;

    @Override
    @Transactional
    public SignUpResponse signUp(SignUpRequest request) {
        log.info("sign-up: attempt email={}", request.email());

        if (userRepository.existsByEmailIgnoreCase(request.email())) {
            log.warn("sign-up: email already registered email={}", request.email());
            throw new ConflictException("email", "Email already registered");
        }

        User user = new User();
        user.setName(request.name().trim());
        user.setEmail(request.email().toLowerCase());
        user.setPassword(passwordEncoder.encode(request.password()));

        User saved = userRepository.save(user);
        log.info("sign-up: user created successfully userId={}", saved.getId());

        String token = jwtProvider.generateToken(saved.getId().toString());
        return new SignUpResponse(token, userMapper.toResponse(saved));
    }

    @Override
    @Transactional(readOnly = true)
    public SignInResponse signIn(SignInRequest request) {
        log.info("sign-in: attempt email={}", request.email());

        User user = userRepository.findByEmailIgnoreCase(request.email())
                .orElseThrow(() -> {
                    log.warn("sign-in: email not found email={}", request.email());
                    return new UnauthorizedException("email", "Email not found");
                });

        if (!passwordEncoder.matches(request.password(), user.getPassword())) {
            log.warn("sign-in: password not match userId={}", user.getId());
            throw new UnauthorizedException("password", "Password does not match");
        }

        String token = jwtProvider.generateToken(user.getId().toString());
        log.info("sign-in: success userId={}", user.getId());
        return new SignInResponse(token, userMapper.toResponse(user));
    }
}