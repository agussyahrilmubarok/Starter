using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.RazorPages;
using System.ComponentModel.DataAnnotations;
using Web.Domain.User;

namespace Web.Pages;

public class SignInModel : PageModel
{
    private readonly IUserRepository _userRepository;

    public SignInModel(IUserRepository userRepository)
    {
        _userRepository = userRepository;
    }

    [BindProperty]
    public SignInInput Input { get; set; } = new();

    public string? ErrorMessage { get; set; }

    public void OnGet() { }

    public async Task<IActionResult> OnPostAsync(CancellationToken ct)
    {
        if (!ModelState.IsValid) return Page();

        var user = await _userRepository.FindByEmailAsync(Input.Email.ToLower(), ct);
        if (user == null || !BCrypt.Net.BCrypt.Verify(Input.Password, user.Password))
        {
            ErrorMessage = "Invalid email or password.";
            return Page();
        }

        HttpContext.Session.SetString("UserId", user.Id.ToString());
        HttpContext.Session.SetString("UserName", user.Name);
        HttpContext.Session.SetString("UserEmail", user.Email);

        return RedirectToPage("/Dashboard/Index");
    }

    public class SignInInput
    {
        [Required(ErrorMessage = "Email is required")]
        [EmailAddress(ErrorMessage = "Email is not valid")]
        [StringLength(150)]
        public string Email { get; set; } = string.Empty;

        [Required(ErrorMessage = "Password is required")]
        [StringLength(72, MinimumLength = 8, ErrorMessage = "Password must be between 8 and 72 characters")]
        public string Password { get; set; } = string.Empty;
    }
}