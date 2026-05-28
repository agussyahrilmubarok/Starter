using Backend.Application.DTOs;
using Backend.Common.Exceptions;
using Backend.Common.Helpers;
using Backend.Domain;
using Backend.Infrastructure.Persistence.Repositories;

namespace Backend.Application.Services;

public sealed class AuthServiceImpl(IUserRepository userRepository) : IAuthService
{
    public async Task<AuthResponse> SignUpAsync(SignUpRequest request, CancellationToken ct = default)
    {
        var emailNormalized = request.Email.ToLowerInvariant();

        var exists = await userRepository.ExistsByEmailAsync(emailNormalized, ct);
        if (exists)
            throw new ConflictException("Email already registered.", "email");

        var user = new User
        {
            Name = request.Name,
            Email = emailNormalized,
            Password = PasswordHelper.Hash(request.Password)
        };

        var created = await userRepository.CreateAsync(user, ct);
        var token = TokenHelper.Generate(created.Id.ToString(), created.Email);

        return BuildResponse(token, created);
    }

    public async Task<AuthResponse> SignInAsync(SignInRequest request, CancellationToken ct = default)
    {
        var emailNormalized = request.Email.ToLowerInvariant();

        var user = await userRepository.GetByEmailAsync(emailNormalized, ct)
                   ?? throw new UnauthorizedException("Invalid email or password.", "credentials");

        var valid = PasswordHelper.Verify(request.Password, user.Password);
        if (!valid)
            throw new UnauthorizedException("Invalid email or password.", "credentials");

        var token = TokenHelper.Generate(user.Id.ToString(), user.Email);

        return BuildResponse(token, user);
    }

    private static AuthResponse BuildResponse(string token, User user)
    {
        return new AuthResponse(token, new AuthUserResponse(
            user.Id.ToString(),
            user.Name,
            user.Email,
            user.CreatedAt,
            user.UpdatedAt
        ));
    }
}