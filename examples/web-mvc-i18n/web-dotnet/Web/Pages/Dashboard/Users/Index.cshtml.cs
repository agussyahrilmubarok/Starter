using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.RazorPages;
using Web.Application.DTO.User;
using Web.Application.Service;
using Web.Common.Exceptions;
using Web.Common.Utils;
using Web.Resources.Lang;

namespace Web.Pages.Dashboard.Users;

public class IndexModel : PageModel
{
    private readonly IUserService _userService;
    private readonly MessageHelper _msg;

    public IndexModel(IUserService userService, MessageHelper msg)
    {
        _userService = userService;
        _msg = msg;
    }

    public IEnumerable<UserResponse> Users { get; private set; } = [];
    public string UserEmail { get; private set; } = string.Empty;

    public async Task<IActionResult> OnGetAsync(CancellationToken ct)
    {
        if (!SessionHelper.IsAuthenticated(HttpContext.Session))
            return RedirectToPage("/SignIn");

        UserEmail = SessionHelper.GetUserEmail(HttpContext.Session);
        Users = await _userService.GetAllAsync(ct);
        return Page();
    }

    public async Task<IActionResult> OnPostDeleteAsync(Guid id, CancellationToken ct)
    {
        if (!SessionHelper.IsAuthenticated(HttpContext.Session))
            return RedirectToPage("/SignIn");

        try
        {
            await _userService.DeleteAsync(id, ct);
            TempData[WebUtils.MsgSuccess] = _msg.Get("user.delete.success");
        }
        catch (NotFoundException)
        {
            TempData[WebUtils.MsgError] = _msg.Get("user.notFound");
        }

        return RedirectToPage();
    }
}