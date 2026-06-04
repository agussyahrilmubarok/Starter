using System.IdentityModel.Tokens.Jwt;
using System.Security.Claims;
using System.Text;
using Backend.Common.Exceptions;
using Microsoft.IdentityModel.Tokens;

namespace Backend.Infrastructure.Security;

public class JwtManager : IJwtManager
{
    private readonly string _secret;
    private readonly int _expiresInMinutes;
    private readonly string _issuer;
    private readonly string _audience;

    public JwtManager(IConfiguration configuration)
    {
        _secret = configuration["Jwt:Secret"]
            ?? throw new InvalidOperationException("Jwt:Secret is not configured");
        _issuer = configuration["Jwt:Issuer"] ?? "Backend";
        _audience = configuration["Jwt:Audience"] ?? "BackendUsers";
        _expiresInMinutes = int.Parse(configuration["Jwt:ExpiresInMinutes"] ?? "1440");

        JwtSecurityTokenHandler.DefaultInboundClaimTypeMap.Clear();
    }

    public string GenerateToken(Guid userId)
    {
        var key = new SymmetricSecurityKey(Encoding.UTF8.GetBytes(_secret));
        var credentials = new SigningCredentials(key, SecurityAlgorithms.HmacSha256);

        var claims = new[]
        {
            new Claim(JwtRegisteredClaimNames.Sub, userId.ToString()),
            new Claim(JwtRegisteredClaimNames.Jti, Guid.NewGuid().ToString()),
            new Claim(JwtRegisteredClaimNames.Iat,
                DateTimeOffset.UtcNow.ToUnixTimeSeconds().ToString(),
                ClaimValueTypes.Integer64)
        };

        var token = new JwtSecurityToken(
            issuer: _issuer,
            audience: _audience,
            claims: claims,
            expires: DateTime.UtcNow.AddMinutes(_expiresInMinutes),
            signingCredentials: credentials
        );

        return new JwtSecurityTokenHandler().WriteToken(token);
    }

    public Guid ValidateToken(string token)
    {
        var key = new SymmetricSecurityKey(Encoding.UTF8.GetBytes(_secret));

        try
        {
            var principal = new JwtSecurityTokenHandler().ValidateToken(
                token,
                new TokenValidationParameters
                {
                    ValidateIssuerSigningKey = true,
                    IssuerSigningKey = key,
                    ValidateIssuer = true,
                    ValidIssuer = _issuer,
                    ValidateAudience = true,
                    ValidAudience = _audience,
                    ValidateLifetime = true,
                    ClockSkew = TimeSpan.Zero
                },
                out _);

            var userIdStr = principal.FindFirstValue(JwtRegisteredClaimNames.Sub)
                ?? throw new JwtInvalidTokenException();

            return Guid.Parse(userIdStr);
        }
        catch (SecurityTokenExpiredException)
        {
            throw new JwtExpiredTokenException();
        }
        catch (JwtExpiredTokenException)
        {
            throw;
        }
        catch
        {
            throw new JwtInvalidTokenException();
        }
    }
}