package io.github.agussyahrilmubarok.backend.delivery.http.model;

import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.Map;
import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
@JsonInclude(JsonInclude.Include.NON_NULL)
public class ApiResponse<T> {
    private String message;
    private T data;
    private Map<String, String> errors;

    public ApiResponse(String message, T data) {
        this.message = message;
        this.data = data;
    }

    public ApiResponse(String message, Map<String, String> errors) {
        this.message = message;
        this.errors = errors;
    }
}
