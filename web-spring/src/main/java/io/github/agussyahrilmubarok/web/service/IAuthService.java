package io.github.agussyahrilmubarok.web.service;

import io.github.agussyahrilmubarok.web.model.auth.SignUpRequest;
import io.github.agussyahrilmubarok.web.model.user.UserResponse;

public interface IAuthService {

    UserResponse signUp(final SignUpRequest request);
}
