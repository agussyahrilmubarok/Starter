using Backend.Common.Exceptions;
using Backend.Domain;
using Microsoft.EntityFrameworkCore;

namespace Backend.Infrastructure.Persistence.Repositories;

public sealed class UserRepositoryImpl(AppDbContext db) : IUserRepository
{
    public async Task<IEnumerable<User>> GetAllAsync(CancellationToken ct = default)
    {
        return await db.Users
            .AsNoTracking()
            .OrderByDescending(u => u.CreatedAt)
            .ToListAsync(ct);
    }

    public async Task<User?> GetByIdAsync(Guid id, CancellationToken ct = default)
    {
        return await db.Users
            .AsNoTracking()
            .FirstOrDefaultAsync(u => u.Id == id, ct);
    }

    public async Task<User?> GetByEmailAsync(string email, CancellationToken ct = default)
    {
        return await db.Users
            .AsNoTracking()
            .FirstOrDefaultAsync(u => u.Email == email.ToLowerInvariant(), ct);
    }

    public async Task<bool> ExistsByEmailAsync(string email, CancellationToken ct = default)
    {
        return await db.Users
            .AsNoTracking()
            .AnyAsync(u => u.Email == email.ToLowerInvariant(), ct);
    }

    public async Task<User> CreateAsync(User user, CancellationToken ct = default)
    {
        user.Email = user.Email.ToLowerInvariant();

        db.Users.Add(user);
        await db.SaveChangesAsync(ct);
        return user;
    }

    public async Task<User> UpdateAsync(User user, CancellationToken ct = default)
    {
        if (user.Email is not null)
            user.Email = user.Email.ToLowerInvariant();

        db.Users.Update(user);
        await db.SaveChangesAsync(ct);
        return user;
    }

    public async Task DeleteAsync(Guid id, CancellationToken ct = default)
    {
        var user = await db.Users.FindAsync([id], ct)
                   ?? throw new NotFoundException($"User with id '{id}' was not found.");

        db.Users.Remove(user);
        await db.SaveChangesAsync(ct);
    }
}