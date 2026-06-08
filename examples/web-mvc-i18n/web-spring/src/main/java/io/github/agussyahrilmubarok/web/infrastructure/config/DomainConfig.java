package io.github.agussyahrilmubarok.web.infrastructure.config;

import java.time.OffsetDateTime;
import java.util.Optional;
import org.springframework.boot.persistence.autoconfigure.EntityScan;
import org.springframework.context.MessageSource;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.data.auditing.DateTimeProvider;
import org.springframework.data.jpa.repository.config.EnableJpaAuditing;
import org.springframework.data.jpa.repository.config.EnableJpaRepositories;
import org.springframework.transaction.annotation.EnableTransactionManagement;
import org.springframework.validation.beanvalidation.LocalValidatorFactoryBean;

@Configuration
@EntityScan("io.github.agussyahrilmubarok.web.domain")
@EnableJpaRepositories("io.github.agussyahrilmubarok.web.infrastructure.persistence.repository")
@EnableTransactionManagement
@EnableJpaAuditing(dateTimeProviderRef = "auditingDateTimeProvider")
public class DomainConfig {

    @Bean(name = "auditingDateTimeProvider")
    public DateTimeProvider dateTimeProvider() {
        return () -> Optional.of(OffsetDateTime.now());
    }

    @Bean
    public LocalValidatorFactoryBean validator(MessageSource messageSource) {
        LocalValidatorFactoryBean factory = new LocalValidatorFactoryBean();
        factory.setValidationMessageSource(messageSource);
        return factory;
    }
}
