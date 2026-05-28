namespace Backend.Common.Exceptions;

public class NotFoundException(string message, string? field = null) : Exception(message)
{
    public string? Field { get; } = field;
}

public class ConflictException(string message, string? field = null) : Exception(message)
{
    public string? Field { get; } = field;
}

public class UnauthorizedException(string message, string? field = null) : Exception(message)
{
    public string? Field { get; } = field;
}