using System.Text.Json.Serialization;

namespace Backend.Application.DTO.User;

public record UserResponse(
    [property: JsonPropertyName("id")]
    Guid Id,

    [property: JsonPropertyName("name")]
    string Name,

    [property: JsonPropertyName("email")]
    string Email,

    [property: JsonPropertyName("created_at")]
    DateTime CreatedAt,

    [property: JsonPropertyName("updated_at")]
    DateTime UpdatedAt
);