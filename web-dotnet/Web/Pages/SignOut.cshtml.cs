using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.RazorPages;
using Web.Common.Utils;

namespace Web.Pages;

[IgnoreAntiforgeryToken]
public class SignOutModel : PageModel
{
    public IActionResult OnGet()
        => RedirectToPage("/Dashboard/Index");

    public IActionResult OnPost()
    {
        HttpContext.Session.Clear();
        TempData[WebUtils.MsgSuccess] = "You have been signed out";
        return RedirectToPage("/SignIn");
    }
}