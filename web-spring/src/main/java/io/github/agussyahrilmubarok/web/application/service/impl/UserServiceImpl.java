package io.github.agussyahrilmubarok.web.application.service.impl;

import io.github.agussyahrilmubarok.web.application.dto.user.CreateUserRequest;
import io.github.agussyahrilmubarok.web.application.dto.user.UpdateUserRequest;
import io.github.agussyahrilmubarok.web.application.dto.user.UserMapper;
import io.github.agussyahrilmubarok.web.application.dto.user.UserResponse;
import io.github.agussyahrilmubarok.web.application.service.UserService;
import io.github.agussyahrilmubarok.web.common.exception.EmailAlreadyExistsException;
import io.github.agussyahrilmubarok.web.common.exception.NotFoundException;
import io.github.agussyahrilmubarok.web.domain.User;
import io.github.agussyahrilmubarok.web.infrastructure.persistence.repository.UserRepository;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;

import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;
import java.util.Locale;
import java.util.UUID;

@Service
@RequiredArgsConstructor
@Slf4j
public class UserServiceImpl implements UserService {

    private final UserRepository userRepository;
    private final UserMapper userMapper;
    private final PasswordEncoder passwordEncoder;

    @Override
    @Transactional(readOnly = true)
    public List<UserResponse> getAll() {
        List<User> users = userRepository.findAll();
        log.info("Users fetched {}", users.size());
        return users.stream().map(userMapper::toResponse).toList();
    }

    @Override
    @Transactional(readOnly = true)
    public UserResponse getById(UUID id) {
        User user = userRepository.findById(id)
                .orElseThrow(() -> {
                    log.warn("User with id {} not found", id);
                    return new NotFoundException("User not found");
                });
        log.info("User fetched {}", user.getId());
        return userMapper.toResponse(user);
    }

    @Override
    @Transactional
    public UserResponse create(CreateUserRequest request) {
        if (userRepository.existsByEmail(request.email().toLowerCase(Locale.ROOT))) {
            log.warn("The email has already been taken");
            throw new EmailAlreadyExistsException();
        }

        User user = new User();
        user.setName(request.name());
        user.setEmail(request.email().toLowerCase(Locale.ROOT));
        user.setPassword(passwordEncoder.encode(request.password()));
        userRepository.save(user);

        log.info("User created {}", user.getId());
        return userMapper.toResponse(user);
    }

    @Override
    @Transactional
    public UserResponse update(UUID id, UpdateUserRequest request) {
        User user = userRepository.findById(id)
                .orElseThrow(() -> {
                    log.warn("User with id {} not found", id);
                    return new NotFoundException("User not found");
                });

        if (request.name() != null && !request.name().isBlank()) {
            user.setName(request.name());
        }

        if (request.email() != null && !request.email().isBlank() &&
                !user.getEmail().equals(request.email().toLowerCase(Locale.ROOT))) {
            if (userRepository.existsByEmail(request.email().toLowerCase(Locale.ROOT))) {
                log.warn("The email has already been taken {}", request.email());
                throw new EmailAlreadyExistsException();
            }
            user.setEmail(request.email().toLowerCase(Locale.ROOT));
        }

        if (request.password() != null && !request.password().isBlank()) {
            user.setPassword(passwordEncoder.encode(request.password()));
        }

        userRepository.save(user);

        log.info("User updated {}", user.getId());
        return userMapper.toResponse(user);
    }

    @Override
    @Transactional
    public void delete(UUID id) {
        User user = userRepository.findById(id)
                .orElseThrow(() -> {
                    log.warn("User with id {} not found", id);
                    return new NotFoundException("User not found");
                });

        userRepository.deleteById(user.getId());
        log.info("User deleted {}", user.getId());
    }
}
