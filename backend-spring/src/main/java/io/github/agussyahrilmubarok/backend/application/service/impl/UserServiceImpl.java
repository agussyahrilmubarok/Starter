package io.github.agussyahrilmubarok.backend.application.service.impl;

import io.github.agussyahrilmubarok.backend.application.dto.user.CreateUserRequest;
import io.github.agussyahrilmubarok.backend.application.dto.user.UpdateUserRequest;
import io.github.agussyahrilmubarok.backend.application.dto.user.UserMapper;
import io.github.agussyahrilmubarok.backend.application.dto.user.UserPageRequest;
import io.github.agussyahrilmubarok.backend.application.dto.user.UserResponse;
import io.github.agussyahrilmubarok.backend.application.service.UserService;
import io.github.agussyahrilmubarok.backend.common.exception.EmailAlreadyInUseException;
import io.github.agussyahrilmubarok.backend.common.exception.NotFoundException;
import io.github.agussyahrilmubarok.backend.domain.User;
import io.github.agussyahrilmubarok.backend.infrastructure.persistence.repository.UserRepository;
import io.github.agussyahrilmubarok.backend.infrastructure.persistence.specification.UserSpecification;
import java.util.Locale;
import java.util.UUID;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;
import org.springframework.data.domain.Pageable;
import org.springframework.data.domain.Sort;
import org.springframework.data.jpa.domain.Specification;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
@RequiredArgsConstructor
@Slf4j
public class UserServiceImpl implements UserService {

    private final UserRepository userRepository;
    private final UserMapper userMapper;
    private final PasswordEncoder passwordEncoder;

    @Override
    @Transactional(readOnly = true)
    public Page<UserResponse> getAll(UserPageRequest request) {
        Sort sort = parseSort(request.sort());
        Pageable pageable = PageRequest.of(request.page() - 1, request.size(), sort);
        Specification<User> spec = UserSpecification.hasSearch(request.search());
        log.info("Users fetched");
        return userRepository.findAll(spec, pageable).map(userMapper::toResponse);
    }

    @Override
    @Transactional(readOnly = true)
    public UserResponse getById(UUID id) {
        User user = userRepository.findById(id).orElseThrow(() -> {
            log.warn("User with id {} not found", id);
            return new NotFoundException("User not found");
        });
        log.info("User fetched {}", user.getId());
        return userMapper.toResponse(user);
    }

    @Override
    @Transactional
    public UserResponse create(CreateUserRequest request) {
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

        log.info("User created {}", user.getId());
        return userMapper.toResponse(user);
    }

    @Override
    @Transactional
    public UserResponse update(UUID id, UpdateUserRequest request) {
        User user = userRepository.findById(id).orElseThrow(() -> {
            log.warn("User with id {} not found", id);
            return new NotFoundException("User not found");
        });

        if (request.name() != null && !request.name().isBlank()) {
            user.setName(request.name());
        }

        if (request.email() != null
                && !request.email().isBlank()
                && !user.getEmail().equals(request.email().toLowerCase(Locale.ROOT))) {
            String email = request.email().trim().toLowerCase(Locale.ROOT);
            if (userRepository.existsByEmail(email)) {
                log.warn("Email is already in use");
                throw new EmailAlreadyInUseException();
            }
            user.setEmail(email);
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
        User user = userRepository.findById(id).orElseThrow(() -> {
            log.warn("User with id {} not found", id);
            return new NotFoundException("User not found");
        });
        userRepository.deleteById(user.getId());
        log.info("User deleted {}", user.getId());
    }

    private Sort parseSort(String sortParam) {
        if (sortParam == null || sortParam.isBlank()) {
            return Sort.by(Sort.Direction.DESC, "createdAt");
        }

        try {
            String[] parts = sortParam.split(",");
            String field = convertToCamelCase(parts[0].trim());
            Sort.Direction direction = parts.length > 1 && "asc".equalsIgnoreCase(parts[1].trim())
                    ? Sort.Direction.ASC
                    : Sort.Direction.DESC;
            return Sort.by(direction, field);
        } catch (Exception e) {
            return Sort.by(Sort.Direction.DESC, "createdAt");
        }
    }

    private String convertToCamelCase(String input) {
        if (input == null || input.isBlank()) return "createdAt";

        if (input.contains("_")) {
            StringBuilder result = new StringBuilder();
            String[] tokens = input.split("_");
            for (int i = 0; i < tokens.length; i++) {
                String token = tokens[i];
                if (i == 0) {
                    result.append(token.toLowerCase());
                } else {
                    result.append(Character.toUpperCase(token.charAt(0)))
                            .append(token.substring(1).toLowerCase());
                }
            }
            return result.toString();
        }

        return input.trim();
    }
}
