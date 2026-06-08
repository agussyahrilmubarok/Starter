using Web.Application.DTO.User;
using Web.Common.Exceptions;
using Web.Domain.User;

namespace Web.Application.Service;

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
        var list = users.ToList();
        _logger.LogInformation("Users fetched {Count}", list.Count);
        return list.Select(ToResponse);
    }

    public async Task<UserResponse> GetByIdAsync(Guid id, CancellationToken ct = default)
    {
        var user = await _userRepository.FindByIdAsync(id, ct);
        if (user is null)
        {
            _logger.LogWarning("User with id {Id} not found", id);
            throw new NotFoundException("User not found");
        }

        return ToResponse(user);
    }

    public async Task<UserResponse> CreateAsync(CreateUserRequest request, CancellationToken ct = default)
    {
        if (await _userRepository.ExistsByEmailAsync(request.Email.ToLowerInvariant(), ct))
        {
            _logger.LogWarning("Email already taken: {Email}", request.Email);
            throw new EmailAlreadyExistsException();
        }

        var hashedPassword = BCrypt.Net.BCrypt.HashPassword(request.Password);
        var user = User.Create(request.Name, request.Email.ToLowerInvariant(), hashedPassword);
        await _userRepository.CreateAsync(user, ct);

        _logger.LogInformation("User created {Id}", user.Id);
        return ToResponse(user);
    }

    public async Task<UserResponse> UpdateAsync(Guid id, UpdateUserRequest request, CancellationToken ct = default)
    {
        var user = await _userRepository.FindByIdAsync(id, ct);
        if (user is null)
        {
            _logger.LogWarning("User with id {Id} not found", id);
            throw new NotFoundException("User not found");
        }

        if (!string.IsNullOrWhiteSpace(request.Name))
            user.UpdateName(request.Name);

        if (!string.IsNullOrWhiteSpace(request.Email) &&
            !user.Email.Equals(request.Email, StringComparison.OrdinalIgnoreCase))
        {
            if (await _userRepository.ExistsByEmailAsync(request.Email.ToLowerInvariant(), ct))
            {
                _logger.LogWarning("Email already taken: {Email}", request.Email);
                throw new EmailAlreadyExistsException();
            }
            user.UpdateEmail(request.Email.ToLowerInvariant());
        }

        if (!string.IsNullOrWhiteSpace(request.Password))
            user.UpdatePassword(BCrypt.Net.BCrypt.HashPassword(request.Password));

        await _userRepository.UpdateAsync(user, ct);
        _logger.LogInformation("User updated {Id}", user.Id);
        return ToResponse(user);
    }

    public async Task DeleteAsync(Guid id, CancellationToken ct = default)
    {
        var user = await _userRepository.FindByIdAsync(id, ct);
        if (user is null)
        {
            _logger.LogWarning("User with id {Id} not found", id);
            throw new NotFoundException("User not found");
        }

        await _userRepository.DeleteAsync(id, ct);
        _logger.LogInformation("User deleted {Id}", id);
    }

    private static UserResponse ToResponse(User user) =>
        new(user.Id, user.Name, user.Email, user.CreatedAt, user.UpdatedAt);
}