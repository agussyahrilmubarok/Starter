using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.RazorPages;
using Web.Application.DTO.Auth;
using Web.Application.Service;
using Web.Common.Exceptions;
using Web.Common.Utils;

namespace Web.Pages;

public class SignUpModel : PageModel
{
    private readonly IAuthService _authService;

    public SignUpModel(IAuthService authService)
    {
        _authService = authService;
    }

    [BindProperty]
    public SignUpRequest Input { get; set; } = new();

    public string? ErrorMessage { get; set; }

    public IActionResult OnGet()
    {
        if (!string.IsNullOrEmpty(HttpContext.Session.GetString("UserId")))
            return RedirectToPage("/Dashboard/Index");
        return Page();
    }

    public async Task<IActionResult> OnPostAsync(CancellationToken ct)
    {
        if (!ModelState.IsValid) return Page();

        try
        {
            await _authService.SignUpAsync(Input, ct);
        }
        catch (EmailAlreadyExistsException)
        {
            ModelState.AddModelError("Input.Email", "This email is already registered");
            return Page();
        }
        catch (Exception)
        {
            ErrorMessage = "Something went wrong. Please try again";
            return Page();
        }

        TempData[WebUtils.MsgSuccess] = "Sign up successfully! Please sign in";
        return RedirectToPage("/SignIn");
    }
}