namespace Backend.Common.Exceptions;

public class JwtInvalidTokenException : Exception
{
    public JwtInvalidTokenException()
        : base("Invalid token.") { }
}