package io.github.agussyahrilmubarok.backend.service.impl;

import io.github.agussyahrilmubarok.backend.domain.User;
import io.github.agussyahrilmubarok.backend.model.user.CreateUserRequest;
import io.github.agussyahrilmubarok.backend.model.user.UpdateUserRequest;
import io.github.agussyahrilmubarok.backend.model.user.UserMapper;
import io.github.agussyahrilmubarok.backend.model.user.UserResponse;
import io.github.agussyahrilmubarok.backend.repository.UserRepository;
import io.github.agussyahrilmubarok.backend.service.IUserService;
import lombok.RequiredArgsConstructor;
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

    private final UserRepository userRepository;
    private final PasswordEncoder passwordEncoder;
    private final UserMapper userMapper;

    @Override
    @Transactional(readOnly = true)
    public List<UserResponse> findAll() {
        return userRepository.findAll()
                .stream()
                .map(userMapper::toResponse)
                .toList();
    }

    @Override
    @Transactional(readOnly = true)
    public UserResponse findById(String userId) {
        User user = userRepository.findById(parseUUID(userId))
                .orElseThrow(() -> new ResponseStatusException(
                        HttpStatus.NOT_FOUND, "User not found"));

        return userMapper.toResponse(user);
    }

    @Override
    @Transactional
    public UserResponse create(CreateUserRequest request) {
        if (userRepository.existsByEmailIgnoreCase(request.email())) {
            throw new ResponseStatusException(HttpStatus.CONFLICT, "Email already registered");
        }

        User user = new User();
        user.setName(request.name().trim());
        user.setEmail(request.email().toLowerCase());
        user.setPassword(passwordEncoder.encode(request.password()));

        return userMapper.toResponse(userRepository.save(user));
    }

    @Override
    @Transactional
    public UserResponse updateById(String userId, UpdateUserRequest request) {
        User user = userRepository.findById(parseUUID(userId))
                .orElseThrow(() -> new ResponseStatusException(
                        HttpStatus.NOT_FOUND, "User not found"));

        if (request.name() != null) {
            user.setName(request.name().trim());
        }

        if (request.email() != null) {
            boolean emailTaken = userRepository.existsByEmailIgnoreCase(request.email());
            if (emailTaken) {
                throw new ResponseStatusException(HttpStatus.CONFLICT, "Email already registered");
            }
            user.setEmail(request.email().toLowerCase());
        }

        if (request.password() != null) {
            user.setPassword(passwordEncoder.encode(request.password()));
        }

        return userMapper.toResponse(userRepository.save(user));
    }

    @Override
    @Transactional
    public void deleteById(String userId) {
        if (!userRepository.existsById(parseUUID(userId))) {
            throw new ResponseStatusException(HttpStatus.NOT_FOUND, "User not found");
        }

        userRepository.deleteById(parseUUID(userId));
    }

    private UUID parseUUID(String userId) {
        try {
            return UUID.fromString(userId);
        } catch (IllegalArgumentException e) {
            throw new ResponseStatusException(HttpStatus.BAD_REQUEST, "Invalid user id format");
        }
    }
}