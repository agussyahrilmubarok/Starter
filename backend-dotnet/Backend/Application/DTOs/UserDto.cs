using System.ComponentModel.DataAnnotations;
using System.Text.Json.Serialization;

namespace Backend.Application.DTOs;

public sealed record CreateUserRequest(
    [property: JsonPropertyName("name")]
    [Required]
    [MinLength(2)]
    [MaxLength(100)]
    string Name,
    [property: JsonPropertyName("email")]
    [Required]
    [EmailAddress]
    [MaxLength(150)]
    string Email,
    [property: JsonPropertyName("password")]
    [Required]
    [MinLength(8)]
    [MaxLength(72)]
    string Password
);

public sealed record UpdateUserRequest(
    [property: JsonPropertyName("name")]
    [MinLength(2)]
    [MaxLength(100)]
    string? Name,
    [property: JsonPropertyName("email")]
    [EmailAddress]
    [MaxLength(150)]
    string? Email,
    [property: JsonPropertyName("password")]
    [MinLength(8)]
    [MaxLength(72)]
    string? Password
);

public sealed record UserResponse(
    [property: JsonPropertyName("id")] string Id,
    [property: JsonPropertyName("name")] string Name,
    [property: JsonPropertyName("email")] string Email,
    [property: JsonPropertyName("created_at")]
    DateTimeOffset CreatedAt,
    [property: JsonPropertyName("updated_at")]
    DateTimeOffset UpdatedAt
);