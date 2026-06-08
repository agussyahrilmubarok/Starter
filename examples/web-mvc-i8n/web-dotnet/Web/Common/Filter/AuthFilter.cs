using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.Filters;
using Web.Common.Utils;

namespace Web.Common.Filter;

public class AuthPageFilter : IPageFilter
{
    private static readonly string[] PublicPaths =
    [
        "/Index", "/SignIn", "/SignUp", "/SignOut", "/Error"
    ];

    public void OnPageHandlerSelected(PageHandlerSelectedContext context) { }

    public void OnPageHandlerExecuting(PageHandlerExecutingContext context)
    {
        var path = context.HttpContext.Request.Path.Value ?? string.Empty;

        bool isPublic = PublicPaths.Any(p =>
            path.Equals(p, StringComparison.OrdinalIgnoreCase) ||
            path.Equals("/", StringComparison.OrdinalIgnoreCase));

        if (!isPublic && !SessionHelper.IsAuthenticated(context.HttpContext.Session))
        {
            context.Result = new RedirectToPageResult("/SignIn");
        }
    }

    public void OnPageHandlerExecuted(PageHandlerExecutedContext context) { }
}