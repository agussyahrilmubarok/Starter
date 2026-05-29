package io.github.agussyahrilmubarok.web.service.impl;

import io.github.agussyahrilmubarok.web.domain.User;
import io.github.agussyahrilmubarok.web.model.auth.SignUpRequest;
import io.github.agussyahrilmubarok.web.model.user.UserMapper;
import io.github.agussyahrilmubarok.web.model.user.UserResponse;
import io.github.agussyahrilmubarok.web.repository.UserRepository;
import io.github.agussyahrilmubarok.web.service.IAuthService;
import io.github.agussyahrilmubarok.web.util.EmailAlreadyExistsException;
import lombok.RequiredArgsConstructor;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
public class AuthServiceImpl implements IAuthService {

    private static final Logger log = LoggerFactory.getLogger(AuthServiceImpl.class);

    private final UserRepository userRepository;
    private final PasswordEncoder passwordEncoder;
    private final UserMapper userMapper;

    @Override
    public UserResponse signUp(SignUpRequest request) {
        if (userRepository.existsByEmailIgnoreCase(request.getEmail())) {
            log.warn("Email already exists userEmail={}", request.getEmail());
            throw new EmailAlreadyExistsException();
        }

        User user = new User();
        user.setName(request.getName().trim());
        user.setEmail(request.getEmail().toLowerCase());
        user.setPassword(passwordEncoder.encode(request.getPassword()));

        User saved = userRepository.save(user);
        log.info("Sign up user successfully userId={}", saved.getId());
        return userMapper.toResponse(saved);
    }
}