package io.github.agussyahrilmubarok.backend.application.dto.user;

public record UserPageRequest(int page, int size, String sort, String search) {
    public UserPageRequest {
        if (page <= 0) page = 1;
        if (size <= 0) size = 10;
        if (size > 100) size = 100;
        if (sort == null || sort.isBlank()) sort = "created_at,desc";
    }
}
