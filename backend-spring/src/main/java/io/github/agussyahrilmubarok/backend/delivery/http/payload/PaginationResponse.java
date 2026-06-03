package io.github.agussyahrilmubarok.backend.delivery.http.payload;

import com.fasterxml.jackson.annotation.JsonProperty;

public record PaginationResponse(
        @JsonProperty("current_page") int currentPage,
        @JsonProperty("page_size") int pageSize,
        @JsonProperty("total_items") long totalItems,
        @JsonProperty("total_pages") int totalPages,
        @JsonProperty("has_next") boolean hasNext,
        @JsonProperty("has_previous") boolean hasPrevious
) {
}