namespace Web.Common.Exceptions;

public class NotFoundException : System.Exception
{
    public NotFoundException(string message) : base(message) { }
}