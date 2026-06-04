using Backend.Application.Service;
using Backend.Common.Extensions;
using Backend.Common.Logging;
using Backend.Delivery.Http.Middleware;
using Backend.Domain.User;
using Backend.Infrastructure.Persistence;
using Backend.Infrastructure.Persistence.Repository;
using Backend.Infrastructure.Security;
using Microsoft.EntityFrameworkCore;
using Serilog;

AppLogger.Initialize();

var builder = WebApplication.CreateBuilder(args);

builder.Host.UseSerilog();

builder.Services.AddDbContext<AppDbContext>(options =>
    options.UseNpgsql(builder.Configuration.GetConnectionString("DefaultConnection")));

builder.Services.AddScoped<IUserRepository, UserRepository>();
builder.Services.AddScoped<IUserService, UserService>();
builder.Services.AddScoped<IAuthService, AuthService>();
builder.Services.AddSingleton<IJwtManager, JwtManager>();

builder.Services.AddTransient<RequestIdMiddleware>();
builder.Services.AddTransient<GlobalExceptionHandler>();

builder.Services.AddJwtAuthentication(builder.Configuration);
builder.Services.AddAuthorization();
builder.Services.AddControllers();
builder.Services.AddValidationErrorResponse();
builder.Services.AddOpenApiWithBearer();
builder.Services.AddCorsPolicy(builder.Configuration);

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

if (app.Environment.IsDevelopment())
{
    app.MapOpenApi();
    app.UseSwaggerUI(options => options.SwaggerEndpoint("/openapi/v1.json", "Swagger"));
}

app.UseCors("DefaultPolicy");

if (!app.Environment.IsDevelopment())
{
    app.UseHttpsRedirection();
}

app.UseMiddleware<GlobalExceptionHandler>();
app.UseMiddleware<RequestIdMiddleware>();
app.UseAuthentication();
app.UseAuthorization();
app.MapControllers();

try
{
    app.Run();
}
finally
{
    AppLogger.Flush();
}