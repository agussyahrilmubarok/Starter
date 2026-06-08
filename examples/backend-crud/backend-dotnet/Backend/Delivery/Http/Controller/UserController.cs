using Backend.Application.DTO.User;
using Backend.Application.Service;
using Backend.Delivery.Http.Payload;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;

namespace Backend.Delivery.Http.Controller;


[ApiController]
[Route("api/v1/users")]
[Authorize]
public class UserController : ControllerBase
{
    private readonly IUserService _userService;

    public UserController(IUserService userService)
    {
        _userService = userService;
    }

    [HttpGet]
    public async Task<IActionResult> GetAll(CancellationToken ct)
    {
        var users = await _userService.GetAllAsync(ct);
        return Ok(new ApiResponse<IEnumerable<UserResponse>>("Users retrieved successfully", users));
    }

    [HttpGet("{id:guid}")]
    public async Task<IActionResult> GetById(Guid id, CancellationToken ct)
    {
        var user = await _userService.GetByIdAsync(id, ct);
        return Ok(new ApiResponse<UserResponse>("User retrieved successfully", user));
    }

    [HttpPost]
    public async Task<IActionResult> Create([FromBody] CreateUserRequest request, CancellationToken ct)
    {
        var user = await _userService.CreateAsync(request, ct);
        return StatusCode(StatusCodes.Status201Created,
            new ApiResponse<UserResponse>("User created successfully", user));
    }

    [HttpPut("{id:guid}")]
    public async Task<IActionResult> Update(Guid id, [FromBody] UpdateUserRequest request, CancellationToken ct)
    {
        var user = await _userService.UpdateAsync(id, request, ct);
        return Ok(new ApiResponse<UserResponse>("User updated successfully", user));
    }

    [HttpDelete("{id:guid}")]
    public async Task<IActionResult> Delete(Guid id, CancellationToken ct)
    {
        await _userService.DeleteAsync(id, ct);
        return Ok(new ApiResponse<object>("User deleted successfully", new { id }));
    }
}