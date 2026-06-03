namespace Backend.Common.Exceptions;

public class JwtExpiredTokenException : Exception
{
    public JwtExpiredTokenException()
        : base("Token has expired.") { }
}