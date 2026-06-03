using Serilog;
using Serilog.Events;

namespace Backend.Common.Logging;

public static class AppLogger
{
    public static void Initialize(string logPath = "logs/app-.log")
    {
        Log.Logger = new LoggerConfiguration()
            .MinimumLevel.Information()
            .MinimumLevel.Override("Microsoft.AspNetCore", LogEventLevel.Warning)
            .WriteTo.Console(outputTemplate:
                "[{Timestamp:yyyy-MM-dd HH:mm:ss} {Level:u3}] {SourceContext} {RequestId}: {Message:lj}{NewLine}{Exception}")
            .WriteTo.File(
                path: logPath,
                rollingInterval: RollingInterval.Day,
                retainedFileCountLimit: 10,
                fileSizeLimitBytes: 5 * 1024 * 1024,
                rollOnFileSizeLimit: true,
                outputTemplate:
                    "{Timestamp:yyyy-MM-dd HH:mm:ss} [{Level:u3}] {SourceContext} {RequestId}: {Message:lj}{Exception}{NewLine}")
            .Enrich.FromLogContext()
            .CreateLogger();
    }

    public static void Flush() => Log.CloseAndFlush();
}