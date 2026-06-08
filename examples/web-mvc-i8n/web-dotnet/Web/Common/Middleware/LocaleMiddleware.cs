namespace Web.Common.Middleware;

public class LocaleMiddleware : IMiddleware
{
    private const string LangParam  = "lang";
    private const string LangCookie = "lang";

    public async Task InvokeAsync(HttpContext context, RequestDelegate next)
    {
        var lang = context.Request.Query[LangParam].FirstOrDefault();

        if (!string.IsNullOrWhiteSpace(lang))
        {
            context.Response.Cookies.Append(LangCookie, lang, new CookieOptions
            {
                Expires  = DateTimeOffset.UtcNow.AddYears(1),
                HttpOnly = false,
                IsEssential = true
            });
        }

        await next(context);
    }
}