package io.github.agussyahrilmubarok.backend.infrastructure.config;

import io.swagger.v3.oas.models.Components;
import io.swagger.v3.oas.models.OpenAPI;
import io.swagger.v3.oas.models.info.Info;
import io.swagger.v3.oas.models.media.*;
import io.swagger.v3.oas.models.responses.ApiResponse;
import io.swagger.v3.oas.models.security.SecurityRequirement;
import io.swagger.v3.oas.models.security.SecurityScheme;
import org.springdoc.core.customizers.OperationCustomizer;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class SwaggerConfig {

    private static final String SECURITY_SCHEME_NAME = "bearerAuth";

    @Bean
    public OpenAPI openApiSpec() {
        return new OpenAPI()
                .info(new Info().title("API Documentation").version("v2").description("Spring Boot REST API"))
                .components(new Components()
                        .addSecuritySchemes(
                                SECURITY_SCHEME_NAME,
                                new SecurityScheme()
                                        .name(SECURITY_SCHEME_NAME)
                                        .type(SecurityScheme.Type.HTTP)
                                        .scheme("bearer")
                                        .bearerFormat("JWT"))
                        .addSchemas(
                                "ApiErrorResponse",
                                new ObjectSchema()
                                        .addProperty("message", new StringSchema().example("Something went wrong"))
                                        .addProperty(
                                                "errors",
                                                new ObjectSchema()
                                                        .additionalProperties(new StringSchema())
                                                        .example("{\"error\": \"detail message\"}")))
                        .addSchemas(
                                "ApiValidationErrorResponse",
                                new ObjectSchema()
                                        .addProperty("message", new StringSchema().example("Validation failed"))
                                        .addProperty(
                                                "errors",
                                                new ObjectSchema()
                                                        .additionalProperties(new StringSchema())
                                                        .example("{\"email\": \"Email is not valid\"}"))))
                .addSecurityItem(new SecurityRequirement().addList(SECURITY_SCHEME_NAME));
    }

    @Bean
    public OperationCustomizer operationCustomizer() {
        return (operation, handlerMethod) -> {
            var responses = operation.getResponses();

            responses.addApiResponse(
                    "400",
                    new ApiResponse()
                            .description("Validation failed")
                            .content(new Content()
                                    .addMediaType(
                                            "application/json",
                                            new MediaType()
                                                    .schema(new Schema<>().$ref("ApiValidationErrorResponse")))));

            responses.addApiResponse(
                    "401",
                    new ApiResponse()
                            .description("Unauthorized")
                            .content(new Content()
                                    .addMediaType(
                                            "application/json",
                                            new MediaType().schema(new Schema<>().$ref("ApiErrorResponse")))));

            responses.addApiResponse(
                    "404",
                    new ApiResponse()
                            .description("Not found")
                            .content(new Content()
                                    .addMediaType(
                                            "application/json",
                                            new MediaType().schema(new Schema<>().$ref("ApiErrorResponse")))));

            responses.addApiResponse(
                    "409",
                    new ApiResponse()
                            .description("Conflict")
                            .content(new Content()
                                    .addMediaType(
                                            "application/json",
                                            new MediaType().schema(new Schema<>().$ref("ApiErrorResponse")))));

            responses.addApiResponse(
                    "500",
                    new ApiResponse()
                            .description("Internal server error")
                            .content(new Content()
                                    .addMediaType(
                                            "application/json",
                                            new MediaType().schema(new Schema<>().$ref("ApiErrorResponse")))));

            return operation;
        };
    }
}
