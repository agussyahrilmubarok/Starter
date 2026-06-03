using Backend.Application.DTO.User;

namespace Backend.Application.Service;

public interface IUserService
{
    Task<IEnumerable<UserResponse>> GetAllAsync(CancellationToken ct = default);
    Task<UserResponse> GetByIdAsync(Guid id, CancellationToken ct = default);
    Task<UserResponse> CreateAsync(CreateUserRequest request, CancellationToken ct = default);
    Task<UserResponse> UpdateAsync(Guid id, UpdateUserRequest request, CancellationToken ct = default);
    Task DeleteAsync(Guid id, CancellationToken ct = default);
}