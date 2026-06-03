package io.github.agussyahrilmubarok.backend.infrastructure.security;

import io.github.agussyahrilmubarok.backend.common.exception.JwtExpiredTokenException;
import io.github.agussyahrilmubarok.backend.common.exception.JwtInvalidTokenException;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.security.Keys;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;

import javax.crypto.SecretKey;
import java.nio.charset.StandardCharsets;
import java.util.Date;
import java.util.UUID;

import static org.assertj.core.api.Assertions.assertThat;
import static org.assertj.core.api.Assertions.assertThatThrownBy;

class JwtManagerTest {

    private static final String SECRET_KEY = "test-secret-key-that-is-long-enough-for-hmac256!!";
    private static final long EXP_MILLIS = 86_400_000L; // 24 jam

    private JwtManager jwtManager;

    @BeforeEach
    void setUp() {
        jwtManager = new JwtManager(SECRET_KEY, EXP_MILLIS);
    }

    @Nested
    @DisplayName("generateToken()")
    class GenerateToken {

        @Test
        @DisplayName("should return non-blank token")
        void shouldReturnNonBlankToken() {
            String userId = UUID.randomUUID().toString();

            String token = jwtManager.generateToken(userId);

            assertThat(token).isNotBlank();
        }

        @Test
        @DisplayName("should return different tokens for different users")
        void shouldReturnDifferentTokensForDifferentUsers() {
            String token1 = jwtManager.generateToken(UUID.randomUUID().toString());
            String token2 = jwtManager.generateToken(UUID.randomUUID().toString());

            assertThat(token1).isNotEqualTo(token2);
        }

        @Test
        @DisplayName("should return different tokens on consecutive calls for same user")
        void shouldReturnDifferentTokensOnConsecutiveCallsForSameUser() {
            String sameUserId = UUID.randomUUID().toString();

            String token1 = jwtManager.generateToken(sameUserId);
            String token2 = jwtManager.generateToken(sameUserId);

            assertThat(token1).isNotEqualTo(token2);
            assertThat(jwtManager.validateToken(token1)).isEqualTo(sameUserId);
            assertThat(jwtManager.validateToken(token2)).isEqualTo(sameUserId);
        }
    }

    @Nested
    @DisplayName("validateToken()")
    class ValidateToken {

        @Test
        @DisplayName("should return userId from valid token")
        void shouldReturnUserIdFromValidToken() {
            String userId = UUID.randomUUID().toString();
            String token = jwtManager.generateToken(userId);

            String result = jwtManager.validateToken(token);

            assertThat(result).isEqualTo(userId);
        }

        @Test
        @DisplayName("should throw JwtExpiredTokenException for expired token")
        void shouldThrowWhenTokenExpired() {
            JwtManager shortLivedManager = new JwtManager(SECRET_KEY, 1L);
            String token = shortLivedManager.generateToken(UUID.randomUUID().toString());

            try {
                Thread.sleep(10);
            } catch (InterruptedException ignored) {
            }

            assertThatThrownBy(() -> shortLivedManager.validateToken(token))
                    .isInstanceOf(JwtExpiredTokenException.class)
                    .hasMessage("Token has expired.");
        }

        @Test
        @DisplayName("should throw JwtInvalidTokenException for malformed token")
        void shouldThrowWhenTokenMalformed() {
            assertThatThrownBy(() -> jwtManager.validateToken("this.is.not.a.valid.jwt"))
                    .isInstanceOf(JwtInvalidTokenException.class)
                    .hasMessage("Invalid token.");
        }

        @Test
        @DisplayName("should throw JwtInvalidTokenException for empty token")
        void shouldThrowWhenTokenEmpty() {
            assertThatThrownBy(() -> jwtManager.validateToken(""))
                    .isInstanceOf(JwtInvalidTokenException.class);
        }

        @Test
        @DisplayName("should throw JwtInvalidTokenException when signed with different key")
        void shouldThrowWhenSignedWithDifferentKey() {
            JwtManager otherManager = new JwtManager("other-secret-key-that-is-long-enough!!", EXP_MILLIS);
            String tokenFromOtherKey = otherManager.generateToken(UUID.randomUUID().toString());

            assertThatThrownBy(() -> jwtManager.validateToken(tokenFromOtherKey))
                    .isInstanceOf(JwtInvalidTokenException.class);
        }

        @Test
        @DisplayName("should throw JwtInvalidTokenException when issuer is wrong")
        void shouldThrowWhenIssuerIsWrong() {
            SecretKey key = Keys.hmacShaKeyFor(SECRET_KEY.getBytes(StandardCharsets.UTF_8));
            Date now = new Date();

            String tokenWithWrongIssuer = Jwts.builder()
                    .issuer("wrong-issuer")
                    .claim("user_id", UUID.randomUUID().toString())
                    .issuedAt(now)
                    .expiration(new Date(now.getTime() + EXP_MILLIS))
                    .signWith(key)
                    .compact();

            assertThatThrownBy(() -> jwtManager.validateToken(tokenWithWrongIssuer))
                    .isInstanceOf(JwtInvalidTokenException.class);
        }

        @Test
        @DisplayName("should throw JwtInvalidTokenException when user_id claim is missing")
        void shouldThrowWhenUserIdClaimMissing() {
            SecretKey key = Keys.hmacShaKeyFor(SECRET_KEY.getBytes(StandardCharsets.UTF_8));
            Date now = new Date();

            String tokenWithoutUserId = Jwts.builder()
                    .issuer("backend")
                    .issuedAt(now)
                    .expiration(new Date(now.getTime() + EXP_MILLIS))
                    .signWith(key)
                    .compact();

            String result = jwtManager.validateToken(tokenWithoutUserId);

            assertThat(result).isNull();
        }
    }
}