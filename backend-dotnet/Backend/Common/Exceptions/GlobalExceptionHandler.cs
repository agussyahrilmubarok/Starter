using System.Text.Json;
using Backend.Application.DTOs;
using Microsoft.AspNetCore.Diagnostics;

namespace Backend.Common.Exceptions;

public class GlobalExceptionHandler(ILogger<GlobalExceptionHandler> logger) : IExceptionHandler
{
    public async ValueTask<bool> TryHandleAsync(
        HttpContext context,
        Exception exception,
        CancellationToken cancellationToken)
    {
        var (statusCode, title, errors) = exception switch
        {
            NotFoundException ex => (
                StatusCodes.Status404NotFound,
                "Not Found Error",
                BuildErrors(ex.Field, ex.Message)
            ),
            ConflictException ex => (
                StatusCodes.Status409Conflict,
                "Conflict Error",
                BuildErrors(ex.Field, ex.Message)
            ),
            UnauthorizedException ex => (
                StatusCodes.Status401Unauthorized,
                "Unauthorized Error",
                BuildErrors(ex.Field, ex.Message)
            ),
            JsonException => (
                StatusCodes.Status400BadRequest,
                "Bad Request Error",
                new Dictionary<string, string> { ["error"] = "Malformed JSON request body" }
            ),
            _ => (
                StatusCodes.Status500InternalServerError,
                "Internal Server Error",
                new Dictionary<string, string> { ["error"] = "Internal server error" }
            )
        };

        if (statusCode >= 500)
            logger.LogError(exception, "Unhandled exception: {Message}", exception.Message);
        else
            logger.LogWarning(exception, "Client error {StatusCode}: {Message}", statusCode, exception.Message);

        context.Response.StatusCode = statusCode;
        context.Response.ContentType = "application/json";

        await context.Response.WriteAsJsonAsync(
            new ApiResponse<object>(title, errors),
            cancellationToken
        );

        return true;
    }

    private static Dictionary<string, string> BuildErrors(string? field, string message)
    {
        return new Dictionary<string, string> { [field ?? "error"] = message };
    }
}