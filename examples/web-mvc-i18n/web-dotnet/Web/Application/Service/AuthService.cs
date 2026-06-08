using Web.Application.DTO.Auth;
using Web.Application.DTO.User;
using Web.Common.Exceptions;
using Web.Domain.User;

namespace Web.Application.Service;

public class AuthService : IAuthService
{

    private readonly IUserRepository _userRepository;
    private readonly ILogger<AuthService> _logger;

    public AuthService(IUserRepository userRepository, ILogger<AuthService> logger)
    {
        _userRepository = userRepository;
        _logger = logger;
    }

    public async Task<UserResponse> SignUpAsync(SignUpRequest request, CancellationToken ct = default)
    {
        if (await _userRepository.ExistsByEmailAsync(request.Email.ToLowerInvariant(), ct))
        {
            _logger.LogWarning("Email already taken: {Email}", request.Email);
            throw new EmailAlreadyExistsException();
        }

        var hashedPassword = BCrypt.Net.BCrypt.HashPassword(request.Password);
        var user = User.Create(request.Name, request.Email.ToLowerInvariant(), hashedPassword);
        await _userRepository.CreateAsync(user, ct);

        _logger.LogInformation("Signed up successfully {Id}", user.Id);
        return ToResponse(user);
    }

    public async Task<UserResponse> SignInAsync(SignInRequest request, CancellationToken ct = default)
    {
        var user = await _userRepository.FindByEmailAsync(request.Email.ToLowerInvariant(), ct);

        if (user is null || !BCrypt.Net.BCrypt.Verify(request.Password, user.Password))
        {
            _logger.LogWarning("Sign in failed for email: {Email}", request.Email);
            throw new InvalidCredentialsException();
        }

        _logger.LogInformation("Signed in successfully {Id}", user.Id);
        return ToResponse(user);
    }

    private static UserResponse ToResponse(User user) =>
        new(user.Id, user.Name, user.Email, user.CreatedAt, user.UpdatedAt);
}