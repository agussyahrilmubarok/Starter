using Web.Domain.User;
using Web.Infrastructure.Persistence;
using Web.Infrastructure.Persistence.Repository;
using Microsoft.EntityFrameworkCore;
using Web.Common.Logging;
using Serilog;
using Web.Common.Middleware;
using Web.Application.Service;

AppLogger.Initialize();

var builder = WebApplication.CreateBuilder(args);

builder.Host.UseSerilog();

builder.Services.AddRazorPages();

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

builder.Services.AddTransient<RequestIdMiddleware>();

var app = builder.Build();

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
    app.UseStatusCodePagesWithReExecute("/Error", "?statusCode={0}");
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