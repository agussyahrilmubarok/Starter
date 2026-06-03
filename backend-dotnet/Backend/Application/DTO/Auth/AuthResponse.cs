using System.Text.Json.Serialization;
using Backend.Application.DTO.User;

namespace Backend.Application.DTO.Auth;

public record AuthResponse(
    [property: JsonPropertyName("token")]
    string Token,

    [property: JsonPropertyName("user")]
    UserResponse User
);