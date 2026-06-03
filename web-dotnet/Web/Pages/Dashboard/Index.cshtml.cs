using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.RazorPages;

namespace Web.Pages.Dashboard;

public class IndexModel : PageModel
{
    public string UserName { get; private set; } = string.Empty;
    public string UserEmail { get; private set; } = string.Empty;

    public IActionResult OnGet()
    {
        var userId = HttpContext.Session.GetString("UserId");
        if (string.IsNullOrEmpty(userId)) return RedirectToPage("/SignIn");

        UserName = HttpContext.Session.GetString("UserName") ?? "User";
        UserEmail = HttpContext.Session.GetString("UserEmail") ?? string.Empty;

        return Page();
    }
}