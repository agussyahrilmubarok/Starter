namespace Web.Common.Exceptions;

public class EmailAlreadyExistsException : System.Exception
{
    public EmailAlreadyExistsException() : base("The email has already been taken") { }
}