package io.github.agussyahrilmubarok.web.controller;

import io.github.agussyahrilmubarok.web.model.user.CreateUserRequest;
import io.github.agussyahrilmubarok.web.model.user.UpdateUserRequest;
import io.github.agussyahrilmubarok.web.security.CustomUserDetails;
import io.github.agussyahrilmubarok.web.service.IUserService;
import io.github.agussyahrilmubarok.web.util.ConflictException;
import io.github.agussyahrilmubarok.web.util.NotFoundException;
import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.springframework.security.core.annotation.AuthenticationPrincipal;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.validation.BindingResult;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.servlet.mvc.support.RedirectAttributes;

import java.util.UUID;

@Controller
@RequestMapping("/dashboard/users")
@RequiredArgsConstructor
public class UserController {

    private final IUserService userService;

    @GetMapping
    public String list(final Model model) {
        model.addAttribute("users", userService.findAll());
        return "dashboard/users/index";
    }

    @GetMapping("/create")
    public String createForm(@ModelAttribute("createUser") final CreateUserRequest request) {
        return "dashboard/users/create";
    }

    @PostMapping("/create")
    public String createSubmit(
            @ModelAttribute("createUser") @Valid final CreateUserRequest request,
            final BindingResult bindingResult,
            final Model model,
            final RedirectAttributes redirectAttributes
    ) {
        if (bindingResult.hasErrors()) {
            return "dashboard/users/create";
        }

        try {
            userService.create(request);
        } catch (final ConflictException e) {
            bindingResult.rejectValue(e.getField(), "Exists.user.email");
            return "dashboard/users/create";
        } catch (final Exception e) {
            model.addAttribute("MSG_ERROR", "Something went wrong. Please try again.");
            return "dashboard/users/create";
        }

        redirectAttributes.addFlashAttribute("MSG_SUCCESS", "User was created successfully.");
        return "redirect:/dashboard/users";
    }

    @GetMapping("/edit/{id}")
    public String editForm(
            @PathVariable final UUID id,
            final Model model,
            final RedirectAttributes redirectAttributes
    ) {
        try {
            model.addAttribute("editUser", userService.findById(id));
        } catch (final NotFoundException e) {
            redirectAttributes.addFlashAttribute("MSG_ERROR", "User not found.");
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
            final RedirectAttributes redirectAttributes
    ) {
        if (bindingResult.hasErrors()) {
            return "dashboard/users/edit";
        }

        try {
            userService.updateById(id, request);
        } catch (final ConflictException e) {
            bindingResult.rejectValue(e.getField(), "Exists.user.email");
            return "dashboard/users/edit";
        } catch (final NotFoundException e) {
            redirectAttributes.addFlashAttribute("MSG_ERROR", "User not found.");
            return "redirect:/dashboard/users";
        } catch (final Exception e) {
            model.addAttribute("MSG_ERROR", "Something went wrong. Please try again.");
            return "dashboard/users/edit";
        }

        redirectAttributes.addFlashAttribute("MSG_SUCCESS", "User was updated successfully.");
        return "redirect:/dashboard/users";
    }

    @PostMapping("/delete/{id}")
    public String delete(
            @PathVariable final UUID id,
            final RedirectAttributes redirectAttributes
    ) {
        try {
            userService.deleteById(id);
            redirectAttributes.addFlashAttribute("MSG_SUCCESS", "User was deleted successfully.");
        } catch (final NotFoundException e) {
            redirectAttributes.addFlashAttribute("MSG_ERROR", "User not found.");
        }
        return "redirect:/dashboard/users";
    }

    @ModelAttribute
    public void globalAttributes(
            @AuthenticationPrincipal final CustomUserDetails userDetails,
            final Model model
    ) {
        if (userDetails != null) {
            model.addAttribute("userProfile", userDetails.getUser());
        }
    }
}