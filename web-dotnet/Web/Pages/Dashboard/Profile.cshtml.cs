using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.RazorPages;
using Web.Domain.User;

namespace Web.Pages.Dashboard;

public class ProfileModel : PageModel
{
    private readonly IUserRepository _userRepository;

    public ProfileModel(IUserRepository userRepository)
    {
        _userRepository = userRepository;
    }

    public User? User { get; private set; }
    public string UserEmail { get; private set; } = string.Empty;

    public async Task<IActionResult> OnGetAsync(CancellationToken ct)
    {
        var userId = HttpContext.Session.GetString("UserId");
        if (string.IsNullOrEmpty(userId)) return RedirectToPage("/SignIn");

        UserEmail = HttpContext.Session.GetString("UserEmail") ?? string.Empty;

        if (Guid.TryParse(userId, out var guid))
            User = await _userRepository.FindByIdAsync(guid, ct);

        return Page();
    }
}