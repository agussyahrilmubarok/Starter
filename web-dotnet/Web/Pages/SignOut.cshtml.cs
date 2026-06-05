using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.RazorPages;
using Web.Common.Utils;
using Web.Resources.Lang;

namespace Web.Pages;

[IgnoreAntiforgeryToken]
public class SignOutModel : PageModel
{
    private readonly MessageHelper _msg;

    public SignOutModel(MessageHelper msg)
    {
        _msg = msg;
    }

    public IActionResult OnGet()
        => RedirectToPage("/Dashboard/Index");

    public IActionResult OnPost()
    {
        HttpContext.Session.Clear();
        TempData[WebUtils.MsgSuccess] = _msg.Get("auth.signOut.success");
        return RedirectToPage("/SignIn");
    }
}