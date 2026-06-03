using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.RazorPages;
using Web.Domain.User;

namespace Web.Pages.Dashboard.Users;

public class IndexModel : PageModel
{
    private readonly IUserRepository _userRepository;

    public IndexModel(IUserRepository userRepository)
    {
        _userRepository = userRepository;
    }

    public IEnumerable<User> Users { get; private set; } = [];
    public string UserEmail { get; private set; } = string.Empty;

    public async Task<IActionResult> OnGetAsync(CancellationToken ct)
    {
        if (string.IsNullOrEmpty(HttpContext.Session.GetString("UserId")))
            return RedirectToPage("/SignIn");

        UserEmail = HttpContext.Session.GetString("UserEmail") ?? string.Empty;
        Users = await _userRepository.FindAllAsync(ct);
        return Page();
    }

    public async Task<IActionResult> OnPostDeleteAsync(Guid id, CancellationToken ct)
    {
        if (string.IsNullOrEmpty(HttpContext.Session.GetString("UserId")))
            return RedirectToPage("/SignIn");

        var user = await _userRepository.FindByIdAsync(id, ct);
        if (user == null)
        {
            TempData["MSG_ERROR"] = "User not found.";
            return RedirectToPage();
        }

        await _userRepository.DeleteAsync(id, ct);
        TempData["MSG_SUCCESS"] = "User was deleted successfully.";
        return RedirectToPage();
    }
}