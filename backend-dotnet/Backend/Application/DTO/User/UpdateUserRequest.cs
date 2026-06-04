using System.ComponentModel.DataAnnotations;
using System.Text.Json.Serialization;

namespace Backend.Application.DTO.User;

public record UpdateUserRequest(
    [property: JsonPropertyName("name")]
    string? Name,

    [property: JsonPropertyName("email")]
    [EmailAddress(ErrorMessage = "Email is not validUser retrieved successfully")]
    string? Email,

    [property: JsonPropertyName("password")]
    string? Password
);