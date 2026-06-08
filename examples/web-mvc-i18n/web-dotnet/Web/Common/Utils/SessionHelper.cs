using Web.Application.DTO.User;

namespace Web.Common.Utils;

public static class SessionHelper
{
    public const string KeyUserId = "UserId";
    public const string KeyUserName = "UserName";
    public const string KeyUserEmail = "UserEmail";

    public static bool IsAuthenticated(ISession session)
        => !string.IsNullOrEmpty(session.GetString(KeyUserId));

    public static Guid? GetUserId(ISession session)
        => Guid.TryParse(session.GetString(KeyUserId), out var id) ? id : null;

    public static string GetUserName(ISession session)
        => session.GetString(KeyUserName) ?? "User";

    public static string GetUserEmail(ISession session)
        => session.GetString(KeyUserEmail) ?? string.Empty;

    public static void SetUser(ISession session, UserResponse user)
    {
        session.SetString(KeyUserId, user.Id.ToString());
        session.SetString(KeyUserName, user.Name);
        session.SetString(KeyUserEmail, user.Email);
    }

    public static void Clear(ISession session) => session.Clear();
}