using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.RazorPages;
using System.ComponentModel.DataAnnotations;
using Web.Domain.User;

namespace Web.Pages;

public class SignUpModel : PageModel
{
    private readonly IUserRepository _userRepository;

    public SignUpModel(IUserRepository userRepository)
    {
        _userRepository = userRepository;
    }

    [BindProperty]
    public SignUpInput Input { get; set; } = new();

    public string? ErrorMessage { get; set; }

    public void OnGet() { }

    public async Task<IActionResult> OnPostAsync(CancellationToken ct)
    {
        if (!ModelState.IsValid) return Page();

        if (await _userRepository.ExistsByEmailAsync(Input.Email.ToLower(), ct))
        {
            ModelState.AddModelError("Input.Email", "Email already exists.");
            return Page();
        }

        try
        {
            var hashedPassword = BCrypt.Net.BCrypt.HashPassword(Input.Password);
            var user = Domain.User.User.Create(Input.Name, Input.Email.ToLowerInvariant(), hashedPassword);
            await _userRepository.CreateAsync(user, ct);
        }
        catch (Exception)
        {
            ErrorMessage = "Something went wrong. Please try again.";
            return Page();
        }

        TempData["MSG_SUCCESS"] = "Sign up successfully! Please sign in.";
        return RedirectToPage("/SignIn");
    }

    public class SignUpInput
    {
        [Required(ErrorMessage = "Name is required")]
        [StringLength(100, MinimumLength = 2, ErrorMessage = "Name must be between 2 and 100 characters")]
        public string Name { get; set; } = string.Empty;

        [Required(ErrorMessage = "Email is required")]
        [EmailAddress(ErrorMessage = "Email is not valid")]
        [StringLength(150)]
        public string Email { get; set; } = string.Empty;

        [Required(ErrorMessage = "Password is required")]
        [StringLength(72, MinimumLength = 8, ErrorMessage = "Password must be between 8 and 72 characters")]
        public string Password { get; set; } = string.Empty;
    }
}