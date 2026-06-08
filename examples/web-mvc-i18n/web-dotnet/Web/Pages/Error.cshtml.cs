using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.RazorPages;
using System.Diagnostics;
using Web.Resources.Lang;

namespace Web.Pages;

[ResponseCache(Duration = 0, Location = ResponseCacheLocation.None, NoStore = true)]
[IgnoreAntiforgeryToken]
public class ErrorModel : PageModel
{
    private readonly MessageHelper _msg;

    public ErrorModel(MessageHelper msg)
    {
        _msg = msg;
    }

    public string RequestId { get; set; } = string.Empty;
    public bool ShowRequestId => !string.IsNullOrEmpty(RequestId);
    public new int StatusCode { get; set; } = 500;
    public string Message { get; set; } = string.Empty;

    public void OnGet(int? statusCode = null)
    {
        RequestId = Activity.Current?.Id ?? HttpContext.TraceIdentifier;

        StatusCode = statusCode ?? 500;
        Message = StatusCode switch
        {
            404 => _msg.Get("error.page.message"),
            403 => _msg.Get("error.unauthorized"),
            401 => _msg.Get("error.unauthorized"),
            _   => _msg.Get("error.general")
        };
    }
}