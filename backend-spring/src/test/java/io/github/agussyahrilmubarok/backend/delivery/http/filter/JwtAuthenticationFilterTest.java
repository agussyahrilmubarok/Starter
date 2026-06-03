package io.github.agussyahrilmubarok.backend.delivery.http.filter;

import io.github.agussyahrilmubarok.backend.common.exception.JwtExpiredTokenException;
import io.github.agussyahrilmubarok.backend.common.exception.JwtInvalidTokenException;
import io.github.agussyahrilmubarok.backend.infrastructure.security.JwtManager;
import jakarta.servlet.FilterChain;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;
import org.springframework.security.core.context.SecurityContextHolder;

import java.util.UUID;

import static org.assertj.core.api.Assertions.assertThat;
import static org.assertj.core.api.Assertions.assertThatThrownBy;
import static org.mockito.Mockito.*;

@ExtendWith(MockitoExtension.class)
class JwtAuthenticationFilterTest {

    @Mock
    private JwtManager jwtManager;

    @Mock
    private HttpServletRequest request;

    @Mock
    private HttpServletResponse response;

    @Mock
    private FilterChain filterChain;

    @InjectMocks
    private JwtAuthenticationFilter filter;

    @BeforeEach
    void setUp() {
        SecurityContextHolder.clearContext();
    }

    @AfterEach
    void tearDown() {
        SecurityContextHolder.clearContext();
    }

    @Test
    @DisplayName("should skip filter and pass through when Authorization header is missing")
    void shouldSkipWhenNoAuthHeader() throws Exception {
        when(request.getHeader("Authorization")).thenReturn(null);

        filter.doFilterInternal(request, response, filterChain);

        verify(filterChain).doFilter(request, response);
        verifyNoInteractions(jwtManager);
        assertThat(SecurityContextHolder.getContext().getAuthentication()).isNull();
    }

    @Test
    @DisplayName("should skip filter and pass through when Authorization header is not Bearer")
    void shouldSkipWhenNotBearerToken() throws Exception {
        when(request.getHeader("Authorization")).thenReturn("Basic alice:password123");

        filter.doFilterInternal(request, response, filterChain);

        verify(filterChain).doFilter(request, response);
        verifyNoInteractions(jwtManager);
        assertThat(SecurityContextHolder.getContext().getAuthentication()).isNull();
    }

    @Test
    @DisplayName("should set authentication in SecurityContext when token is valid")
    void shouldSetAuthenticationWhenTokenValid() throws Exception {
        String aliceId = UUID.randomUUID().toString();

        when(request.getHeader("Authorization")).thenReturn("Bearer alice-valid-token");
        when(jwtManager.validateToken("alice-valid-token")).thenReturn(aliceId);

        filter.doFilterInternal(request, response, filterChain);

        var auth = SecurityContextHolder.getContext().getAuthentication();
        assertThat(auth).isNotNull();
        assertThat(auth.getPrincipal()).isEqualTo(aliceId);
        assertThat(auth.getAuthorities()).anyMatch(a -> a.getAuthority().equals("ROLE_USER"));
        verify(filterChain).doFilter(request, response);
    }

    @Test
    @DisplayName("should throw JwtExpiredTokenException when token is expired")
    void shouldThrowWhenTokenExpired() {
        when(request.getHeader("Authorization")).thenReturn("Bearer alice-expired-token");
        when(jwtManager.validateToken("alice-expired-token")).thenThrow(new JwtExpiredTokenException());

        assertThatThrownBy(() -> filter.doFilterInternal(request, response, filterChain))
                .isInstanceOf(JwtExpiredTokenException.class)
                .hasMessage("Token has expired");

        assertThat(SecurityContextHolder.getContext().getAuthentication()).isNull();
    }

    @Test
    @DisplayName("should throw JwtInvalidTokenException when token is invalid")
    void shouldThrowWhenTokenInvalid() {
        when(request.getHeader("Authorization")).thenReturn("Bearer alice-invalid-token");
        when(jwtManager.validateToken("alice-invalid-token")).thenThrow(new JwtInvalidTokenException());

        assertThatThrownBy(() -> filter.doFilterInternal(request, response, filterChain))
                .isInstanceOf(JwtInvalidTokenException.class)
                .hasMessage("Invalid token");

        assertThat(SecurityContextHolder.getContext().getAuthentication()).isNull();
    }

    @Test
    @DisplayName("should throw JwtInvalidTokenException when unexpected exception occurs")
    void shouldWrapUnexpectedExceptionAsInvalidToken() {
        when(request.getHeader("Authorization")).thenReturn("Bearer alice-broken-token");
        when(jwtManager.validateToken("alice-broken-token")).thenThrow(new RuntimeException("unexpected"));

        assertThatThrownBy(() -> filter.doFilterInternal(request, response, filterChain))
                .isInstanceOf(JwtInvalidTokenException.class);
    }
}