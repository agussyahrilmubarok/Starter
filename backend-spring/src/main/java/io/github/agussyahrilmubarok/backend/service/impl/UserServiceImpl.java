package io.github.agussyahrilmubarok.backend.service.impl;

import io.github.agussyahrilmubarok.backend.domain.User;
import io.github.agussyahrilmubarok.backend.exception.ConflictException;
import io.github.agussyahrilmubarok.backend.exception.NotFoundException;
import io.github.agussyahrilmubarok.backend.model.user.CreateUserRequest;
import io.github.agussyahrilmubarok.backend.model.user.UpdateUserRequest;
import io.github.agussyahrilmubarok.backend.model.user.UserMapper;
import io.github.agussyahrilmubarok.backend.model.user.UserResponse;
import io.github.agussyahrilmubarok.backend.repository.UserRepository;
import io.github.agussyahrilmubarok.backend.service.IUserService;
import lombok.RequiredArgsConstructor;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.http.HttpStatus;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.web.server.ResponseStatusException;

import java.util.List;
import java.util.UUID;

@Service
@RequiredArgsConstructor
public class UserServiceImpl implements IUserService {

    private static final Logger log = LoggerFactory.getLogger(UserServiceImpl.class);

    private final UserRepository userRepository;
    private final PasswordEncoder passwordEncoder;
    private final UserMapper userMapper;

    @Override
    @Transactional(readOnly = true)
    public List<UserResponse> findAll() {
        List<UserResponse> result = userRepository.findAll()
                .stream()
                .map(userMapper::toResponse)
                .toList();

        log.debug("Fetched all users count={}", result.size());
        return result;
    }

    @Override
    @Transactional(readOnly = true)
    public UserResponse findById(String userId) {
        User user = userRepository.findById(parseUUID(userId))
                .orElseThrow(() -> {
                    log.warn("User not found userId={}", userId);
                    return new NotFoundException("User not found");
                });

        log.debug("Fetched successfully userId={}", userId);
        return userMapper.toResponse(user);
    }

    @Override
    @Transactional
    public UserResponse create(CreateUserRequest request) {
        if (userRepository.existsByEmailIgnoreCase(request.email())) {
            log.warn("Email already registered email={}", request.email());
            throw new ConflictException("email", "Email already registered");
        }

        User user = new User();
        user.setName(request.name().trim());
        user.setEmail(request.email().toLowerCase());
        user.setPassword(passwordEncoder.encode(request.password()));

        User saved = userRepository.save(user);
        log.info("Created successfully userId={}", saved.getId());
        return userMapper.toResponse(saved);
    }

    @Override
    @Transactional
    public UserResponse updateById(String userId, UpdateUserRequest request) {
        User user = userRepository.findById(parseUUID(userId))
                .orElseThrow(() -> {
                    log.warn("Not found for update userId={}", userId);
                    return new NotFoundException("User not found");
                });

        if (request.name() != null && !request.name().isBlank()) {
            user.setName(request.name().trim());
        }

        if (request.email() != null && !request.email().isBlank() && !request.email().equalsIgnoreCase(user.getEmail())) {
            boolean emailTaken = userRepository.existsByEmailIgnoreCase(request.email().trim());
            if (emailTaken) {
                log.warn("Email already registered on update email={} userId={}", request.email(), userId);
                throw new ConflictException("email", "Email already registered");
            }
            user.setEmail(request.email().trim().toLowerCase());
        }

        if (request.password() != null && !request.password().isBlank()) {
            user.setPassword(passwordEncoder.encode(request.password()));
        }

        User saved = userRepository.save(user);
        log.info("Updated successfully userId={}", saved.getId());
        return userMapper.toResponse(saved);
    }

    @Override
    @Transactional
    public void deleteById(String userId) {
        if (!userRepository.existsById(parseUUID(userId))) {
            log.warn("Not found for delete userId={}", userId);
            throw new ResponseStatusException(HttpStatus.NOT_FOUND, "User not found");
        }

        userRepository.deleteById(parseUUID(userId));
        log.info("Deleted successfully userId={}", userId);
    }

    private UUID parseUUID(String userId) {
        try {
            return UUID.fromString(userId);
        } catch (IllegalArgumentException e) {
            log.warn("Invalid UUID format userId={}", userId);
            throw new ResponseStatusException(HttpStatus.BAD_REQUEST, "Invalid user id format");
        }
    }
}