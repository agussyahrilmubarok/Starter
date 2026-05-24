package io.github.agussyahrilmubarok.backend.rest;

import io.github.agussyahrilmubarok.backend.model.ApiResponse;
import io.github.agussyahrilmubarok.backend.model.auth.SignInRequest;
import io.github.agussyahrilmubarok.backend.model.auth.SignInResponse;
import io.github.agussyahrilmubarok.backend.model.auth.SignUpRequest;
import io.github.agussyahrilmubarok.backend.model.auth.SignUpResponse;
import io.github.agussyahrilmubarok.backend.service.IAuthService;
import jakarta.validation.Valid;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping(value = "/api/v1/auth", produces = MediaType.APPLICATION_JSON_VALUE)
public class AuthResource {

    private final IAuthService authService;

    public AuthResource(final IAuthService authService) {
        this.authService = authService;
    }

    @PostMapping("/sign-up")
    public ResponseEntity<ApiResponse<SignUpResponse>> signUp(
            @Valid @RequestBody SignUpRequest request) {
        SignUpResponse response = authService.signUp(request);
        return ResponseEntity.status(HttpStatus.CREATED)
                .body(new ApiResponse<>("Sign up successfully", response));
    }

    @PostMapping("/sign-in")
    public ResponseEntity<ApiResponse<SignInResponse>> signIn(
            @Valid @RequestBody SignInRequest request) {
        SignInResponse response = authService.signIn(request);
        return ResponseEntity.ok()
                .body(new ApiResponse<>("Sign in successfully", response));
    }
}