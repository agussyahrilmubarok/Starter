using System.ComponentModel.DataAnnotations;

namespace Backend.Application.DTOs;

public sealed record CreateUserRequest(
    [Required, MaxLength(100)] string Name,
    [Required, EmailAddress, MaxLength(150)] string Email,
    [Required, MinLength(8), MaxLength(72)] string Password
);

public sealed record UpdateUserRequest(
    [Required, MaxLength(100)] string Name,
    [Required, EmailAddress, MaxLength(150)] string Email
);

public sealed record UserResponse(
    Guid Id,
    string Name,
    string Email,
    DateTimeOffset CreatedAt,
    DateTimeOffset UpdatedAt
);