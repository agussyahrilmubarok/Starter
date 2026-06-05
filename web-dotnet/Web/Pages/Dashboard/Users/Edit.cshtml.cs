using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.RazorPages;
using Web.Application.DTO.User;
using Web.Application.Service;
using Web.Common.Exceptions;
using Web.Common.Utils;

namespace Web.Pages.Dashboard.Users;

public class EditModel : PageModel
{
    private readonly IUserService _userService;

    public EditModel(IUserService userService)
    {
        _userService = userService;
    }

    [BindProperty(SupportsGet = true)]
    public Guid Id { get; set; }

    [BindProperty]
    public UpdateUserRequest Input { get; set; } = new();

    public string? ErrorMessage { get; set; }
    public string UserEmail { get; private set; } = string.Empty;

    public async Task<IActionResult> OnGetAsync(CancellationToken ct)
    {
        if (!SessionHelper.IsAuthenticated(HttpContext.Session))
            return RedirectToPage("/SignIn");

        UserEmail = SessionHelper.GetUserEmail(HttpContext.Session);

        try
        {
            var user = await _userService.GetByIdAsync(Id, ct);
            Input = new UpdateUserRequest(user.Name, user.Email);
        }
        catch (NotFoundException)
        {
            TempData[WebUtils.MsgError] = "User not found";
            return RedirectToPage("/Dashboard/Users/Index");
        }

        return Page();
    }

    public async Task<IActionResult> OnPostAsync(CancellationToken ct)
    {
        UserEmail = SessionHelper.GetUserEmail(HttpContext.Session);

        if (!ModelState.IsValid) return Page();

        try
        {
            await _userService.UpdateAsync(Id, Input, ct);
        }
        catch (EmailAlreadyExistsException)
        {
            ModelState.AddModelError("Input.Email", "This email is already registered");
            return Page();
        }
        catch (NotFoundException)
        {
            TempData[WebUtils.MsgError] = "User not found";
            return RedirectToPage("/Dashboard/Users/Index");
        }
        catch (Exception)
        {
            ErrorMessage = "Something went wrong. Please try again";
            return Page();
        }

        TempData[WebUtils.MsgSuccess] = "User updated successfully";
        return RedirectToPage("/Dashboard/Users/Index");
    }
}