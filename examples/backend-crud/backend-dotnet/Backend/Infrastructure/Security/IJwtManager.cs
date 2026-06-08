namespace Backend.Infrastructure.Security;

public interface IJwtManager
{
    string GenerateToken(Guid userId);
    Guid ValidateToken(string token);
}