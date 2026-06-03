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
    public string Message { get; set; } = "An unexpected error occurred. Please try again.";

    public void OnGet(int? statusCode = null)
    {
        RequestId = Activity.Current?.Id ?? HttpContext.TraceIdentifier;

        if (statusCode.HasValue)
        {
            StatusCode = statusCode.Value;
            Message = statusCode switch
            {
                404 => "The page you are looking for does not exist or has been moved.",
                403 => "You do not have permission to access this page.",
                401 => "You are not authorized to access this page.",
                500 => "An unexpected error occurred. Please try again.",
                _ => "An unexpected error occurred. Please try again."
            };
        }
    }
}