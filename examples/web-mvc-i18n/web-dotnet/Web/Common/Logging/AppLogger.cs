using Serilog;
using Serilog.Events;

namespace Web.Common.Logging;

public static class AppLogger
{
    public static void Initialize(string logPath = "logs/app-.log")
    {
        Log.Logger = new LoggerConfiguration()
            .MinimumLevel.Information()
            .MinimumLevel.Override("Microsoft.AspNetCore", LogEventLevel.Warning)
            .WriteTo.Console(outputTemplate:
                "{Timestamp:yyyy-MM-dd HH:mm:ss} [{ThreadId}] {Level:u5} [{RequestId}] {SourceContext} - {Message:lj}{NewLine}{Exception}")
            .WriteTo.File(
                path: logPath,
                rollingInterval: RollingInterval.Day,
                retainedFileCountLimit: 7,
                fileSizeLimitBytes: 10 * 1024 * 1024,
                rollOnFileSizeLimit: true,
                outputTemplate:
                    "{Timestamp:yyyy-MM-dd HH:mm:ss} [{ThreadId}] {Level:u5} [{RequestId}] {SourceContext} - {Message:lj}{Exception}{NewLine}")
            .Enrich.FromLogContext()
            .Enrich.WithThreadId()
            .CreateLogger();
    }

    public static void Flush() => Log.CloseAndFlush();
}