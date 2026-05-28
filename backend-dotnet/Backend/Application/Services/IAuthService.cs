using Backend.Application.DTOs;

namespace Backend.Application.Services;

public interface IAuthService
{
    Task<AuthResponse> SignUpAsync(SignUpRequest request, CancellationToken ct = default);
    Task<AuthResponse> SignInAsync(SignInRequest request, CancellationToken ct = default);
}