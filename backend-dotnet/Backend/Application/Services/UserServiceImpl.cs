using Backend.Application.DTOs;
using Backend.Common.Exceptions;
using Backend.Common.Helpers;
using Backend.Domain;
using Backend.Infrastructure.Persistence.Repositories;

namespace Backend.Application.Services;

public sealed class UserServiceImpl(IUserRepository userRepository) : IUserService
{
    public async Task<IEnumerable<UserResponse>> GetAllAsync(CancellationToken ct = default)
    {
        var users = await userRepository.GetAllAsync(ct);
        return users.Select(ToResponse);
    }

    public async Task<UserResponse> GetByIdAsync(Guid id, CancellationToken ct = default)
    {
        var user = await userRepository.GetByIdAsync(id, ct)
                   ?? throw new NotFoundException($"User with id '{id}' was not found.", "id");

        return ToResponse(user);
    }

    public async Task<UserResponse> CreateAsync(CreateUserRequest request, CancellationToken ct = default)
    {
        var emailNormalized = request.Email.ToLowerInvariant();

        var exists = await userRepository.ExistsByEmailAsync(emailNormalized, ct);
        if (exists)
            throw new ConflictException("Email already exists.", "email");

        var user = new User
        {
            Name = request.Name,
            Email = emailNormalized,
            Password = PasswordHelper.Hash(request.Password)
        };

        var created = await userRepository.CreateAsync(user, ct);
        return ToResponse(created);
    }

    public async Task<UserResponse> UpdateAsync(Guid id, UpdateUserRequest request, CancellationToken ct = default)
    {
        var user = await userRepository.GetByIdAsync(id, ct)
                   ?? throw new NotFoundException($"User with id '{id}' was not found.", "id");

        if (request.Email is not null)
        {
            var emailNormalized = request.Email.ToLowerInvariant();
            if (emailNormalized != user.Email)
            {
                var exists = await userRepository.ExistsByEmailAsync(emailNormalized, ct);
                if (exists)
                    throw new ConflictException("Email already exists.", "email");

                user.Email = emailNormalized;
            }
        }

        if (request.Name is not null)
            user.Name = request.Name;

        if (request.Password is not null)
            user.Password = PasswordHelper.Hash(request.Password);

        var updated = await userRepository.UpdateAsync(user, ct);
        return ToResponse(updated);
    }

    public async Task DeleteAsync(Guid id, CancellationToken ct = default)
    {
        await userRepository.DeleteAsync(id, ct);
    }

    private static UserResponse ToResponse(User u)
    {
        return new UserResponse(u.Id.ToString(), u.Name, u.Email, u.CreatedAt, u.UpdatedAt);
    }
}