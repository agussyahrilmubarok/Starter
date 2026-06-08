using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.RazorPages;
using Web.Application.DTO.Auth;
using Web.Application.Service;
using Web.Common.Exceptions;
using Web.Common.Utils;

namespace Web.Pages;

public class SignInModel : PageModel
{
    private readonly IAuthService _authService;

    public SignInModel(IAuthService authService)
    {
        _authService = authService;
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
            ErrorMessage = "Invalid email or password";
            return Page();
        }
        catch (Exception)
        {
            ErrorMessage = "Something went wrong. Please try again";
            return Page();
        }

        TempData[WebUtils.MsgSuccess] = "Signed in successfully";
        return RedirectToPage("/Dashboard/Index");
    }
}