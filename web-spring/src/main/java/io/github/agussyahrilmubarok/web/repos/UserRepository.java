package io.github.agussyahrilmubarok.web.repos;

import io.github.agussyahrilmubarok.web.domain.User;
import java.util.UUID;
import org.springframework.data.jpa.repository.JpaRepository;


public interface UserRepository extends JpaRepository<User, UUID> {

    boolean existsByEmailIgnoreCase(String email);

}
