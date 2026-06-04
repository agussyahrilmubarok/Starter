package io.github.agussyahrilmubarok.web.delivery.web.controller;

import io.github.agussyahrilmubarok.web.application.dto.auth.SignInRequest;
import io.github.agussyahrilmubarok.web.application.dto.user.CreateUserRequest;
import io.github.agussyahrilmubarok.web.application.service.UserService;
import io.github.agussyahrilmubarok.web.common.exception.EmailAlreadyExistsException;
import io.github.agussyahrilmubarok.web.common.util.WebUtils;
import io.github.agussyahrilmubarok.web.infrastructure.security.CustomUserDetails;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpSession;
import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
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
@RequiredArgsConstructor
public class AuthController {

    private final UserService userService;
    private final AuthenticationManager authenticationManager;

    @GetMapping("/sign-up")
    public String signUpForm(
            @AuthenticationPrincipal final CustomUserDetails userDetails,
            @ModelAttribute("signUp") final CreateUserRequest createUserRequest) {
        if (userDetails != null)
            return "redirect:/dashboard";
        return "auth/sign-up";
    }

    @PostMapping("/sign-up")
    public String signUpSubmit(
            @ModelAttribute("signUp") @Valid final CreateUserRequest createUserRequest,
            final BindingResult bindingResult,
            final Model model,
            final RedirectAttributes redirectAttributes) {
        if (bindingResult.hasErrors())
            return "auth/sign-up";

        try {
            userService.create(createUserRequest);
        } catch (final EmailAlreadyExistsException e) {
            bindingResult.rejectValue("email", "Exists.user.email");
            return "auth/sign-up";
        } catch (final Exception e) {
            model.addAttribute(WebUtils.MSG_ERROR, WebUtils.getMessage("auth.general.error"));
            return "auth/sign-up";
        }

        redirectAttributes.addFlashAttribute(WebUtils.MSG_SUCCESS,
                WebUtils.getMessage("auth.signUp.success"));
        return "redirect:/sign-in";
    }

    @GetMapping("/sign-in")
    public String signInForm(
            @AuthenticationPrincipal final CustomUserDetails userDetails,
            @ModelAttribute("signIn") final SignInRequest signInRequest) {
        if (userDetails != null)
            return "redirect:/dashboard";
        return "auth/sign-in";
    }

    @PostMapping("/sign-in")
    public String signInSubmit(
            @ModelAttribute("signIn") @Valid final SignInRequest signInRequest,
            final BindingResult bindingResult,
            final HttpServletRequest request,
            final Model model) {
        if (bindingResult.hasErrors())
            return "auth/sign-in";

        try {
            UsernamePasswordAuthenticationToken authToken = new UsernamePasswordAuthenticationToken(
                    signInRequest.email(), signInRequest.password());

            Authentication authentication = authenticationManager.authenticate(authToken);

            SecurityContext context = SecurityContextHolder.createEmptyContext();
            context.setAuthentication(authentication);
            SecurityContextHolder.setContext(context);

            HttpSession session = request.getSession(true);
            session.setAttribute(
                    HttpSessionSecurityContextRepository.SPRING_SECURITY_CONTEXT_KEY, context);

        } catch (final BadCredentialsException e) {
            model.addAttribute(WebUtils.MSG_ERROR, WebUtils.getMessage("auth.signIn.error"));
            return "auth/sign-in";
        } catch (final Exception e) {
            model.addAttribute(WebUtils.MSG_ERROR, WebUtils.getMessage("auth.general.error"));
            return "auth/sign-in";
        }

        return "redirect:/dashboard";
    }

    @PostMapping("/sign-out")
    public String signOut(
            final HttpServletRequest request,
            final RedirectAttributes redirectAttributes) {
        SecurityContextHolder.clearContext();

        HttpSession session = request.getSession(false);
        if (session != null)
            session.invalidate();

        redirectAttributes.addFlashAttribute(WebUtils.MSG_SUCCESS, WebUtils.getMessage("auth.signOut.success"));
        return "redirect:/sign-in";
    }
}