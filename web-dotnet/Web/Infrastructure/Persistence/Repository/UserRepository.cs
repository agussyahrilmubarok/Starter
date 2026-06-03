using Web.Domain.User;
using Microsoft.EntityFrameworkCore;

namespace Web.Infrastructure.Persistence.Repository;

public class UserRepository : IUserRepository
{
    private readonly AppDbContext _context;

    public UserRepository(AppDbContext context)
    {
        _context = context;
    }

    public async Task<IEnumerable<User>> FindAllAsync(CancellationToken ct = default)
    {
        return await _context.Users
            .AsNoTracking()
            .ToListAsync(ct);
    }

    public async Task<User?> FindByIdAsync(Guid id, CancellationToken ct = default)
    {
        return await _context.Users
            .AsNoTracking()
            .FirstOrDefaultAsync(u => u.Id == id, ct);
    }

    public async Task<User?> FindByEmailAsync(string email, CancellationToken ct = default)
    {
        return await _context.Users
            .AsNoTracking()
            .FirstOrDefaultAsync(u => u.Email == email, ct);
    }

    public async Task<User> CreateAsync(User user, CancellationToken ct = default)
    {
        await _context.Users.AddAsync(user, ct);
        await _context.SaveChangesAsync(ct);
        return user;
    }

    public async Task<User> UpdateAsync(User user, CancellationToken ct = default)
    {
        _context.Users.Update(user);
        await _context.SaveChangesAsync(ct);
        return user;
    }

    public async Task DeleteAsync(Guid id, CancellationToken ct = default)
    {
        await _context.Users
            .Where(u => u.Id == id)
            .ExecuteDeleteAsync(ct);
    }

    public async Task<bool> ExistsByEmailAsync(string email, CancellationToken ct = default)
    {
        return await _context.Users
            .AsNoTracking()
            .AnyAsync(u => u.Email == email, ct);
    }

    public async Task<int> CountAsync(CancellationToken ct = default)
    {
        return await _context.Users
            .AsNoTracking()
            .CountAsync(ct);
    }
}