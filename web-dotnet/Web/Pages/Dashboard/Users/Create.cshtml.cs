using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.RazorPages;
using System.ComponentModel.DataAnnotations;
using Web.Domain.User;
using DomainUser = Web.Domain.User.User;

namespace Web.Pages.Dashboard.Users;

public class CreateModel : PageModel
{
    private readonly IUserRepository _userRepository;

    public CreateModel(IUserRepository userRepository)
    {
        _userRepository = userRepository;
    }

    [BindProperty]
    public CreateInput Input { get; set; } = new();

    public string? ErrorMessage { get; set; }
    public string UserEmail { get; private set; } = string.Empty;

    public IActionResult OnGet()
    {
        if (string.IsNullOrEmpty(HttpContext.Session.GetString("UserId")))
            return RedirectToPage("/SignIn");

        UserEmail = HttpContext.Session.GetString("UserEmail") ?? string.Empty;
        return Page();
    }

    public async Task<IActionResult> OnPostAsync(CancellationToken ct)
    {
        UserEmail = HttpContext.Session.GetString("UserEmail") ?? string.Empty;

        if (!ModelState.IsValid) return Page();

        if (await _userRepository.ExistsByEmailAsync(Input.Email.ToLowerInvariant(), ct))
        {
            ModelState.AddModelError("Input.Email", "Email already exists.");
            return Page();
        }

        try
        {
            var hashedPassword = BCrypt.Net.BCrypt.HashPassword(Input.Password);
            var user = DomainUser.Create(Input.Name, Input.Email.ToLowerInvariant(), hashedPassword);
            await _userRepository.CreateAsync(user, ct);
        }
        catch (Exception)
        {
            ErrorMessage = "Something went wrong. Please try again";
            return Page();
        }

        TempData["MSG_SUCCESS"] = "User created successfully";
        return RedirectToPage("/Dashboard/Users/Index");
    }

    public class CreateInput
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