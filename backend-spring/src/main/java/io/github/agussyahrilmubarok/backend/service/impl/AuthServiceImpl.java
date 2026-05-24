package io.github.agussyahrilmubarok.backend.service.impl;

import io.github.agussyahrilmubarok.backend.domain.User;
import io.github.agussyahrilmubarok.backend.model.auth.SignInRequest;
import io.github.agussyahrilmubarok.backend.model.auth.SignInResponse;
import io.github.agussyahrilmubarok.backend.model.auth.SignUpRequest;
import io.github.agussyahrilmubarok.backend.model.auth.SignUpResponse;
import io.github.agussyahrilmubarok.backend.model.user.UserMapper;
import io.github.agussyahrilmubarok.backend.repository.UserRepository;
import io.github.agussyahrilmubarok.backend.security.JwtProvider;
import io.github.agussyahrilmubarok.backend.service.IAuthService;
import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.web.server.ResponseStatusException;

@Service
@RequiredArgsConstructor
public class AuthServiceImpl implements IAuthService {

    private final UserRepository userRepository;
    private final PasswordEncoder passwordEncoder;
    private final JwtProvider jwtProvider;
    private final UserMapper userMapper;

    @Override
    @Transactional
    public SignUpResponse signUp(SignUpRequest request) {
        if (userRepository.existsByEmailIgnoreCase(request.email())) {
            throw new ResponseStatusException(HttpStatus.CONFLICT, "Email already registered");
        }

        User user = new User();
        user.setName(request.name().trim());
        user.setEmail(request.email().toLowerCase());
        user.setPassword(passwordEncoder.encode(request.password()));

        User saved = userRepository.save(user);

        String token = jwtProvider.generateToken(saved.getId().toString());
        return new SignUpResponse(token, userMapper.toResponse(saved));
    }

    @Override
    @Transactional(readOnly = true)
    public SignInResponse signIn(SignInRequest request) {
        User user = userRepository.findByEmailIgnoreCase(request.email())
                .orElseThrow(() -> new ResponseStatusException(
                        HttpStatus.UNAUTHORIZED, "Invalid email or password"));

        if (!passwordEncoder.matches(request.password(), user.getPassword())) {
            throw new ResponseStatusException(HttpStatus.UNAUTHORIZED, "Invalid email or password");
        }

        String token = jwtProvider.generateToken(user.getId().toString());
        return new SignInResponse(token, userMapper.toResponse(user));
    }
}