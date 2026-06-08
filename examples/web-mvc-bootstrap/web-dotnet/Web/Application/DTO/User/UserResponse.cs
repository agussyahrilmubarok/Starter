namespace Web.Application.DTO.User;

public record UserResponse(
    Guid Id,
    string Name,
    string Email,
    DateTime CreatedAt,
    DateTime UpdatedAt
);