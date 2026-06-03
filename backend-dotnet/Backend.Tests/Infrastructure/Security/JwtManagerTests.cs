using Backend.Common.Exceptions;
using Backend.Infrastructure.Security;
using Microsoft.Extensions.Configuration;

namespace Backend.Tests.Unit;

public class JwtManagerTests
{
    private readonly IJwtManager _sut;

    public JwtManagerTests()
    {
        var config = new ConfigurationBuilder()
            .AddInMemoryCollection(new Dictionary<string, string?>
            {
                ["Jwt:Secret"] = "this-is-a-test-secret-key-minimum-32-chars!!",
                ["Jwt:Issuer"] = "Backend",
                ["Jwt:Audience"] = "BackendUsers",
                ["Jwt:ExpiresInMinutes"] = "1440"
            })
            .Build();

        _sut = new JwtManager(config);
    }

    [Fact]
    public void GenerateToken_ShouldReturnNonEmptyString()
    {
        var userId = Guid.NewGuid();

        var token = _sut.GenerateToken(userId);

        Assert.NotNull(token);
        Assert.NotEmpty(token);
    }

    [Fact]
    public void GenerateToken_ShouldReturnValidJwtFormat()
    {
        var userId = Guid.NewGuid();

        var token = _sut.GenerateToken(userId);

        var parts = token.Split('.');
        Assert.Equal(3, parts.Length);
    }

    [Fact]
    public void GenerateToken_TwoCalls_ShouldReturnDifferentTokens()
    {
        var userId = Guid.NewGuid();

        var token1 = _sut.GenerateToken(userId);
        var token2 = _sut.GenerateToken(userId);

        Assert.NotEqual(token1, token2);
    }

    [Fact]
    public void ValidateToken_WithValidToken_ShouldReturnUserId()
    {
        var userId = Guid.NewGuid();
        var token = _sut.GenerateToken(userId);

        var result = _sut.ValidateToken(token);

        Assert.Equal(userId, result);
    }

    [Fact]
    public void ValidateToken_WithInvalidToken_ShouldThrowJwtInvalidTokenException()
    {
        var invalidToken = "this.is.invalid";

        Assert.Throws<JwtInvalidTokenException>(
            () => _sut.ValidateToken(invalidToken));
    }

    [Fact]
    public void ValidateToken_WithTamperedToken_ShouldThrowJwtInvalidTokenException()
    {
        var userId = Guid.NewGuid();
        var token = _sut.GenerateToken(userId);
        var parts = token.Split('.');
        var tamperedToken = $"{parts[0]}.{parts[1]}.invalidsignature";

        Assert.Throws<JwtInvalidTokenException>(
            () => _sut.ValidateToken(tamperedToken));
    }

    [Fact]
    public void ValidateToken_WithEmptyToken_ShouldThrowJwtInvalidTokenException()
    {
        Assert.Throws<JwtInvalidTokenException>(
            () => _sut.ValidateToken(string.Empty));
    }

    [Fact]
    public void ValidateToken_WithExpiredToken_ShouldThrowJwtExpiredTokenException()
    {
        var config = new ConfigurationBuilder()
            .AddInMemoryCollection(new Dictionary<string, string?>
            {
                ["Jwt:Secret"] = "this-is-a-test-secret-key-minimum-32-chars!!",
                ["Jwt:Issuer"] = "Backend",
                ["Jwt:Audience"] = "BackendUsers",
                ["Jwt:ExpiresInMinutes"] = "0"
            })
            .Build();

        var expiredJwtManager = new JwtManager(config);
        var token = expiredJwtManager.GenerateToken(Guid.NewGuid());

        Assert.Throws<JwtExpiredTokenException>(
            () => _sut.ValidateToken(token));
    }

    [Fact]
    public void ValidateToken_WithWrongSecret_ShouldThrowJwtInvalidTokenException()
    {
        var config = new ConfigurationBuilder()
            .AddInMemoryCollection(new Dictionary<string, string?>
            {
                ["Jwt:Secret"] = "completely-different-secret-key-32-chars!!",
                ["Jwt:Issuer"] = "Backend",
                ["Jwt:Audience"] = "BackendUsers",
                ["Jwt:ExpiresInMinutes"] = "1440"
            })
            .Build();

        var otherJwtManager = new JwtManager(config);
        var token = otherJwtManager.GenerateToken(Guid.NewGuid());

        Assert.Throws<JwtInvalidTokenException>(
            () => _sut.ValidateToken(token));
    }

    [Fact]
    public void ValidateToken_WithWrongIssuer_ShouldThrowJwtInvalidTokenException()
    {
        var config = new ConfigurationBuilder()
            .AddInMemoryCollection(new Dictionary<string, string?>
            {
                ["Jwt:Secret"] = "this-is-a-test-secret-key-minimum-32-chars!!",
                ["Jwt:Issuer"] = "OtherIssuer",
                ["Jwt:Audience"] = "BackendUsers",
                ["Jwt:ExpiresInMinutes"] = "1440"
            })
            .Build();

        var otherJwtManager = new JwtManager(config);
        var token = otherJwtManager.GenerateToken(Guid.NewGuid());

        Assert.Throws<JwtInvalidTokenException>(
            () => _sut.ValidateToken(token));
    }
}