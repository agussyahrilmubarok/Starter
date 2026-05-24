package io.github.agussyahrilmubarok.backend.service;

import io.github.agussyahrilmubarok.backend.model.auth.SignInRequest;
import io.github.agussyahrilmubarok.backend.model.auth.SignInResponse;
import io.github.agussyahrilmubarok.backend.model.auth.SignUpRequest;
import io.github.agussyahrilmubarok.backend.model.auth.SignUpResponse;

public interface IAuthService {
    SignUpResponse signUp(SignUpRequest request);

    SignInResponse signIn(SignInRequest request);
}
