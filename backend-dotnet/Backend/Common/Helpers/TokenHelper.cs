namespace Backend.Common.Helpers;

public static class TokenHelper
{
    public static string Generate(string userId, string email)
    {
        var raw = $"{userId}:{email}:{DateTimeOffset.UtcNow.ToUnixTimeSeconds()}";
        var bytes = System.Text.Encoding.UTF8.GetBytes(raw);
        return Convert.ToBase64String(bytes);
    }
}