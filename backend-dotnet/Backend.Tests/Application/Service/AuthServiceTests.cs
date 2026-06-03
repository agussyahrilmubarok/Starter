using Backend.Application.DTO.Auth;
using Backend.Application.Service;
using Backend.Common.Exceptions;
using Backend.Domain.User;
using Backend.Infrastructure.Security;
using Microsoft.Extensions.Logging;
using Moq;

namespace Backend.Tests.Application.Service;

public class AuthServiceTests
{
    private readonly Mock<IUserRepository> _repoMock;
    private readonly Mock<IJwtManager> _jwtMock;
    private readonly AuthService _sut;

    public AuthServiceTests()
    {
        _repoMock = new Mock<IUserRepository>();
        _jwtMock = new Mock<IJwtManager>();
        var loggerMock = new Mock<ILogger<AuthService>>();
        _sut = new AuthService(_repoMock.Object, _jwtMock.Object, loggerMock.Object);
    }

    private static User MakeUser(
        string name = "John Doe",
        string email = "john@mail.com",
        string password = "password123")
        => User.Create(name, email, BCrypt.Net.BCrypt.HashPassword(password));

    [Fact]
    public async Task SignUpAsync_WhenEmailNotTaken_ShouldReturnAuthResponse()
    {
        var request = new SignUpRequest("John Doe", "john@mail.com", "password123");
        var fakeToken = "fake.jwt.token";

        _repoMock
            .Setup(r => r.ExistsByEmailAsync(request.Email, It.IsAny<CancellationToken>()))
            .ReturnsAsync(false);

        _repoMock
            .Setup(r => r.CreateAsync(It.IsAny<User>(), It.IsAny<CancellationToken>()))
            .ReturnsAsync((User u, CancellationToken _) => u);

        _jwtMock
            .Setup(j => j.GenerateToken(It.IsAny<Guid>()))
            .Returns(fakeToken);

        var result = await _sut.SignUpAsync(request);

        Assert.NotNull(result);
        Assert.Equal(fakeToken, result.Token);
        Assert.Equal(request.Name, result.User.Name);
        Assert.Equal(request.Email, result.User.Email);
    }

    [Fact]
    public async Task SignUpAsync_ShouldCallCreateAsyncOnce()
    {
        var request = new SignUpRequest("John Doe", "john@mail.com", "password123");

        _repoMock
            .Setup(r => r.ExistsByEmailAsync(request.Email, It.IsAny<CancellationToken>()))
            .ReturnsAsync(false);

        _repoMock
            .Setup(r => r.CreateAsync(It.IsAny<User>(), It.IsAny<CancellationToken>()))
            .ReturnsAsync((User u, CancellationToken _) => u);

        _jwtMock
            .Setup(j => j.GenerateToken(It.IsAny<Guid>()))
            .Returns("fake.jwt.token");

        await _sut.SignUpAsync(request);

        _repoMock.Verify(
            r => r.CreateAsync(It.IsAny<User>(), It.IsAny<CancellationToken>()),
            Times.Once);
    }

    [Fact]
    public async Task SignUpAsync_WhenEmailAlreadyExists_ShouldThrowEmailAlreadyExistsException()
    {
        var request = new SignUpRequest("John Doe", "john@mail.com", "password123");

        _repoMock
            .Setup(r => r.ExistsByEmailAsync(request.Email, It.IsAny<CancellationToken>()))
            .ReturnsAsync(true);

        var ex = await Assert.ThrowsAsync<EmailAlreadyExistsException>(
            () => _sut.SignUpAsync(request));

        Assert.Contains(request.Email, ex.Message);

        _repoMock.Verify(
            r => r.CreateAsync(It.IsAny<User>(), It.IsAny<CancellationToken>()),
            Times.Never);
    }

    [Fact]
    public async Task SignUpAsync_ShouldHashPassword()
    {
        var request = new SignUpRequest("John Doe", "john@mail.com", "password123");
        User? capturedUser = null;

        _repoMock
            .Setup(r => r.ExistsByEmailAsync(request.Email, It.IsAny<CancellationToken>()))
            .ReturnsAsync(false);

        _repoMock
            .Setup(r => r.CreateAsync(It.IsAny<User>(), It.IsAny<CancellationToken>()))
            .Callback<User, CancellationToken>((u, _) => capturedUser = u)
            .ReturnsAsync((User u, CancellationToken _) => u);

        _jwtMock
            .Setup(j => j.GenerateToken(It.IsAny<Guid>()))
            .Returns("fake.jwt.token");

        await _sut.SignUpAsync(request);

        Assert.NotNull(capturedUser);
        Assert.NotEqual(request.Password, capturedUser.Password);
        Assert.True(BCrypt.Net.BCrypt.Verify(request.Password, capturedUser.Password));
    }

    [Fact]
    public async Task SignInAsync_WhenCredentialsValid_ShouldReturnAuthResponse()
    {
        var plainPassword = "password123";
        var user = MakeUser(password: plainPassword);
        var request = new SignInRequest(user.Email, plainPassword);
        var fakeToken = "fake.jwt.token";

        _repoMock
            .Setup(r => r.FindByEmailAsync(request.Email, It.IsAny<CancellationToken>()))
            .ReturnsAsync(user);

        _jwtMock
            .Setup(j => j.GenerateToken(user.Id))
            .Returns(fakeToken);

        var result = await _sut.SignInAsync(request);

        Assert.NotNull(result);
        Assert.Equal(fakeToken, result.Token);
        Assert.Equal(user.Email, result.User.Email);
    }

    [Fact]
    public async Task SignInAsync_WhenEmailNotRegistered_ShouldThrowEmailNotRegisteredException()
    {
        var request = new SignInRequest("ghost@mail.com", "password123");

        _repoMock
            .Setup(r => r.FindByEmailAsync(request.Email, It.IsAny<CancellationToken>()))
            .ReturnsAsync((User?)null);

        var ex = await Assert.ThrowsAsync<EmailNotRegisteredException>(
            () => _sut.SignInAsync(request));

        Assert.Contains(request.Email, ex.Message);

        _jwtMock.Verify(
            j => j.GenerateToken(It.IsAny<Guid>()),
            Times.Never);
    }

    [Fact]
    public async Task SignInAsync_WhenPasswordWrong_ShouldThrowWrongPasswordException()
    {
        var user = MakeUser(password: "correctpassword");
        var request = new SignInRequest(user.Email, "wrongpassword");

        _repoMock
            .Setup(r => r.FindByEmailAsync(request.Email, It.IsAny<CancellationToken>()))
            .ReturnsAsync(user);

        await Assert.ThrowsAsync<WrongPasswordException>(
            () => _sut.SignInAsync(request));

        _jwtMock.Verify(
            j => j.GenerateToken(It.IsAny<Guid>()),
            Times.Never);
    }

    [Fact]
    public async Task SignInAsync_ShouldCallGenerateTokenWithCorrectUserId()
    {
        var plainPassword = "password123";
        var user = MakeUser(password: plainPassword);
        var request = new SignInRequest(user.Email, plainPassword);

        _repoMock
            .Setup(r => r.FindByEmailAsync(request.Email, It.IsAny<CancellationToken>()))
            .ReturnsAsync(user);

        _jwtMock
            .Setup(j => j.GenerateToken(user.Id))
            .Returns("fake.jwt.token");

        await _sut.SignInAsync(request);

        _jwtMock.Verify(
            j => j.GenerateToken(user.Id),
            Times.Once);
    }
}