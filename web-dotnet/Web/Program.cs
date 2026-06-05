using Web.Domain.User;
using Web.Infrastructure.Persistence;
using Web.Infrastructure.Persistence.Repository;
using Microsoft.EntityFrameworkCore;
using Web.Common.Logging;
using Serilog;
using Web.Common.Middleware;
using Web.Application.Service;
using System.Globalization;
using Microsoft.AspNetCore.Localization;

AppLogger.Initialize();

var builder = WebApplication.CreateBuilder(args);

builder.Host.UseSerilog();

builder.Services.AddLocalization(options => options.ResourcesPath = "");

builder.Services.AddRazorPages()
    .AddViewLocalization();

builder.Services.AddDistributedMemoryCache();
builder.Services.AddSession(options =>
{
    options.IdleTimeout = TimeSpan.FromHours(8);
    options.Cookie.HttpOnly = true;
    options.Cookie.IsEssential = true;
});

builder.Services.AddHttpContextAccessor();

builder.Services.AddDbContext<AppDbContext>(opt =>
    opt.UseNpgsql(builder.Configuration.GetConnectionString("DefaultConnection")));

builder.Services.AddScoped<IUserRepository, UserRepository>();
builder.Services.AddScoped<IAuthService, AuthService>();
builder.Services.AddScoped<IUserService, UserService>();
builder.Services.AddScoped<Web.Resources.Lang.MessageHelper>();

builder.Services.AddTransient<RequestIdMiddleware>();

var app = builder.Build();

var supportedCultures = new[]
{
    new CultureInfo("en")
};

app.UseRequestLocalization(new RequestLocalizationOptions
{
    DefaultRequestCulture = new RequestCulture("en"),
    SupportedCultures = supportedCultures,
    SupportedUICultures = supportedCultures,
    RequestCultureProviders = new List<IRequestCultureProvider>
    {
        new QueryStringRequestCultureProvider(),
        new CookieRequestCultureProvider(),
        new AcceptLanguageHeaderRequestCultureProvider()
    }
});

using (var scope = app.Services.CreateScope())
{
    var db = scope.ServiceProvider.GetRequiredService<AppDbContext>();
    try
    {
        await db.Database.CanConnectAsync();
        Log.Information("Database connected successfully");
    }
    catch (Exception ex)
    {
        Log.Fatal(ex, "Database connection failed");
    }
}

if (!app.Environment.IsDevelopment())
{
    app.UseExceptionHandler("/Error");
    app.UseHsts();
}

app.UseHttpsRedirection();
app.UseMiddleware<RequestIdMiddleware>();
app.UseSession();
app.UseRouting();
app.UseAuthorization();
app.MapStaticAssets();
app.MapRazorPages().WithStaticAssets();

app.Run();