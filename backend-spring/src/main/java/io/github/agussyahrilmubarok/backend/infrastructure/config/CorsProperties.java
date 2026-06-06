package io.github.agussyahrilmubarok.backend.infrastructure.config;

import java.util.List;
import org.springframework.boot.context.properties.ConfigurationProperties;

@ConfigurationProperties(prefix = "cors")
public record CorsProperties(List<String> allowedOrigins) {

    public CorsProperties {
        if (allowedOrigins == null || allowedOrigins.isEmpty()) {
            throw new IllegalStateException("CORS config is missing. Please set CORS_ALLOWED_ORIGINS");
        }
    }
}
