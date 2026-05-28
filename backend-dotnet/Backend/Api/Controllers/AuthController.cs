using Backend.Application.DTOs;
using Microsoft.AspNetCore.Mvc;

namespace Backend.Api.Controllers;

[ApiController]
[Route("api/v1/auth")]
public class AuthController : ControllerBase
{
    [HttpPost("sign-up")]
    public async Task<IActionResult> SignUp([FromBody] SignUpRequest request)
    {
        var userResponse = new AuthUserResponse(
            "user-id",
            request.Name,
            request.Email,
            DateTimeOffset.UtcNow,
            DateTimeOffset.UtcNow
        );

        var response = new AuthResponse("jwt_token_secret_xyz123", userResponse);

        return StatusCode(StatusCodes.Status201Created,
            new ApiResponse<AuthResponse>("Sign up successfully", response));
    }

    [HttpPost("sign-in")]
    public async Task<IActionResult> SignIn([FromBody] SignInRequest request)
    {
        var userResponse = new AuthUserResponse(
            "user-id",
            "name",
            request.Email,
            DateTimeOffset.UtcNow,
            DateTimeOffset.UtcNow
        );

        var response = new AuthResponse("jwt_token_secret_xyz123", userResponse);

        return Ok(new ApiResponse<AuthResponse>("Sign in successfully", response));
    }
}