using Backend.Application.DTOs;
using Backend.Application.Services;
using Microsoft.AspNetCore.Mvc;

namespace Backend.Api.Controllers;

[ApiController]
[Route("api/v1/auth")]
public class AuthController(IAuthService authService) : ControllerBase
{
    [HttpPost("sign-up")]
    public async Task<IActionResult> SignUp([FromBody] SignUpRequest request, CancellationToken ct)
    {
        var response = await authService.SignUpAsync(request, ct);
        return StatusCode(StatusCodes.Status201Created,
            new ApiResponse<AuthResponse>("Sign up successfully", response));
    }

    [HttpPost("sign-in")]
    public async Task<IActionResult> SignIn([FromBody] SignInRequest request, CancellationToken ct)
    {
        var response = await authService.SignInAsync(request, ct);
        return Ok(new ApiResponse<AuthResponse>("Sign in successfully", response));
    }
}