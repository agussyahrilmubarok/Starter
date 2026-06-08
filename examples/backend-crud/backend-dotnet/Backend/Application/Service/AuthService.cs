using Backend.Application.DTO.Auth;
using Backend.Application.DTO.User;
using Backend.Common.Exceptions;
using Backend.Domain.User;
using Backend.Infrastructure.Security;

namespace Backend.Application.Service;

public class AuthService : IAuthService
{
    private readonly IUserRepository _userRepository;
    private readonly IJwtManager _jwtManager;
    private readonly ILogger<AuthService> _logger;

    public AuthService(IUserRepository userRepository, IJwtManager jwtManager, ILogger<AuthService> logger)
    {
        _userRepository = userRepository;
        _jwtManager = jwtManager;
        _logger = logger;
    }

    public async Task<AuthResponse> SignUpAsync(SignUpRequest request, CancellationToken ct = default)
    {
        var alreadyExists = await _userRepository.ExistsByEmailAsync(request.Email, ct);
        if (alreadyExists)
        {
            _logger.LogWarning("The email has already been taken");
            throw new EmailAlreadyExistsException();
        }

        var hashedPassword = BCrypt.Net.BCrypt.HashPassword(request.Password);
        var user = User.Create(request.Name, request.Email, hashedPassword);

        await _userRepository.CreateAsync(user, ct);

        var token = _jwtManager.GenerateToken(user.Id);

        _logger.LogInformation("User signed up {UserId}", user.Id);
        return ToResponse(token, user);
    }

    public async Task<AuthResponse> SignInAsync(SignInRequest request, CancellationToken ct = default)
    {
        var user = await _userRepository.FindByEmailAsync(request.Email, ct);
        if (user is null)
        {
            _logger.LogWarning("Email not registered");
            throw new EmailNotRegisteredException();
        }

        var isPasswordValid = BCrypt.Net.BCrypt.Verify(request.Password, user.Password);
        if (!isPasswordValid)
        {
            _logger.LogWarning("Invalid password attempt for user {UserId}", user.Id);
            throw new WrongPasswordException();
        }

        var token = _jwtManager.GenerateToken(user.Id);

        _logger.LogInformation("User signed in {UserId}", user.Id);
        return ToResponse(token, user);
    }

    private static AuthResponse ToResponse(string token, User user) => new(
        token,
        ToUserResponse(user)
    );

    private static UserResponse ToUserResponse(User user) => new(
        user.Id,
        user.Name,
        user.Email,
        user.CreatedAt,
        user.UpdatedAt
    );
}