package io.github.agussyahrilmubarok.web.controller;

import io.github.agussyahrilmubarok.web.model.auth.SignInRequest;
import io.github.agussyahrilmubarok.web.model.auth.SignUpRequest;
import io.github.agussyahrilmubarok.web.security.CustomUserDetails;
import io.github.agussyahrilmubarok.web.service.IAuthService;
import io.github.agussyahrilmubarok.web.util.EmailAlreadyExistsException;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpSession;
import jakarta.validation.Valid;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.BadCredentialsException;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.annotation.AuthenticationPrincipal;
import org.springframework.security.core.context.SecurityContext;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.web.context.HttpSessionSecurityContextRepository;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.validation.BindingResult;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.ModelAttribute;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.servlet.mvc.support.RedirectAttributes;

@Controller
public class AuthController {

    private static final Logger log = LoggerFactory.getLogger(AuthController.class);

    private final IAuthService authService;
    private final AuthenticationManager authenticationManager;

    public AuthController(IAuthService authService, AuthenticationManager authenticationManager) {
        this.authService = authService;
        this.authenticationManager = authenticationManager;
    }

    @GetMapping("/sign-up")
    public String signUpForm(
            @AuthenticationPrincipal final CustomUserDetails userDetails,
            @ModelAttribute("signUp") final SignUpRequest signUpRequest
    ) {
        if (userDetails != null) {
            return "redirect:/dashboard";
        }

        return "auth/sign-up";
    }

    @PostMapping("/sign-up")
    public String signUpSubmit(
            @ModelAttribute("signUp") @Valid final SignUpRequest signUpRequest,
            final BindingResult bindingResult,
            final Model model,
            final RedirectAttributes redirectAttributes
    ) {
        if (bindingResult.hasErrors()) {
            return "auth/sign-up";
        }

        try {
            authService.signUp(signUpRequest);
        } catch (final EmailAlreadyExistsException e) {
            bindingResult.rejectValue("email", "Exists.user.email");
            return "auth/sign-up";
        } catch (final Exception e) {
            model.addAttribute("MSG_ERROR", "Something went wrong. Please try again.");
            return "auth/sign-up";
        }

        redirectAttributes.addFlashAttribute("MSG_SUCCESS", "Sign up successfully! Please sign in.");
        return "redirect:/sign-in";
    }

    @GetMapping("/sign-in")
    public String signInForm(
            @AuthenticationPrincipal final CustomUserDetails userDetails,
            @ModelAttribute("signIn") final SignInRequest signInRequest
    ) {
        if (userDetails != null) {
            return "redirect:/dashboard";
        }

        return "auth/sign-in";
    }

    @PostMapping("/sign-in")
    public String signInSubmit(
            @ModelAttribute("signIn") @Valid final SignInRequest signInRequest,
            final BindingResult bindingResult,
            final HttpServletRequest request,
            final Model model
    ) {
        if (bindingResult.hasErrors()) {
            return "auth/sign-in";
        }

        try {
            UsernamePasswordAuthenticationToken authToken =
                    new UsernamePasswordAuthenticationToken(signInRequest.getEmail(), signInRequest.getPassword());

            Authentication authentication = authenticationManager.authenticate(authToken);

            SecurityContext context = SecurityContextHolder.createEmptyContext();
            context.setAuthentication(authentication);
            SecurityContextHolder.setContext(context);

            HttpSession session = request.getSession(true);
            session.setAttribute(HttpSessionSecurityContextRepository.SPRING_SECURITY_CONTEXT_KEY, context);

            log.info("Sign in successfully userEmail={}", signInRequest.getEmail());
        } catch (final BadCredentialsException e) {
            model.addAttribute("MSG_ERROR", "Invalid email or password.");
            return "auth/sign-in";
        } catch (final Exception e) {
            log.error("Authentication failed", e);
            model.addAttribute("MSG_ERROR", "Authentication failed. Please try again.");
            return "auth/sign-in";
        }

        return "redirect:/dashboard";
    }

    @PostMapping("/sign-out")
    public String signOut(
            final HttpServletRequest request,
            final RedirectAttributes redirectAttributes
    ) {
        SecurityContextHolder.clearContext();

        HttpSession session = request.getSession(false);
        if (session != null) {
            session.invalidate();
        }

        redirectAttributes.addFlashAttribute("MSG_SUCCESS", "You have signed out.");
        return "redirect:/sign-in";
    }
}