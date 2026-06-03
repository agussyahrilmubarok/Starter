package io.github.agussyahrilmubarok.backend.application.service;

import io.github.agussyahrilmubarok.backend.application.dto.auth.AuthResponse;
import io.github.agussyahrilmubarok.backend.application.dto.auth.SignInRequest;
import io.github.agussyahrilmubarok.backend.application.dto.auth.SignUpRequest;

public interface AuthService {
    AuthResponse signUp(final SignUpRequest request);

    AuthResponse signIn(final SignInRequest request);
}
