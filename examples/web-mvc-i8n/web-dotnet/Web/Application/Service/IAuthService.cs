using Web.Application.DTO.Auth;
using Web.Application.DTO.User;

namespace Web.Application.Service;

public interface IAuthService
{
    Task<UserResponse> SignUpAsync(SignUpRequest request, CancellationToken ct = default);
    Task<UserResponse> SignInAsync(SignInRequest request, CancellationToken ct = default);
}