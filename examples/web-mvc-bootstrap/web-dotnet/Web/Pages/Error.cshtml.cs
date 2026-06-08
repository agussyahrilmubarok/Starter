using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.RazorPages;
using System.Diagnostics;

namespace Web.Pages;

[ResponseCache(Duration = 0, Location = ResponseCacheLocation.None, NoStore = true)]
[IgnoreAntiforgeryToken]
public class ErrorModel : PageModel
{
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
            404 => "The page you are looking for was not found.",
            403 => "You are not authorized to access this page.",
            401 => "You are not authorized to access this page.",
            _   => "An unexpected error occurred. Please try again later."
        };
    }
}