using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.RazorPages;
using Web.Application.DTO.User;
using Web.Application.Service;
using Web.Common.Exceptions;
using Web.Common.Utils;
using Web.Resources.Lang;

namespace Web.Pages.Dashboard.Users;

public class EditModel : PageModel
{
    private readonly IUserService _userService;
    private readonly MessageHelper _msg;

    public EditModel(IUserService userService, MessageHelper msg)
    {
        _userService = userService;
        _msg = msg;
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
            TempData[WebUtils.MsgError] = _msg.Get("user.notFound");
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
            ModelState.AddModelError("Input.Email", _msg.Get("user.notFound"));
            return Page();
        }
        catch (NotFoundException)
        {
            TempData[WebUtils.MsgError] = _msg.Get("user.notFound");
            return RedirectToPage("/Dashboard/Users/Index");
        }
        catch (Exception)
        {
            ErrorMessage = _msg.Get("error.general");
            return Page();
        }

        TempData[WebUtils.MsgSuccess] = _msg.Get("user.update.success");
        return RedirectToPage("/Dashboard/Users/Index");
    }
}