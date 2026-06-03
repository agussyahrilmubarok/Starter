using Backend.Domain.User;
using Backend.Infrastructure.Persistence;
using Backend.Infrastructure.Persistence.Repository;
using Microsoft.EntityFrameworkCore;

namespace Backend.Tests.Infrastructure.Persistence.Repository;

public class UserRepositoryTests : IDisposable
{
    private readonly AppDbContext _context;
    private readonly UserRepository _sut;

    public UserRepositoryTests()
    {
        var options = new DbContextOptionsBuilder<AppDbContext>()
            .UseInMemoryDatabase(databaseName: Guid.NewGuid().ToString())
            .Options;

        _context = new AppDbContext(options);
        _sut = new UserRepository(_context);
    }

    public void Dispose()
    {
        _context.Database.EnsureDeleted();
        _context.Dispose();
    }

    private static User MakeUser(
        string name = "John Doe",
        string email = "john@mail.com",
        string password = "hashed_password")
        => User.Create(name, email, password);

    private async Task<User> SeedUser(
        string name = "John Doe",
        string email = "john@mail.com",
        string password = "hashed_password")
    {
        var user = MakeUser(name, email, password);
        await _context.Users.AddAsync(user);
        await _context.SaveChangesAsync();
        return user;
    }

    [Fact]
    public async Task FindAllAsync_WhenUsersExist_ShouldReturnAllUsers()
    {
        await SeedUser("Alice", "alice@mail.com", "hash1");
        await SeedUser("Bob", "bob@mail.com", "hash2");

        var result = await _sut.FindAllAsync();

        Assert.Equal(2, result.Count());
    }

    [Fact]
    public async Task FindAllAsync_WhenEmpty_ShouldReturnEmptyList()
    {
        var result = await _sut.FindAllAsync();

        Assert.Empty(result);
    }

    [Fact]
    public async Task FindByIdAsync_WhenUserExists_ShouldReturnUser()
    {
        var seeded = await SeedUser();

        var result = await _sut.FindByIdAsync(seeded.Id);

        Assert.NotNull(result);
        Assert.Equal(seeded.Id, result.Id);
        Assert.Equal(seeded.Email, result.Email);
    }

    [Fact]
    public async Task FindByIdAsync_WhenUserNotFound_ShouldReturnNull()
    {
        var result = await _sut.FindByIdAsync(Guid.NewGuid());

        Assert.Null(result);
    }

    [Fact]
    public async Task FindByEmailAsync_WhenEmailExists_ShouldReturnUser()
    {
        var seeded = await SeedUser();

        var result = await _sut.FindByEmailAsync(seeded.Email);

        Assert.NotNull(result);
        Assert.Equal(seeded.Email, result.Email);
    }

    [Fact]
    public async Task FindByEmailAsync_WhenEmailNotFound_ShouldReturnNull()
    {
        var result = await _sut.FindByEmailAsync("notexist@mail.com");

        Assert.Null(result);
    }

    [Fact]
    public async Task CreateAsync_ShouldPersistUser()
    {
        var user = MakeUser();

        var result = await _sut.CreateAsync(user);

        var inDb = await _context.Users.FindAsync(result.Id);
        Assert.NotNull(inDb);
        Assert.Equal(user.Email, inDb.Email);
        Assert.Equal(user.Name, inDb.Name);
    }

    [Fact]
    public async Task CreateAsync_ShouldReturnCreatedUser()
    {
        var user = MakeUser();

        var result = await _sut.CreateAsync(user);

        Assert.Equal(user.Id, result.Id);
        Assert.Equal(user.Name, result.Name);
        Assert.Equal(user.Email, result.Email);
    }

    [Fact]
    public async Task UpdateAsync_ShouldPersistChanges()
    {
        var seeded = await SeedUser();
        seeded.UpdateName("Updated Name");

        var result = await _sut.UpdateAsync(seeded);

        var inDb = await _context.Users
            .AsNoTracking()
            .FirstOrDefaultAsync(u => u.Id == seeded.Id);

        Assert.NotNull(inDb);
        Assert.Equal("Updated Name", inDb.Name);
    }

    [Fact]
    public async Task ExistsByEmailAsync_WhenEmailExists_ShouldReturnTrue()
    {
        var seeded = await SeedUser();

        var result = await _sut.ExistsByEmailAsync(seeded.Email);

        Assert.True(result);
    }

    [Fact]
    public async Task ExistsByEmailAsync_WhenEmailNotExists_ShouldReturnFalse()
    {
        var result = await _sut.ExistsByEmailAsync("ghost@mail.com");

        Assert.False(result);
    }

    [Fact]
    public async Task CountAsync_ShouldReturnCorrectCount()
    {
        await SeedUser("Alice", "alice@mail.com", "hash1");
        await SeedUser("Bob", "bob@mail.com", "hash2");
        await SeedUser("Charlie", "charlie@mail.com", "hash3");

        var count = await _sut.CountAsync();

        Assert.Equal(3, count);
    }

    [Fact]
    public async Task CountAsync_WhenEmpty_ShouldReturnZero()
    {
        var count = await _sut.CountAsync();

        Assert.Equal(0, count);
    }
}