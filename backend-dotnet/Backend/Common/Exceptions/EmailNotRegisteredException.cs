namespace Backend.Common.Exceptions;

public class EmailNotRegisteredException : Exception
{
    public EmailNotRegisteredException()
        : base("The email address is not registered") { }
}