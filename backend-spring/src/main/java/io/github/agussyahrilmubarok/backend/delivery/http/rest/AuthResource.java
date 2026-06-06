package io.github.agussyahrilmubarok.backend.delivery.http.rest;

import io.github.agussyahrilmubarok.backend.application.dto.auth.AuthResponse;
import io.github.agussyahrilmubarok.backend.application.dto.auth.SignInRequest;
import io.github.agussyahrilmubarok.backend.application.dto.auth.SignUpRequest;
import io.github.agussyahrilmubarok.backend.application.service.AuthService;
import io.github.agussyahrilmubarok.backend.delivery.http.model.ApiResponse;
import io.swagger.v3.oas.annotations.security.SecurityRequirements;
import jakarta.validation.Valid;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping(value = "/api/v2/auth", produces = MediaType.APPLICATION_JSON_VALUE)
public class AuthResource {

    private final AuthService authService;

    public AuthResource(final AuthService authService) {
        this.authService = authService;
    }

    @PostMapping("/sign-up")
    @SecurityRequirements
    public ResponseEntity<ApiResponse<AuthResponse>> signUp(@Valid @RequestBody SignUpRequest request) {
        AuthResponse response = authService.signUp(request);
        return ResponseEntity.status(HttpStatus.CREATED).body(new ApiResponse<>("Signed up successfully", response));
    }

    @PostMapping("/sign-in")
    @SecurityRequirements
    public ResponseEntity<ApiResponse<AuthResponse>> signIn(@Valid @RequestBody SignInRequest request) {
        AuthResponse response = authService.signIn(request);
        return ResponseEntity.ok().body(new ApiResponse<>("Signed in successfully", response));
    }
}
