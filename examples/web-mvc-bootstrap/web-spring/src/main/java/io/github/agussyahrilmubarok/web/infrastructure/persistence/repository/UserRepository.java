package io.github.agussyahrilmubarok.web.infrastructure.persistence.repository;

import io.github.agussyahrilmubarok.web.domain.User;
import java.util.Optional;
import java.util.UUID;
import org.springframework.data.jpa.repository.JpaRepository;

public interface UserRepository extends JpaRepository<User, UUID> {
    Optional<User> findByEmailIgnoreCase(String email);

    boolean existsByEmail(String email);
}
