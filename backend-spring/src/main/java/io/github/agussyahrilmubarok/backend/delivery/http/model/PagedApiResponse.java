package io.github.agussyahrilmubarok.backend.delivery.http.model;

import com.fasterxml.jackson.annotation.JsonInclude;
import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
@JsonInclude(JsonInclude.Include.NON_NULL)
public class PagedApiResponse<T> {
    private String message;
    private T data;
    private PaginationResponse pagination;

    public PagedApiResponse(String message, T data, PaginationResponse pagination) {
        this.message = message;
        this.data = data;
        this.pagination = pagination;
    }
}
