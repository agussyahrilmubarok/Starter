namespace Web.Common.Utils;

public static class WebUtils
{
    public const string MsgSuccess = "MSG_SUCCESS";
    public const string MsgInfo    = "MSG_INFO";
    public const string MsgError   = "MSG_ERROR";

    public static string GetMessage(string key, params object[] args)
        => Web.Resources.Lang.MessageHelper.GetMessage(key, args);
}