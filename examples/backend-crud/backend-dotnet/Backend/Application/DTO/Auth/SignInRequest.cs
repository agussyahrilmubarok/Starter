using System.ComponentModel.DataAnnotations;
using System.Text.Json.Serialization;

namespace Backend.Application.DTO.Auth;

public record SignInRequest(
    [property: JsonPropertyName("email")]
    [Required(ErrorMessage = "Email is required")]
    [EmailAddress(ErrorMessage = "Email is not valid")]
    [StringLength(150, ErrorMessage = "Email must not exceed 150 characters")]
    string Email,

    [property: JsonPropertyName("password")]
    [Required(ErrorMessage = "Password is required")]
    [StringLength(72, MinimumLength = 8, ErrorMessage = "Password must be between 8 and 72 characters")]
    string Password
);