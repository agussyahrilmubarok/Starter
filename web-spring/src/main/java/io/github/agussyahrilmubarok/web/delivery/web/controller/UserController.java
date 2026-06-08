package io.github.agussyahrilmubarok.web.delivery.web.controller;

import io.github.agussyahrilmubarok.web.application.dto.user.CreateUserRequest;
import io.github.agussyahrilmubarok.web.application.dto.user.UpdateUserRequest;
import io.github.agussyahrilmubarok.web.application.dto.user.UserResponse;
import io.github.agussyahrilmubarok.web.application.service.UserService;
import io.github.agussyahrilmubarok.web.common.exception.EmailAlreadyInUseException;
import io.github.agussyahrilmubarok.web.common.exception.NotFoundException;
import io.github.agussyahrilmubarok.web.common.WebUtils;
import io.github.agussyahrilmubarok.web.infrastructure.security.CustomUserDetails;
import jakarta.validation.Valid;
import java.util.UUID;
import lombok.RequiredArgsConstructor;
import org.springframework.security.core.annotation.AuthenticationPrincipal;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.validation.BindingResult;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.servlet.mvc.support.RedirectAttributes;

@Controller
@RequestMapping("/dashboard/users")
@RequiredArgsConstructor
public class UserController {

    private final UserService userService;

    @GetMapping
    public String list(final Model model) {
        model.addAttribute("users", userService.getAll());
        return "dashboard/users/index";
    }

    @GetMapping("/create")
    public String createForm(@ModelAttribute("createUser") final CreateUserRequest createUserRequest) {
        return "dashboard/users/create";
    }

    @PostMapping("/create")
    public String createSubmit(
            @ModelAttribute("createUser") @Valid final CreateUserRequest createUserRequest,
            final BindingResult bindingResult,
            final Model model,
            final RedirectAttributes redirectAttributes) {
        if (bindingResult.hasErrors()) return "dashboard/users/create";

        try {
            userService.create(createUserRequest);
        } catch (final EmailAlreadyInUseException e) {
            bindingResult.rejectValue("email", "Exists.user.email");
            return "dashboard/users/create";
        } catch (final Exception e) {
            model.addAttribute(WebUtils.MSG_ERROR, "Something went wrong. Please try again");
            return "dashboard/users/create";
        }

        redirectAttributes.addFlashAttribute(WebUtils.MSG_SUCCESS, "User created successfully");
        return "redirect:/dashboard/users";
    }

    @GetMapping("/edit/{id}")
    public String editForm(
            @PathVariable final UUID id, final Model model, final RedirectAttributes redirectAttributes) {
        try {
            UserResponse user = userService.getById(id);
            UpdateUserRequest editUser = new UpdateUserRequest(user.id(), user.name(), user.email(), null);
            model.addAttribute("editUser", editUser);
        } catch (final NotFoundException e) {
            redirectAttributes.addFlashAttribute(WebUtils.MSG_ERROR, "User not found");
            return "redirect:/dashboard/users";
        }
        return "dashboard/users/edit";
    }

    @PostMapping("/edit/{id}")
    public String editSubmit(
            @PathVariable final UUID id,
            @ModelAttribute("editUser") @Valid final UpdateUserRequest request,
            final BindingResult bindingResult,
            final Model model,
            final RedirectAttributes redirectAttributes) {
        if (bindingResult.hasErrors()) return "dashboard/users/edit";

        try {
            userService.update(id, request);
        } catch (final EmailAlreadyInUseException e) {
            model.addAttribute("editUser", request);
            bindingResult.rejectValue("email", "Exists.user.email", "Email is already in use");
            return "dashboard/users/edit";
        } catch (final NotFoundException e) {
            redirectAttributes.addFlashAttribute(WebUtils.MSG_ERROR, "User not found");
            return "redirect:/dashboard/users";
        } catch (final Exception e) {
            model.addAttribute("editUser", request);
            model.addAttribute(WebUtils.MSG_ERROR, "Something went wrong. Please try again");
            return "dashboard/users/edit";
        }

        redirectAttributes.addFlashAttribute(WebUtils.MSG_SUCCESS, "User updated successfully");
        return "redirect:/dashboard/users";
    }

    @PostMapping("/delete/{id}")
    public String delete(@PathVariable final UUID id, final RedirectAttributes redirectAttributes) {
        try {
            userService.delete(id);
            redirectAttributes.addFlashAttribute(WebUtils.MSG_SUCCESS, "User deleted successfully");
        } catch (final NotFoundException e) {
            redirectAttributes.addFlashAttribute(WebUtils.MSG_ERROR, "User not found");
        }
        return "redirect:/dashboard/users";
    }

    @ModelAttribute
    public void globalAttributes(@AuthenticationPrincipal final CustomUserDetails userDetails, final Model model) {
        if (userDetails != null) {
            model.addAttribute("userProfile", userDetails.getUser());
        }
    }
}
