using Backend.Application.DTO.User;
using Backend.Common.Exceptions;
using Backend.Domain.User;

namespace Backend.Application.Service;

public class UserService : IUserService
{
    private readonly IUserRepository _userRepository;
    private readonly ILogger<UserService> _logger;

    public UserService(IUserRepository userRepository, ILogger<UserService> logger)
    {
        _userRepository = userRepository;
        _logger = logger;
    }

    public async Task<IEnumerable<UserResponse>> GetAllAsync(CancellationToken ct = default)
    {
        var users = await _userRepository.FindAllAsync(ct);
        _logger.LogInformation("Users fetched {Count}", users.Count());
        return users.Select(ToResponse);
    }

    public async Task<UserResponse> GetByIdAsync(Guid id, CancellationToken ct = default)
    {
        var user = await _userRepository.FindByIdAsync(id, ct);
        if (user is null)
        {
            _logger.LogWarning("User not found");
            throw new NotFoundException($"User not found");
        }

        _logger.LogInformation("User fetched {UserId}", user.Id);
        return ToResponse(user);
    }

    public async Task<UserResponse> CreateAsync(CreateUserRequest request, CancellationToken ct = default)
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

        _logger.LogInformation("User created {UserId}", user.Id);
        return ToResponse(user);
    }

    public async Task<UserResponse> UpdateAsync(Guid id, UpdateUserRequest request, CancellationToken ct = default)
    {
        var user = await _userRepository.FindByIdAsync(id, ct);
        if (user is null)
        {
            _logger.LogWarning("User not found");
            throw new NotFoundException($"User not found");
        }

        if (!string.IsNullOrWhiteSpace(request.Name))
            user.UpdateName(request.Name);

        if (!string.IsNullOrWhiteSpace(request.Email) && user.Email != request.Email)
        {
            var emailExists = await _userRepository.ExistsByEmailAsync(request.Email, ct);
            if (emailExists)
            {
                _logger.LogWarning("The email has already been taken {Email}", request.Email);
                throw new EmailAlreadyExistsException();
            }
            user.UpdateEmail(request.Email.ToLowerInvariant());
        }

        if (!string.IsNullOrWhiteSpace(request.Password))
            user.UpdatePassword(BCrypt.Net.BCrypt.HashPassword(request.Password));

        await _userRepository.UpdateAsync(user, ct);

        _logger.LogInformation("User updated {UserId}", user.Id);
        return ToResponse(user);
    }

    public async Task DeleteAsync(Guid id, CancellationToken ct = default)
    {
        var user = await _userRepository.FindByIdAsync(id, ct);
        if (user is null)
        {
            _logger.LogWarning("User not found");
            throw new NotFoundException($"User not found");
        }

        await _userRepository.DeleteAsync(user.Id, ct);
        _logger.LogInformation("User deleted {UserId}", id);
    }

    private static UserResponse ToResponse(User user) => new(
        user.Id,
        user.Name,
        user.Email,
        user.CreatedAt,
        user.UpdatedAt
    );
}