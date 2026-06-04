namespace Backend.Common.Exceptions;

public class EmailAlreadyExistsException : Exception
{
    public EmailAlreadyExistsException()
        : base("The email has already been taken") { }
}