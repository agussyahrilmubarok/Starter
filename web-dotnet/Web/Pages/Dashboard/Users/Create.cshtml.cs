using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.RazorPages;
using Web.Application.DTO.User;
using Web.Application.Service;
using Web.Common.Exceptions;
using Web.Common.Utils;

namespace Web.Pages.Dashboard.Users;

public class CreateModel : PageModel
{
    private readonly IUserService _userService;

    public CreateModel(IUserService userService)
    {
        _userService = userService;
    }

    [BindProperty]
    public CreateUserRequest Input { get; set; } = new();

    public string? ErrorMessage { get; set; }
    public string UserEmail { get; private set; } = string.Empty;

    public IActionResult OnGet()
    {
        if (!SessionHelper.IsAuthenticated(HttpContext.Session))
            return RedirectToPage("/SignIn");

        UserEmail = SessionHelper.GetUserEmail(HttpContext.Session);
        return Page();
    }

    public async Task<IActionResult> OnPostAsync(CancellationToken ct)
    {
        UserEmail = SessionHelper.GetUserEmail(HttpContext.Session);

        if (!ModelState.IsValid) return Page();

        try
        {
            await _userService.CreateAsync(Input, ct);
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

        TempData[WebUtils.MsgSuccess] = "User created successfully";
        return RedirectToPage("/Dashboard/Users/Index");
    }
}