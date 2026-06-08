using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.RazorPages;
using Web.Application.DTO.User;
using Web.Application.Service;
using Web.Common.Exceptions;
using Web.Common.Utils;

namespace Web.Pages.Dashboard;

public class ProfileModel : PageModel
{
    private readonly IUserService _userService;

    public ProfileModel(IUserService userService)
    {
        _userService = userService;
    }

    public new UserResponse? User { get; private set; }
    public string UserEmail { get; private set; } = string.Empty;

    public async Task<IActionResult> OnGetAsync(CancellationToken ct)
    {
        if (!SessionHelper.IsAuthenticated(HttpContext.Session))
            return RedirectToPage("/SignIn");

        UserEmail = SessionHelper.GetUserEmail(HttpContext.Session);

        var userId = SessionHelper.GetUserId(HttpContext.Session);
        if (userId is null)
            return RedirectToPage("/SignIn");

        try
        {
            User = await _userService.GetByIdAsync(userId.Value, ct);
        }
        catch (NotFoundException)
        {
            TempData[WebUtils.MsgError] = "User not found";
            return RedirectToPage("/SignIn");
        }

        return Page();
    }
}