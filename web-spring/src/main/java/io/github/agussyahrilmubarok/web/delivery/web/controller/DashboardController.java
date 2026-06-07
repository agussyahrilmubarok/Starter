package io.github.agussyahrilmubarok.web.delivery.web.controller;

import io.github.agussyahrilmubarok.web.infrastructure.security.CustomUserDetails;
import org.springframework.security.core.annotation.AuthenticationPrincipal;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.ModelAttribute;

@Controller
public class DashboardController {

    @GetMapping("/dashboard")
    public String dashboard() {
        return "dashboard/index";
    }

    @GetMapping("/dashboard/profile")
    public String profile() {
        return "dashboard/profile/index";
    }

    @ModelAttribute
    public void globalAttributes(@AuthenticationPrincipal final CustomUserDetails userDetails, final Model model) {
        if (userDetails != null) {
            model.addAttribute("userProfile", userDetails.getUser());
        }
    }
}
