package io.github.agussyahrilmubarok.backend.infrastructure.security;

import io.github.agussyahrilmubarok.backend.common.exception.JwtExpiredTokenException;
import io.github.agussyahrilmubarok.backend.common.exception.JwtInvalidTokenException;
import io.jsonwebtoken.Claims;
import io.jsonwebtoken.ExpiredJwtException;
import io.jsonwebtoken.JwtException;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.security.Keys;
import java.nio.charset.StandardCharsets;
import java.util.Date;
import java.util.UUID;
import javax.crypto.SecretKey;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

@Component
public class JwtManager {

    private static final String ISSUER = "backend";
    private static final String CLAIM_USER_ID = "user_id";

    private final SecretKey signingKey;
    private final long expMillis;

    public JwtManager(
            @Value("${jwt.secret-key}") String secretKey, @Value("${jwt.exp-time-ms:86400000}") long expMillis) {
        this.signingKey = Keys.hmacShaKeyFor(secretKey.getBytes(StandardCharsets.UTF_8));
        this.expMillis = expMillis;
    }

    public String generateToken(String userId) {
        Date now = new Date();
        return Jwts.builder()
                .id(UUID.randomUUID().toString())
                .issuer(ISSUER)
                .claim(CLAIM_USER_ID, userId)
                .issuedAt(now)
                .expiration(new Date(now.getTime() + expMillis))
                .signWith(signingKey)
                .compact();
    }

    public String validateToken(String token) {
        try {
            Claims claims = Jwts.parser()
                    .requireIssuer(ISSUER)
                    .verifyWith(signingKey)
                    .build()
                    .parseSignedClaims(token)
                    .getPayload();

            return claims.get(CLAIM_USER_ID, String.class);
        } catch (ExpiredJwtException e) {
            throw new JwtExpiredTokenException();
        } catch (JwtException | IllegalArgumentException e) {
            throw new JwtInvalidTokenException();
        }
    }
}
