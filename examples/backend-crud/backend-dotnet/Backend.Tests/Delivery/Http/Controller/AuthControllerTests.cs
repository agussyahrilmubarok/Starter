using System.Net;
using System.Net.Http.Json;
using System.Text.Json;
using Backend.Application.DTO.Auth;
using Backend.Application.DTO.User;
using Backend.Application.Service;
using Backend.Common.Exceptions;
using Backend.Delivery.Http.Payload;
using Microsoft.AspNetCore.Mvc.Testing;
using Microsoft.AspNetCore.TestHost;
using Microsoft.Extensions.DependencyInjection;
using Moq;

namespace Backend.Tests.Delivery.Http.Controller;

public class AuthControllerTests : IDisposable
{
    private readonly WebApplicationFactory<Program> _factory;
    private readonly HttpClient _client;
    private readonly Mock<IAuthService> _serviceMock;

    private static readonly JsonSerializerOptions JsonOptions = new()
    {
        PropertyNameCaseInsensitive = true
    };

    public AuthControllerTests()
    {
        _serviceMock = new Mock<IAuthService>();

        _factory = new WebApplicationFactory<Program>()
            .WithWebHostBuilder(builder =>
            {
                builder.ConfigureTestServices(services =>
                {
                    var descriptor = services.SingleOrDefault(
                        d => d.ServiceType == typeof(IAuthService));
                    if (descriptor != null)
                        services.Remove(descriptor);

                    services.AddSingleton(_serviceMock.Object);
                });
            });

        _client = _factory.CreateClient();
    }

    public void Dispose()
    {
        _client.Dispose();
        _factory.Dispose();
    }

    private static AuthResponse MakeAuthResponse(string email = "john@mail.com")
        => new(
            "fake.jwt.token",
            new UserResponse(Guid.NewGuid(), "John Doe", email, DateTime.UtcNow, DateTime.UtcNow)
        );

    [Fact]
    public async Task SignUp_WhenValidRequest_ShouldReturn201()
    {
        var request = new SignUpRequest("John Doe", "john@mail.com", "password123");
        var authResponse = MakeAuthResponse(request.Email);

        _serviceMock
            .Setup(s => s.SignUpAsync(It.IsAny<SignUpRequest>(), It.IsAny<CancellationToken>()))
            .ReturnsAsync(authResponse);

        var response = await _client.PostAsJsonAsync("/api/v1/auth/sign-up", request);
        var json = await response.Content.ReadAsStringAsync();
        var body = JsonSerializer.Deserialize<ApiResponse<AuthResponse>>(json, JsonOptions);

        Assert.Equal(HttpStatusCode.Created, response.StatusCode);
        Assert.NotNull(body!.Data!.Token);
        Assert.Equal(request.Email, body.Data.User.Email);
    }

    [Fact]
    public async Task SignUp_WhenEmailAlreadyExists_ShouldReturn409()
    {
        var request = new SignUpRequest("John Doe", "john@mail.com", "password123");

        _serviceMock
            .Setup(s => s.SignUpAsync(It.IsAny<SignUpRequest>(), It.IsAny<CancellationToken>()))
            .ThrowsAsync(new EmailAlreadyExistsException());

        var response = await _client.PostAsJsonAsync("/api/v1/auth/sign-up", request);

        Assert.Equal(HttpStatusCode.Conflict, response.StatusCode);
    }

    [Fact]
    public async Task SignUp_WhenInvalidRequest_ShouldReturn400()
    {
        var request = new { name = "", email = "bukan-email", password = "123" };

        var response = await _client.PostAsJsonAsync("/api/v1/auth/sign-up", request);

        Assert.Equal(HttpStatusCode.BadRequest, response.StatusCode);
    }

    [Fact]
    public async Task SignIn_WhenValidCredentials_ShouldReturn200()
    {
        var request = new SignInRequest("john@mail.com", "password123");
        var authResponse = MakeAuthResponse(request.Email);

        _serviceMock
            .Setup(s => s.SignInAsync(It.IsAny<SignInRequest>(), It.IsAny<CancellationToken>()))
            .ReturnsAsync(authResponse);

        var response = await _client.PostAsJsonAsync("/api/v1/auth/sign-in", request);
        var json = await response.Content.ReadAsStringAsync();
        var body = JsonSerializer.Deserialize<ApiResponse<AuthResponse>>(json, JsonOptions);

        Assert.Equal(HttpStatusCode.OK, response.StatusCode);
        Assert.NotNull(body!.Data!.Token);
        Assert.Equal(request.Email, body.Data.User.Email);
    }

    [Fact]
    public async Task SignIn_WhenEmailNotRegistered_ShouldReturn401()
    {
        var request = new SignInRequest("ghost@mail.com", "password123");

        _serviceMock
            .Setup(s => s.SignInAsync(It.IsAny<SignInRequest>(), It.IsAny<CancellationToken>()))
            .ThrowsAsync(new EmailNotRegisteredException());

        var response = await _client.PostAsJsonAsync("/api/v1/auth/sign-in", request);

        Assert.Equal(HttpStatusCode.Unauthorized, response.StatusCode);
    }

    [Fact]
    public async Task SignIn_WhenWrongPassword_ShouldReturn401()
    {
        var request = new SignInRequest("john@mail.com", "wrongpassword");

        _serviceMock
            .Setup(s => s.SignInAsync(It.IsAny<SignInRequest>(), It.IsAny<CancellationToken>()))
            .ThrowsAsync(new WrongPasswordException());

        var response = await _client.PostAsJsonAsync("/api/v1/auth/sign-in", request);

        Assert.Equal(HttpStatusCode.Unauthorized, response.StatusCode);
    }

    [Fact]
    public async Task SignIn_WhenInvalidRequest_ShouldReturn400()
    {
        var request = new { email = "", password = "123" };

        var response = await _client.PostAsJsonAsync("/api/v1/auth/sign-in", request);

        Assert.Equal(HttpStatusCode.BadRequest, response.StatusCode);
    }
}