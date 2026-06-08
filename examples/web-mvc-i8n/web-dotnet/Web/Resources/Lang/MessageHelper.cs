using Microsoft.Extensions.Localization;

namespace Web.Resources.Lang;

public class MessageHelper
{
    private readonly IStringLocalizer<MessageHelper> _localizer;
    private static MessageHelper? _instance;

    public MessageHelper(IStringLocalizer<MessageHelper> localizer)
    {
        _localizer = localizer;
        _instance = this;
    }

    public string Get(string key, params object[] args)
        => _localizer[key, args].Value;

    public static string GetMessage(string key, params object[] args)
        => _instance?.Get(key, args) ?? key;
}