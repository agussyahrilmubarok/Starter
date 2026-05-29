package io.github.agussyahrilmubarok.web.service.impl;

import io.github.agussyahrilmubarok.web.domain.User;
import io.github.agussyahrilmubarok.web.model.user.CreateUserRequest;
import io.github.agussyahrilmubarok.web.model.user.UpdateUserRequest;
import io.github.agussyahrilmubarok.web.model.user.UserMapper;
import io.github.agussyahrilmubarok.web.model.user.UserResponse;
import io.github.agussyahrilmubarok.web.repository.UserRepository;
import io.github.agussyahrilmubarok.web.service.IUserService;
import io.github.agussyahrilmubarok.web.util.ConflictException;
import io.github.agussyahrilmubarok.web.util.NotFoundException;
import lombok.RequiredArgsConstructor;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

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
    public UserResponse findById(final UUID userId) {
        User user = userRepository.findById(userId)
                .orElseThrow(() -> {
                    log.warn("User not found userId={}", userId);
                    return new NotFoundException("User not found");
                });
        log.debug("Fetched successfully userId={}", userId);
        return userMapper.toResponse(user);
    }

    @Override
    @Transactional
    public UserResponse create(final CreateUserRequest request) {
        if (userRepository.existsByEmailIgnoreCase(request.getEmail())) {
            log.warn("Email already registered email={}", request.getEmail());
            throw new ConflictException("email", "Email already registered");
        }

        User user = new User();
        user.setName(request.getName().trim());
        user.setEmail(request.getEmail().toLowerCase());
        user.setPassword(passwordEncoder.encode(request.getPassword()));

        User saved = userRepository.save(user);
        log.info("Created successfully userId={}", saved.getId());
        return userMapper.toResponse(saved);
    }

    @Override
    @Transactional
    public UserResponse updateById(final UUID userId, final UpdateUserRequest request) {
        User user = userRepository.findById(userId)
                .orElseThrow(() -> {
                    log.warn("Not found for update userId={}", userId);
                    return new NotFoundException("User not found");
                });

        if (request.getName() != null && !request.getName().isBlank()) {
            user.setName(request.getName().trim());
        }

        if (request.getEmail() != null && !request.getEmail().isBlank()
                && !request.getEmail().equalsIgnoreCase(user.getEmail())) {
            if (userRepository.existsByEmailIgnoreCase(request.getEmail().trim())) {
                log.warn("Email already registered on update email={} userId={}", request.getEmail(), userId);
                throw new ConflictException("email", "Email already registered");
            }
            user.setEmail(request.getEmail().trim().toLowerCase());
        }

        if (request.getPassword() != null && !request.getPassword().isBlank()) {
            user.setPassword(passwordEncoder.encode(request.getPassword()));
        }

        User saved = userRepository.save(user);
        log.info("Updated successfully userId={}", saved.getId());
        return userMapper.toResponse(saved);
    }

    @Override
    @Transactional
    public void deleteById(final UUID userId) {
        if (!userRepository.existsById(userId)) {
            log.warn("Not found for delete userId={}", userId);
            throw new NotFoundException("User not found");
        }
        userRepository.deleteById(userId);
        log.info("Deleted successfully userId={}", userId);
    }
}