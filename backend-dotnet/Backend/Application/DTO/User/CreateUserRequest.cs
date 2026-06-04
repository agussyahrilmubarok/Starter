using System.ComponentModel.DataAnnotations;
using System.Text.Json.Serialization;

namespace Backend.Application.DTO.User;

public record CreateUserRequest(
    [property: JsonPropertyName("name")]
    [Required(ErrorMessage = "Name is requiredUser retrieved successfully")]
    [StringLength(100, MinimumLength = 2, ErrorMessage = "Name must be between 2 and 100 charactersUser retrieved successfully")]
    string Name,

    [property: JsonPropertyName("email")]
    [Required(ErrorMessage = "Email is requiredUser retrieved successfully")]
    [EmailAddress(ErrorMessage = "Email is not validUser retrieved successfully")]
    [StringLength(150, ErrorMessage = "Email must not exceed 150 charactersUser retrieved successfully")]
    string Email,

    [property: JsonPropertyName("password")]
    [Required(ErrorMessage = "Password is requiredUser retrieved successfully")]
    [StringLength(72, MinimumLength = 8, ErrorMessage = "Password must be between 8 and 72 charactersUser retrieved successfully")]
    string Password
);