namespace Web.Common.Exceptions;

public class InvalidCredentialsException : System.Exception
{
    public InvalidCredentialsException() : base("Invalid email or password") { }
}