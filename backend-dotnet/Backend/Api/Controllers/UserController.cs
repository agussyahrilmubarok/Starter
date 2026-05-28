using Backend.Application.DTOs;
using Microsoft.AspNetCore.Mvc;

namespace Backend.Api.Controllers;

[ApiController]
[Route("api/v1/users")]
public class UserController : ControllerBase
{
    private static readonly List<UserResponse> _users =
    [
        new UserResponse("user-id-1", "John Doe",  "john@mail.com",  DateTimeOffset.UtcNow, DateTimeOffset.UtcNow),
        new UserResponse("user-id-2", "Jane Doe",  "jane@mail.com",  DateTimeOffset.UtcNow, DateTimeOffset.UtcNow),
    ];

    [HttpGet]
    public async Task<IActionResult> GetAll()
    {
        return Ok(new ApiResponse<List<UserResponse>>("Users fetched successfully", _users));
    }

    [HttpGet("{id}")]
    public async Task<IActionResult> GetById(string id)
    {
        var user = _users.FirstOrDefault(u => u.Id == id);
        if (user is null)
            return NotFound(new ApiResponse<object>("User not found", (object)null!));

        return Ok(new ApiResponse<UserResponse>("User fetched successfully", user));
    }

    [HttpPost]
    public async Task<IActionResult> Create([FromBody] CreateUserRequest request)
    {
        var user = new UserResponse(
            Guid.NewGuid().ToString(),
            request.Name,
            request.Email,
            DateTimeOffset.UtcNow,
            DateTimeOffset.UtcNow
        );

        _users.Add(user);

        return StatusCode(StatusCodes.Status201Created,
            new ApiResponse<UserResponse>("User created successfully", user));
    }

    [HttpPut("{id}")]
    public async Task<IActionResult> Update(string id, [FromBody] UpdateUserRequest request)
    {
        var index = _users.FindIndex(u => u.Id == id);
        if (index == -1)
            return NotFound(new ApiResponse<object>("User not found", (object)null!));

        var existing = _users[index];
        var updated = existing with
        {
            Name = request.Name ?? existing.Name,
            Email = request.Email ?? existing.Email,
            UpdatedAt = DateTimeOffset.UtcNow
        };

        _users[index] = updated;

        return Ok(new ApiResponse<UserResponse>("User updated successfully", updated));
    }

    [HttpDelete("{id}")]
    public async Task<IActionResult> Delete(string id)
    {
        var user = _users.FirstOrDefault(u => u.Id == id);
        if (user is null)
            return NotFound(new ApiResponse<object>("User not found", (object)null!));

        _users.Remove(user);

        return Ok(new ApiResponse<UserResponse>("User deleted successfully", user));
    }
}