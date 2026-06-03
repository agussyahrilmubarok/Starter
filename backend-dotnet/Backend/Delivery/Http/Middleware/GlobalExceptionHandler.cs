using System.Net;
using System.Text.Json;
using Backend.Common.Exceptions;
using Backend.Delivery.Http.Payload;

namespace Backend.Delivery.Http.Middleware;

public class GlobalExceptionHandler : IMiddleware
{
    private readonly ILogger<GlobalExceptionHandler> _logger;

    public GlobalExceptionHandler(ILogger<GlobalExceptionHandler> logger)
    {
        _logger = logger;
    }

    public async Task InvokeAsync(HttpContext context, RequestDelegate next)
    {
        try
        {
            await next(context);
        }
        catch (Exception ex)
        {
            _logger.LogError(ex, "Unhandled exception: {Message}", ex.Message);
            await HandleExceptionAsync(context, ex);
        }
    }

    private static async Task HandleExceptionAsync(HttpContext context, Exception exception)
    {
        context.Response.ContentType = "application/json";

        ApiResponse<object> response;

        switch (exception)
        {
            case NotFoundException:
                context.Response.StatusCode = (int)HttpStatusCode.NotFound;
                response = new ApiResponse<object>(
                    "Not found.",
                    new Dictionary<string, string>
                    {
                        ["error"] = exception.Message
                    });
                break;

            case EmailAlreadyExistsException:
                context.Response.StatusCode = (int)HttpStatusCode.Conflict;
                response = new ApiResponse<object>(
                    "Conflict.",
                    new Dictionary<string, string>
                    {
                        ["email"] = exception.Message
                    });
                break;

            case JwtExpiredTokenException:
                context.Response.StatusCode = (int)HttpStatusCode.Unauthorized;
                response = new ApiResponse<object>(
                    "Unauthorized.",
                    new Dictionary<string, string>
                    {
                        ["error"] = exception.Message
                    });
                break;

            case JwtInvalidTokenException:
                context.Response.StatusCode = (int)HttpStatusCode.Unauthorized;
                response = new ApiResponse<object>(
                    "Unauthorized.",
                    new Dictionary<string, string>
                    {
                        ["error"] = exception.Message
                    });
                break;

            case EmailNotRegisteredException:
                context.Response.StatusCode = (int)HttpStatusCode.Unauthorized;
                response = new ApiResponse<object>(
                    "Unauthorized.",
                    new Dictionary<string, string>
                    {
                        ["email"] = exception.Message
                    });
                break;

            case WrongPasswordException:
                context.Response.StatusCode = (int)HttpStatusCode.Unauthorized;
                response = new ApiResponse<object>(
                    "Unauthorized.",
                    new Dictionary<string, string>
                    {
                        ["password"] = exception.Message
                    });
                break;

            case UnauthorizedAccessException:
                context.Response.StatusCode = (int)HttpStatusCode.Unauthorized;
                response = new ApiResponse<object>("Unauthorized.",
                    new Dictionary<string, string> { ["error"] = exception.Message });
                break;

            case AppValidationException ve:
                context.Response.StatusCode = (int)HttpStatusCode.BadRequest;
                response = new ApiResponse<object>(
                    "Validation failed.",
                    ve.Errors);
                break;

            default:
                context.Response.StatusCode = (int)HttpStatusCode.InternalServerError;
                response = new ApiResponse<object>(
                    "Something went wrong.",
                    new Dictionary<string, string>
                    {
                        ["error"] = "Something went wrong."
                    });
                break;
        }

        var json = JsonSerializer.Serialize(response, new JsonSerializerOptions
        {
            PropertyNamingPolicy = JsonNamingPolicy.CamelCase
        });

        await context.Response.WriteAsync(json);
    }
}