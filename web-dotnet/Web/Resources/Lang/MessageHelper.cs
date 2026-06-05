using Microsoft.Extensions.Localization;

namespace Web.Resources.Lang;

public class MessageHelper
{
    private readonly IStringLocalizer<MessageHelper> _localizer;

    public MessageHelper(IStringLocalizer<MessageHelper> localizer)
    {
        _localizer = localizer;
    }

    public string Get(string key, params object[] args)
        => _localizer[key, args].Value;
}