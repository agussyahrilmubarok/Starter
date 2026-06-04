using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.RazorPages;
using System.ComponentModel.DataAnnotations;
using Web.Domain.User;

namespace Web.Pages.Dashboard.Users;

public class EditModel : PageModel
{
    private readonly IUserRepository _userRepository;

    public EditModel(IUserRepository userRepository)
    {
        _userRepository = userRepository;
    }

    [BindProperty(SupportsGet = true)]
    public Guid Id { get; set; }

    [BindProperty]
    public EditInput Input { get; set; } = new();

    public string? ErrorMessage { get; set; }
    public string UserEmail { get; private set; } = string.Empty;

    public async Task<IActionResult> OnGetAsync(CancellationToken ct)
    {
        if (string.IsNullOrEmpty(HttpContext.Session.GetString("UserId")))
            return RedirectToPage("/SignIn");

        UserEmail = HttpContext.Session.GetString("UserEmail") ?? string.Empty;

        var user = await _userRepository.FindByIdAsync(Id, ct);
        if (user == null)
        {
            TempData["MSG_ERROR"] = "User not found";
            return RedirectToPage("/Dashboard/Users/Index");
        }

        Input.Name = user.Name;
        Input.Email = user.Email;
        return Page();
    }

    public async Task<IActionResult> OnPostAsync(CancellationToken ct)
    {
        UserEmail = HttpContext.Session.GetString("UserEmail") ?? string.Empty;

        if (!ModelState.IsValid) return Page();

        var user = await _userRepository.FindByIdAsync(Id, ct);
        if (user == null)
        {
            TempData["MSG_ERROR"] = "User not found";
            return RedirectToPage("/Dashboard/Users/Index");
        }

        try
        {
            if (!string.IsNullOrWhiteSpace(Input.Name))
                user.UpdateName(Input.Name);

            if (!string.IsNullOrWhiteSpace(Input.Email) &&
                !user.Email.Equals(Input.Email, StringComparison.OrdinalIgnoreCase))
            {
                if (await _userRepository.ExistsByEmailAsync(Input.Email.ToLowerInvariant(), ct))
                {
                    ModelState.AddModelError("Input.Email", "The email has already been taken");
                    return Page();
                }
                user.UpdateEmail(Input.Email.ToLowerInvariant());
            }

            if (!string.IsNullOrWhiteSpace(Input.Password))
                user.UpdatePassword(BCrypt.Net.BCrypt.HashPassword(Input.Password));

            await _userRepository.UpdateAsync(user, ct);
        }
        catch (Exception)
        {
            ErrorMessage = "Something went wrong. Please try again";
            return Page();
        }

        TempData["MSG_SUCCESS"] = "User updated successfully";
        return RedirectToPage("/Dashboard/Users/Index");
    }

    public class EditInput
    {
        [StringLength(100, MinimumLength = 2, ErrorMessage = "Name must be between 2 and 100 characters")]
        public string? Name { get; set; }

        [EmailAddress(ErrorMessage = "Email is not valid")]
        [StringLength(150)]
        public string? Email { get; set; }

        [StringLength(72, MinimumLength = 8, ErrorMessage = "Password must be between 8 and 72 characters")]
        public string? Password { get; set; }
    }
}