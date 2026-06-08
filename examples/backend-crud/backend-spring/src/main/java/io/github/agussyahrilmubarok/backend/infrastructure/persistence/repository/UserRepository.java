package io.github.agussyahrilmubarok.backend.infrastructure.persistence.repository;

import io.github.agussyahrilmubarok.backend.domain.User;
import java.util.Optional;
import java.util.UUID;
import org.springframework.data.jpa.repository.JpaRepository;

public interface UserRepository extends JpaRepository<User, UUID> {
    Optional<User> findByEmail(String email);

    boolean existsByEmail(String email);
}
