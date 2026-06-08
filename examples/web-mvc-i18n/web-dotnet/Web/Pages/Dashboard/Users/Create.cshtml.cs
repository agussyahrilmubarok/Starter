using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.RazorPages;
using Web.Application.DTO.User;
using Web.Application.Service;
using Web.Common.Exceptions;
using Web.Common.Utils;
using Web.Resources.Lang;

namespace Web.Pages.Dashboard.Users;

public class CreateModel : PageModel
{
    private readonly IUserService _userService;
    private readonly MessageHelper _msg;

    public CreateModel(IUserService userService, MessageHelper msg)
    {
        _userService = userService;
        _msg = msg;
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
            ModelState.AddModelError("Input.Email", _msg.Get("user.notFound"));
            return Page();
        }
        catch (Exception)
        {
            ErrorMessage = _msg.Get("error.general");
            return Page();
        }

        TempData[WebUtils.MsgSuccess] = _msg.Get("user.create.success");
        return RedirectToPage("/Dashboard/Users/Index");
    }
}