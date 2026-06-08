package io.github.agussyahrilmubarok.web.infrastructure.security;

import io.github.agussyahrilmubarok.web.domain.User;
import io.github.agussyahrilmubarok.web.infrastructure.persistence.repository.UserRepository;
import lombok.RequiredArgsConstructor;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.stereotype.Component;

/**
 * Loads users from the database for Spring Security authentication.
 * Maps our User entity to Spring Security's UserDetails.
 */
@Component
@RequiredArgsConstructor
public class CustomUserDetailsService implements UserDetailsService {

    private final UserRepository userRepository;

    @Override
    public UserDetails loadUserByUsername(final String email) throws UsernameNotFoundException {
        final User user = userRepository
                .findByEmailIgnoreCase(email)
                .orElseThrow(() -> new UsernameNotFoundException("User not found: " + email));

        return new CustomUserDetails(user);
    }
}
