using System.ComponentModel.DataAnnotations;

namespace Web.Application.DTO.User;

public class UpdateUserRequest
{
    [StringLength(100, MinimumLength = 2, ErrorMessage = "Name must be between 2 and 100 characters")]
    public string? Name { get; set; }

    [ValidEmailIfPresent]
    [StringLength(150)]
    public string? Email { get; set; }

    [StringLength(72, MinimumLength = 8, ErrorMessage = "Password must be between 8 and 72 characters")]
    public string? Password { get; set; }

    public UpdateUserRequest()
    {
    }

    public UpdateUserRequest(string name, string email)
    {
        Name = name;
        Email = email;
    }
}

public class ValidEmailIfPresentAttribute : ValidationAttribute
{
    protected override ValidationResult? IsValid(object? value, ValidationContext validationContext)
    {
        if (value is null || string.IsNullOrWhiteSpace(value.ToString()))
            return ValidationResult.Success;

        var isValid = new EmailAddressAttribute().IsValid(value.ToString());
        return isValid
            ? ValidationResult.Success
            : new ValidationResult("Email is not valid");
    }
}