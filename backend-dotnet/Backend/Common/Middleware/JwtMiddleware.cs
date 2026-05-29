using System.Net;
using System.Text.Json;
using Backend.Application.Services;

namespace Backend.Common.Middleware;

public sealed class JwtMiddleware(RequestDelegate next)
{
    private static readonly HashSet<string> GuestPaths =
    [
        "/api/v1/auth/sign-up",
        "/api/v1/auth/sign-in",
        "/openapi/v1.json",
        "/swagger",
        "/swagger/index.html"
    ];

    public async Task InvokeAsync(HttpContext context, IJwtService jwtService)
    {
        var path = context.Request.Path.Value?.ToLowerInvariant() ?? "";
        if (GuestPaths.Contains(path) || path.StartsWith("/swagger") || path.StartsWith("/openapi"))
        {
            await next(context);
            return;
        }

        var authHeader = context.Request.Headers.Authorization.FirstOrDefault();
        if (string.IsNullOrWhiteSpace(authHeader) ||
            !authHeader.StartsWith("Bearer ", StringComparison.OrdinalIgnoreCase))
        {
            await WriteUnauthorizedResponse(context, "Authorization header missing or invalid.");
            return;
        }

        var token = authHeader["Bearer ".Length..].Trim();
        var userId = jwtService.ValidateToken(token);
        if (userId is null)
        {
            await WriteUnauthorizedResponse(context, "Token is invalid or expired.");
            return;
        }

        context.Items["UserId"] = userId;
        await next(context);
    }

    private static async Task WriteUnauthorizedResponse(HttpContext context, string message)
    {
        context.Response.StatusCode = (int)HttpStatusCode.Unauthorized;
        context.Response.ContentType = "application/json";
        var body = JsonSerializer.Serialize(new
        {
            message,
            errors = (object?)null
        });
        await context.Response.WriteAsync(body);
    }
}