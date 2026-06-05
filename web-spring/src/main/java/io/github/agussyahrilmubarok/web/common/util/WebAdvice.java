package io.github.agussyahrilmubarok.web.common.util;

import io.github.agussyahrilmubarok.web.common.exception.NotFoundException;
import jakarta.servlet.http.HttpServletRequest;
import org.springframework.http.HttpStatus;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.ModelAttribute;
import org.springframework.web.bind.annotation.ResponseStatus;

@ControllerAdvice
public class WebAdvice {

    @ModelAttribute("requestUri")
    public String getRequestUri(final HttpServletRequest request) {
        return request.getRequestURI();
    }

    @ExceptionHandler(NotFoundException.class)
    @ResponseStatus(HttpStatus.NOT_FOUND)
    public String handleNotFound(final NotFoundException e, final Model model) {
        model.addAttribute("message", e.getMessage());
        return "error/404";
    }
}