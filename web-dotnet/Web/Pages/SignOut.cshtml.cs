using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.RazorPages;

namespace Web.Pages;

[IgnoreAntiforgeryToken]
public class SignOutModel : PageModel
{
    public IActionResult OnGet()
    {
        return RedirectToPage("/Dashboard/Index");
    }

    public IActionResult OnPost()
    {
        HttpContext.Session.Clear();
        TempData["MSG_SUCCESS"] = "You have signed out";
        return RedirectToPage("/SignIn");
    }
}