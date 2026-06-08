using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.RazorPages;
using Web.Common.Utils;

namespace Web.Pages.Dashboard;

public class IndexModel : PageModel
{
    public string UserName  { get; private set; } = string.Empty;
    public string UserEmail { get; private set; } = string.Empty;

    public IActionResult OnGet()
    {
        if (!SessionHelper.IsAuthenticated(HttpContext.Session))
            return RedirectToPage("/SignIn");

        UserName  = SessionHelper.GetUserName(HttpContext.Session);
        UserEmail = SessionHelper.GetUserEmail(HttpContext.Session);
        return Page();
    }
}