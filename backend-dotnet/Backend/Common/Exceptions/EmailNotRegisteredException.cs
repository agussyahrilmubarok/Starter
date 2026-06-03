namespace Backend.Common.Exceptions;

public class EmailNotRegisteredException : Exception
{
    public EmailNotRegisteredException(string email)
        : base($"Email {email} is not registered.") { }
}