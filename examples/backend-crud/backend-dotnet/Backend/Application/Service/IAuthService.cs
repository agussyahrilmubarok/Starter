using Backend.Application.DTO.Auth;
using Backend.Application.DTO.User;

namespace Backend.Application.Service;

public interface IAuthService
{
    Task<AuthResponse> SignUpAsync(SignUpRequest request, CancellationToken ct = default);
    Task<AuthResponse> SignInAsync(SignInRequest request, CancellationToken ct = default);
}