using Backend.Application.DTO.User;
using Backend.Application.Service;
using Backend.Common.Exceptions;
using Backend.Domain.User;
using Microsoft.Extensions.Logging;
using Moq;

namespace Backend.Tests.Application.Service;

public class UserServiceTests
{
    private readonly Mock<IUserRepository> _repoMock;
    private readonly UserService _sut;

    public UserServiceTests()
    {
        _repoMock = new Mock<IUserRepository>();
        var loggerMock = new Mock<ILogger<UserService>>();
        _sut = new UserService(_repoMock.Object, loggerMock.Object);
    }

    private static User MakeUser(
        string name = "John Doe",
        string email = "john@mail.com",
        string password = "hashed_password")
        => User.Create(name, email, password);

    [Fact]
    public async Task GetAllAsync_ShouldReturnAllUsers()
    {
        var users = new List<User>
        {
            MakeUser("Alice", "alice@mail.com", "hash1"),
            MakeUser("Bob",   "bob@mail.com",   "hash2"),
        };

        _repoMock
            .Setup(r => r.FindAllAsync(It.IsAny<CancellationToken>()))
            .ReturnsAsync(users);

        var result = await _sut.GetAllAsync();

        Assert.Equal(2, result.Count());
        Assert.Contains(result, u => u.Email == "alice@mail.com");
        Assert.Contains(result, u => u.Email == "bob@mail.com");
    }

    [Fact]
    public async Task GetAllAsync_WhenEmpty_ShouldReturnEmptyList()
    {
        _repoMock
            .Setup(r => r.FindAllAsync(It.IsAny<CancellationToken>()))
            .ReturnsAsync(new List<User>());

        var result = await _sut.GetAllAsync();

        Assert.Empty(result);
    }

    [Fact]
    public async Task GetByIdAsync_WhenUserExists_ShouldReturnUser()
    {
        var user = MakeUser();

        _repoMock
            .Setup(r => r.FindByIdAsync(user.Id, It.IsAny<CancellationToken>()))
            .ReturnsAsync(user);

        var result = await _sut.GetByIdAsync(user.Id);

        Assert.Equal(user.Id, result.Id);
        Assert.Equal(user.Name, result.Name);
        Assert.Equal(user.Email, result.Email);
    }

    [Fact]
    public async Task GetByIdAsync_WhenUserNotFound_ShouldThrowNotFoundException()
    {
        var randomId = Guid.NewGuid();

        _repoMock
            .Setup(r => r.FindByIdAsync(randomId, It.IsAny<CancellationToken>()))
            .ReturnsAsync((User?)null);

        var ex = await Assert.ThrowsAsync<NotFoundException>(
            () => _sut.GetByIdAsync(randomId));

        Assert.Contains(randomId.ToString(), ex.Message);
    }

    [Fact]
    public async Task CreateAsync_WhenEmailNotTaken_ShouldReturnCreatedUser()
    {
        var request = new CreateUserRequest("John Doe", "john@mail.com", "password123");

        _repoMock
            .Setup(r => r.ExistsByEmailAsync(request.Email, It.IsAny<CancellationToken>()))
            .ReturnsAsync(false);

        _repoMock
            .Setup(r => r.CreateAsync(It.IsAny<User>(), It.IsAny<CancellationToken>()))
            .ReturnsAsync((User u, CancellationToken _) => u);

        var result = await _sut.CreateAsync(request);

        Assert.Equal(request.Name, result.Name);
        Assert.Equal(request.Email, result.Email);
        Assert.NotEqual(Guid.Empty, result.Id);

        _repoMock.Verify(
            r => r.CreateAsync(It.IsAny<User>(), It.IsAny<CancellationToken>()),
            Times.Once);
    }

    [Fact]
    public async Task CreateAsync_WhenEmailAlreadyExists_ShouldThrowEmailAlreadyExistsException()
    {
        var request = new CreateUserRequest("John Doe", "john@mail.com", "password123");

        _repoMock
            .Setup(r => r.ExistsByEmailAsync(request.Email, It.IsAny<CancellationToken>()))
            .ReturnsAsync(true);

        var ex = await Assert.ThrowsAsync<EmailAlreadyExistsException>(
            () => _sut.CreateAsync(request));

        Assert.Contains(request.Email, ex.Message);

        _repoMock.Verify(
            r => r.CreateAsync(It.IsAny<User>(), It.IsAny<CancellationToken>()),
            Times.Never);
    }

    [Fact]
    public async Task UpdateAsync_WhenUserExists_ShouldUpdateAndReturnUser()
    {
        var user = MakeUser("Old Name", "old@mail.com", "oldhash");
        var request = new UpdateUserRequest("New Name", null, null);

        _repoMock
            .Setup(r => r.FindByIdAsync(user.Id, It.IsAny<CancellationToken>()))
            .ReturnsAsync(user);

        _repoMock
            .Setup(r => r.UpdateAsync(It.IsAny<User>(), It.IsAny<CancellationToken>()))
            .ReturnsAsync((User u, CancellationToken _) => u);

        var result = await _sut.UpdateAsync(user.Id, request);

        Assert.Equal("New Name", result.Name);
        Assert.Equal("old@mail.com", result.Email);

        _repoMock.Verify(
            r => r.UpdateAsync(It.IsAny<User>(), It.IsAny<CancellationToken>()),
            Times.Once);
    }

    [Fact]
    public async Task UpdateAsync_WhenUserNotFound_ShouldThrowNotFoundException()
    {
        var randomId = Guid.NewGuid();

        _repoMock
            .Setup(r => r.FindByIdAsync(randomId, It.IsAny<CancellationToken>()))
            .ReturnsAsync((User?)null);

        await Assert.ThrowsAsync<NotFoundException>(
            () => _sut.UpdateAsync(randomId, new UpdateUserRequest(null, null, null)));
    }

    [Fact]
    public async Task DeleteAsync_WhenUserExists_ShouldCallDeleteOnce()
    {
        var user = MakeUser();

        _repoMock
            .Setup(r => r.FindByIdAsync(user.Id, It.IsAny<CancellationToken>()))
            .ReturnsAsync(user);

        _repoMock
            .Setup(r => r.DeleteAsync(user.Id, It.IsAny<CancellationToken>()))
            .Returns(Task.CompletedTask);

        await _sut.DeleteAsync(user.Id);

        _repoMock.Verify(
            r => r.DeleteAsync(user.Id, It.IsAny<CancellationToken>()),
            Times.Once);
    }

    [Fact]
    public async Task DeleteAsync_WhenUserNotFound_ShouldThrowNotFoundException()
    {
        var randomId = Guid.NewGuid();

        _repoMock
            .Setup(r => r.FindByIdAsync(randomId, It.IsAny<CancellationToken>()))
            .ReturnsAsync((User?)null);

        await Assert.ThrowsAsync<NotFoundException>(
            () => _sut.DeleteAsync(randomId));

        _repoMock.Verify(
            r => r.DeleteAsync(It.IsAny<Guid>(), It.IsAny<CancellationToken>()),
            Times.Never);
    }
}