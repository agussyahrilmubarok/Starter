using Microsoft.AspNetCore.Mvc.Filters;
using Microsoft.AspNetCore.Mvc.RazorPages;
using Web.Common.Utils;

namespace Web.Common.Filter;

public class UserProfileFilter : IPageFilter
{
    public void OnPageHandlerSelected(PageHandlerSelectedContext context) { }

    public void OnPageHandlerExecuting(PageHandlerExecutingContext context)
    {
        if (context.HandlerInstance is PageModel page)
        {
            var session = context.HttpContext.Session;
            if (SessionHelper.IsAuthenticated(session))
            {
                page.ViewData["UserName"]  = SessionHelper.GetUserName(session);
                page.ViewData["UserEmail"] = SessionHelper.GetUserEmail(session);
                page.ViewData["UserId"]    = SessionHelper.GetUserId(session)?.ToString();
            }
        }
    }

    public void OnPageHandlerExecuted(PageHandlerExecutedContext context) { }
}