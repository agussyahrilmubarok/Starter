using Serilog.Context;

namespace Backend.Delivery.Http.Middleware;

public class RequestIdMiddleware : IMiddleware
{
    private const string HeaderName = "X-Request-ID";

    public async Task InvokeAsync(HttpContext context, RequestDelegate next)
    {
        var requestId = context.Request.Headers[HeaderName].FirstOrDefault()
                        ?? Guid.NewGuid().ToString();

        context.Response.Headers[HeaderName] = requestId;

        using (LogContext.PushProperty("RequestId", requestId))
        {
            await next(context);
        }
    }
}