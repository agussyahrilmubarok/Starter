using Backend.Application.DTOs;
using Backend.Domain;
using Backend.Infrastructure.Persistence.Repositories;

namespace Backend.Application.Services;

public sealed class UserServiceImpl : IUserService
{
    private readonly IUserRepository _repo;

    public UserServiceImpl(IUserRepository repo)
    {
        _repo = repo;
    }

    public async Task<IEnumerable<UserResponse>> GetAllAsync(CancellationToken cancellationToken = default)
    {
        var users = await _repo.GetAllAsync(cancellationToken);
        return users.Select(ToResponse);
    }

    public async Task<UserResponse> GetByIdAsync(Guid id, CancellationToken cancellationToken = default)
    {
        var user = await _repo.GetByIdAsync(id, cancellationToken)
            ?? throw new KeyNotFoundException($"User with id '{id}' was not found.");

        return ToResponse(user);
    }

    public async Task<UserResponse> CreateAsync(CreateUserRequest request, CancellationToken cancellationToken = default)
    {
        var emailTaken = await _repo.ExistsByEmailAsync(request.Email, cancellationToken);
        if (emailTaken)
            throw new InvalidOperationException($"Email '{request.Email}' is already in use.");

        var user = new User
        {
            Name = request.Name,
            Email = request.Email,
            Password = BCrypt.Net.BCrypt.HashPassword(request.Password),
        };

        var created = await _repo.CreateAsync(user, cancellationToken);
        return ToResponse(created);
    }

    public async Task<UserResponse> UpdateAsync(Guid id, UpdateUserRequest request, CancellationToken cancellationToken = default)
    {
        var user = await _repo.GetByIdAsync(id, cancellationToken)
            ?? throw new KeyNotFoundException($"User with id '{id}' was not found.");

        var emailTakenByOther = await _repo.GetByEmailAsync(request.Email, cancellationToken);
        if (emailTakenByOther is not null && emailTakenByOther.Id != id)
            throw new InvalidOperationException($"Email '{request.Email}' is already used by another user.");

        user.Name = request.Name;
        user.Email = request.Email;

        var updated = await _repo.UpdateAsync(user, cancellationToken);
        return ToResponse(updated);
    }

    public async Task DeleteAsync(Guid id, CancellationToken cancellationToken = default)
    {
        await _repo.DeleteAsync(id, cancellationToken);
    }

    private static UserResponse ToResponse(User u) =>
        new(u.Id, u.Name, u.Email, u.CreatedAt, u.UpdatedAt);
}