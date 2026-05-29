namespace Backend.Application.Services;

public interface IJwtService
{
    string GenerateToken(string userId);

    string? ValidateToken(string token);
}