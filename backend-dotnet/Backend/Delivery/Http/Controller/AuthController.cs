using Backend.Application.DTO.Auth;
using Backend.Application.Service;
using Backend.Delivery.Http.Payload;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;

namespace Backend.Delivery.Http.Controller;


[ApiController]
[Route("api/v1/auth")]
[AllowAnonymous]
public class AuthController : ControllerBase
{
    private readonly IAuthService _authService;

    public AuthController(IAuthService authService)
    {
        _authService = authService;
    }

    [HttpPost("sign-up")]
    public async Task<IActionResult> SignUp([FromBody] SignUpRequest request, CancellationToken ct)
    {
        var auth = await _authService.SignUpAsync(request, ct);
        return StatusCode(StatusCodes.Status201Created,
            new ApiResponse<AuthResponse>("Signed up successfully", auth));
    }

    [HttpPost("sign-in")]
    public async Task<IActionResult> SignIn([FromBody] SignInRequest request, CancellationToken ct)
    {
        var auth = await _authService.SignInAsync(request, ct);
        return Ok(new ApiResponse<AuthResponse>("Signed in successfully", auth));
    }
}