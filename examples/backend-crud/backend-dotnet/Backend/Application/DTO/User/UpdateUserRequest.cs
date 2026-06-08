using System.ComponentModel.DataAnnotations;
using System.Text.Json.Serialization;

namespace Backend.Application.DTO.User;

public record UpdateUserRequest(
    [property: JsonPropertyName("name")]
    string? Name,

    [property: JsonPropertyName("email")]
    [param: ValidEmailIfPresent]
    string? Email,

    [property: JsonPropertyName("password")]
    string? Password
);

public class ValidEmailIfPresentAttribute : ValidationAttribute
{
    protected override ValidationResult? IsValid(object? value, ValidationContext validationContext)
    {
        if (value is null || string.IsNullOrWhiteSpace(value.ToString()))
            return ValidationResult.Success;

        var email = value.ToString()!;
        var isValid = new EmailAddressAttribute().IsValid(email);

        return isValid
            ? ValidationResult.Success
            : new ValidationResult("Email is not valid");
    }
}