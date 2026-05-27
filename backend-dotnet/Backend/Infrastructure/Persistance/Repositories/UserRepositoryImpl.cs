using Backend.Domain;
using Microsoft.EntityFrameworkCore;

namespace Backend.Infrastructure.Persistence.Repositories;

public sealed class UserRepositoryImpl : IUserRepository
{
    private readonly AppDbContext _db;

    public UserRepositoryImpl(AppDbContext db)
    {
        _db = db;
    }

    public async Task<IEnumerable<User>> GetAllAsync(CancellationToken cancellationToken = default)
    {
        return await _db.Users
            .AsNoTracking()
            .OrderBy(u => u.CreatedAt)
            .ToListAsync(cancellationToken);
    }

    public async Task<User?> GetByIdAsync(Guid id, CancellationToken cancellationToken = default)
    {
        return await _db.Users
            .AsNoTracking()
            .FirstOrDefaultAsync(u => u.Id == id, cancellationToken);
    }

    public async Task<User?> GetByEmailAsync(string email, CancellationToken cancellationToken = default)
    {
        return await _db.Users
            .AsNoTracking()
            .FirstOrDefaultAsync(u => u.Email == email, cancellationToken);
    }

    public async Task<User> CreateAsync(User user, CancellationToken cancellationToken = default)
    {
        _db.Users.Add(user);
        await _db.SaveChangesAsync(cancellationToken);
        return user;
    }

    public async Task<User> UpdateAsync(User user, CancellationToken cancellationToken = default)
    {
        _db.Users.Update(user);
        await _db.SaveChangesAsync(cancellationToken);
        return user;
    }

    public async Task DeleteAsync(Guid id, CancellationToken cancellationToken = default)
    {
        var user = await _db.Users.FindAsync([id], cancellationToken)
            ?? throw new KeyNotFoundException($"User dengan id '{id}' tidak ditemukan.");

        _db.Users.Remove(user);
        await _db.SaveChangesAsync(cancellationToken);
    }

    public async Task<bool> ExistsByEmailAsync(string email, CancellationToken cancellationToken = default)
    {
        return await _db.Users
            .AsNoTracking()
            .AnyAsync(u => u.Email == email, cancellationToken);
    }
}