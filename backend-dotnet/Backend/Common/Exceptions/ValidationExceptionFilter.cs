using Backend.Application.DTOs;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.Filters;

namespace Backend.Api.Exceptions;

public class ValidationExceptionFilter : IActionFilter
{
    public void OnActionExecuting(ActionExecutingContext context)
    {
        if (context.ModelState.IsValid) return;

        var errors = context.ModelState
            .Where(e => e.Value?.Errors.Count > 0)
            .ToDictionary(
                kvp => Camelize(kvp.Key),
                kvp => BuildMessage(Camelize(kvp.Key), kvp.Value!.Errors.First())
            );

        context.Result = new ObjectResult(
            new ApiResponse<object>("Validation Failed", errors))
        {
            StatusCode = StatusCodes.Status422UnprocessableEntity
        };
    }

    public void OnActionExecuted(ActionExecutedContext context) { }

    private static string BuildMessage(string field, Microsoft.AspNetCore.Mvc.ModelBinding.ModelError error)
    {
        var name = Capitalize(field);
        var msg  = error.ErrorMessage ?? string.Empty;

        if (msg.Contains("required", StringComparison.OrdinalIgnoreCase) ||
            msg.Contains("must not be empty", StringComparison.OrdinalIgnoreCase))
            return $"{name} is required";

        if (msg.Contains("valid e-mail", StringComparison.OrdinalIgnoreCase) ||
            msg.Contains("email", StringComparison.OrdinalIgnoreCase))
            return "Invalid email format";

        var minLengthMatch = System.Text.RegularExpressions.Regex.Match(
            msg, @"minimum length of (\d+)");
        var maxLengthMatch = System.Text.RegularExpressions.Regex.Match(
            msg, @"maximum length of (\d+)");

        if (minLengthMatch.Success && maxLengthMatch.Success)
        {
            var min = minLengthMatch.Groups[1].Value;
            var max = maxLengthMatch.Groups[1].Value;
            return $"{name} must be between {min} and {max} characters";
        }

        if (minLengthMatch.Success)
        {
            var min = minLengthMatch.Groups[1].Value;
            return $"{name} must be at least {min} characters";
        }

        if (maxLengthMatch.Success)
        {
            var max = maxLengthMatch.Groups[1].Value;
            return $"{name} must be at most {max} characters";
        }

        var rangeMatch = System.Text.RegularExpressions.Regex.Match(
            msg, @"between (\d+) and (\d+)");

        if (rangeMatch.Success)
        {
            var min = rangeMatch.Groups[1].Value;
            var max = rangeMatch.Groups[2].Value;

            if (max == int.MaxValue.ToString())
                return $"{name} must be at least {min}";
            if (min == int.MinValue.ToString())
                return $"{name} must be at most {max}";

            return $"{name} must be between {min} and {max}";
        }

        if (msg.Contains("match", StringComparison.OrdinalIgnoreCase) ||
            msg.Contains("pattern", StringComparison.OrdinalIgnoreCase))
            return $"Invalid {field} format";

        return "Invalid value";
    }

    private static string Camelize(string value) =>
        string.IsNullOrEmpty(value) ? value : char.ToLower(value[0]) + value[1..];

    private static string Capitalize(string value) =>
        string.IsNullOrEmpty(value) ? value : char.ToUpper(value[0]) + value[1..];
}