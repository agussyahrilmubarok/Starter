using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.RazorPages;
using Web.Application.DTO.Auth;
using Web.Application.Service;
using Web.Common.Exceptions;
using Web.Common.Utils;
using Web.Resources.Lang;

namespace Web.Pages;

public class SignInModel : PageModel
{
    private readonly IAuthService _authService;
    private readonly MessageHelper _msg;

    public SignInModel(IAuthService authService, MessageHelper msg)
    {
        _authService = authService;
        _msg = msg;
    }

    [BindProperty]
    public SignInRequest Input { get; set; } = new();

    public string? ErrorMessage { get; set; }

    public IActionResult OnGet()
    {
        if (SessionHelper.IsAuthenticated(HttpContext.Session))
            return RedirectToPage("/Dashboard/Index");
        return Page();
    }

    public async Task<IActionResult> OnPostAsync(CancellationToken ct)
    {
        if (!ModelState.IsValid) return Page();

        try
        {
            var user = await _authService.SignInAsync(Input, ct);
            SessionHelper.SetUser(HttpContext.Session, user);
        }
        catch (InvalidCredentialsException)
        {
            ErrorMessage = _msg.Get("auth.signIn.error");
            return Page();
        }
        catch (Exception)
        {
            ErrorMessage = _msg.Get("auth.general.error");
            return Page();
        }

        TempData[WebUtils.MsgSuccess] = _msg.Get("auth.signIn.success");
        return RedirectToPage("/Dashboard/Index");
    }
}