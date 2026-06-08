using System.ComponentModel.DataAnnotations;

namespace Web.Application.DTO.Auth;

public class SignInRequest
{
    [Required(ErrorMessage = "Email is required")]
    [EmailAddress(ErrorMessage = "Email is not valid")]
    [StringLength(150)]
    public string Email { get; set; } = string.Empty;

    [Required(ErrorMessage = "Password is required")]
    [StringLength(72, MinimumLength = 8, ErrorMessage = "Password must be between 8 and 72 characters")]
    public string Password { get; set; } = string.Empty;
}
