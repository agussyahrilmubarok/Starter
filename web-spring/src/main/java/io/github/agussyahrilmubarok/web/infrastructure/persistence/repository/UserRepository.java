package io.github.agussyahrilmubarok.web.infrastructure.persistence.repository;

import io.github.agussyahrilmubarok.web.domain.User;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.Optional;
import java.util.UUID;

public interface UserRepository extends JpaRepository<User, UUID> {
    Optional<User> findByEmailIgnoreCase(String email);

    boolean existsByEmail(String email);
}