namespace Backend.Common.Exceptions;

public class WrongPasswordException : Exception
{
    public WrongPasswordException()
        : base("The password is incorrect") { }
}