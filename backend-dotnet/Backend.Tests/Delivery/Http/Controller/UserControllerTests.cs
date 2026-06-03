using System.Net;
using System.Net.Http.Headers;
using System.Net.Http.Json;
using System.Security.Claims;
using System.Text.Encodings.Web;
using System.Text.Json;
using Backend.Application.DTO.User;
using Backend.Application.Service;
using Backend.Common.Exceptions;
using Backend.Delivery.Http.Payload;
using Microsoft.AspNetCore.Authentication;
using Microsoft.AspNetCore.Hosting;
using Microsoft.AspNetCore.Mvc.Testing;
using Microsoft.AspNetCore.TestHost;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Logging;
using Microsoft.Extensions.Options;
using Moq;

namespace Backend.Tests.Delivery.Http.Controller;

public class FakeAuthHandler : AuthenticationHandler<AuthenticationSchemeOptions>
{
    public FakeAuthHandler(
        IOptionsMonitor<AuthenticationSchemeOptions> options,
        ILoggerFactory logger,
        UrlEncoder encoder) : base(options, logger, encoder) { }

    protected override Task<AuthenticateResult> HandleAuthenticateAsync()
    {
        var claims = new[] { new Claim("sub", Guid.NewGuid().ToString()) };
        var identity = new ClaimsIdentity(claims, "FakeScheme");
        var principal = new ClaimsPrincipal(identity);
        var ticket = new AuthenticationTicket(principal, "FakeScheme");
        return Task.FromResult(AuthenticateResult.Success(ticket));
    }
}

public class AuthenticatedFactory : WebApplicationFactory<Program>
{
    private readonly Action<IServiceCollection>? _extraServices;

    public AuthenticatedFactory(Action<IServiceCollection>? extraServices = null)
    {
        _extraServices = extraServices;
    }

    protected override void ConfigureWebHost(IWebHostBuilder builder)
    {
        builder.UseEnvironment("Testing");

        builder.ConfigureTestServices(services =>
        {
            var authDescriptors = services
                .Where(d => d.ServiceType.FullName?.Contains("Authentication") == true
                         || d.ServiceType.FullName?.Contains("JwtBearer") == true)
                .ToList();

            foreach (var d in authDescriptors)
                services.Remove(d);

            services.AddAuthentication("FakeScheme")
                .AddScheme<AuthenticationSchemeOptions, FakeAuthHandler>(
                    "FakeScheme", _ => { });

            _extraServices?.Invoke(services);
        });
    }
}

public class UserControllerTests : IDisposable
{
    private readonly AuthenticatedFactory _factory;
    private readonly HttpClient _client;
    private readonly Mock<IUserService> _serviceMock;

    private static readonly JsonSerializerOptions JsonOptions = new()
    {
        PropertyNameCaseInsensitive = true
    };

    public UserControllerTests()
    {
        _serviceMock = new Mock<IUserService>();

        _factory = new AuthenticatedFactory(services =>
        {
            var descriptor = services.SingleOrDefault(
                d => d.ServiceType == typeof(IUserService));
            if (descriptor != null)
                services.Remove(descriptor);

            services.AddSingleton(_serviceMock.Object);
        });

        _client = _factory.CreateClient();
        _client.DefaultRequestHeaders.Authorization =
            new AuthenticationHeaderValue("FakeScheme", "fake-token");
    }

    public void Dispose()
    {
        _client.Dispose();
        _factory.Dispose();
    }

    private static UserResponse MakeUserResponse(
        string name = "John Doe",
        string email = "john@mail.com")
        => new(Guid.NewGuid(), name, email, DateTime.UtcNow, DateTime.UtcNow);

    [Fact]
    public async Task GetAll_ShouldReturn200WithUsers()
    {
        var users = new List<UserResponse>
        {
            MakeUserResponse("Alice", "alice@mail.com"),
            MakeUserResponse("Bob",   "bob@mail.com"),
        };

        _serviceMock
            .Setup(s => s.GetAllAsync(It.IsAny<CancellationToken>()))
            .ReturnsAsync(users);

        var response = await _client.GetAsync("/api/v1/users");
        var json = await response.Content.ReadAsStringAsync();
        var body = JsonSerializer.Deserialize<ApiResponse<List<UserResponse>>>(json, JsonOptions);

        Assert.Equal(HttpStatusCode.OK, response.StatusCode);
        Assert.NotNull(body);
        Assert.Equal("Get all users successfully", body.Message);
        Assert.Equal(2, body.Data!.Count);
    }

    [Fact]
    public async Task GetAll_WhenEmpty_ShouldReturn200WithEmptyList()
    {
        _serviceMock
            .Setup(s => s.GetAllAsync(It.IsAny<CancellationToken>()))
            .ReturnsAsync(new List<UserResponse>());

        var response = await _client.GetAsync("/api/v1/users");
        var json = await response.Content.ReadAsStringAsync();
        var body = JsonSerializer.Deserialize<ApiResponse<List<UserResponse>>>(json, JsonOptions);

        Assert.Equal(HttpStatusCode.OK, response.StatusCode);
        Assert.Empty(body!.Data!);
    }

    [Fact]
    public async Task GetById_WhenUserExists_ShouldReturn200()
    {
        var user = MakeUserResponse();

        _serviceMock
            .Setup(s => s.GetByIdAsync(user.Id, It.IsAny<CancellationToken>()))
            .ReturnsAsync(user);

        var response = await _client.GetAsync($"/api/v1/users/{user.Id}");
        var json = await response.Content.ReadAsStringAsync();
        var body = JsonSerializer.Deserialize<ApiResponse<UserResponse>>(json, JsonOptions);

        Assert.Equal(HttpStatusCode.OK, response.StatusCode);
        Assert.Equal(user.Id, body!.Data!.Id);
        Assert.Equal(user.Email, body.Data.Email);
    }

    [Fact]
    public async Task GetById_WhenUserNotFound_ShouldReturn404()
    {
        var randomId = Guid.NewGuid();

        _serviceMock
            .Setup(s => s.GetByIdAsync(randomId, It.IsAny<CancellationToken>()))
            .ThrowsAsync(new NotFoundException($"User with id {randomId} not found."));

        var response = await _client.GetAsync($"/api/v1/users/{randomId}");

        Assert.Equal(HttpStatusCode.NotFound, response.StatusCode);
    }

    [Fact]
    public async Task Create_WhenValidRequest_ShouldReturn201()
    {
        var request = new CreateUserRequest("John Doe", "john@mail.com", "password123");
        var user = MakeUserResponse(request.Name, request.Email);

        _serviceMock
            .Setup(s => s.CreateAsync(It.IsAny<CreateUserRequest>(), It.IsAny<CancellationToken>()))
            .ReturnsAsync(user);

        var response = await _client.PostAsJsonAsync("/api/v1/users", request);
        var json = await response.Content.ReadAsStringAsync();
        var body = JsonSerializer.Deserialize<ApiResponse<UserResponse>>(json, JsonOptions);

        Assert.Equal(HttpStatusCode.Created, response.StatusCode);
        Assert.Equal(request.Email, body!.Data!.Email);
    }

    [Fact]
    public async Task Create_WhenEmailAlreadyExists_ShouldReturn409()
    {
        var request = new CreateUserRequest("John Doe", "john@mail.com", "password123");

        _serviceMock
            .Setup(s => s.CreateAsync(It.IsAny<CreateUserRequest>(), It.IsAny<CancellationToken>()))
            .ThrowsAsync(new EmailAlreadyExistsException(request.Email));

        var response = await _client.PostAsJsonAsync("/api/v1/users", request);

        Assert.Equal(HttpStatusCode.Conflict, response.StatusCode);
    }

    [Fact]
    public async Task Create_WhenInvalidRequest_ShouldReturn400()
    {
        var request = new { name = "J", email = "not-an-email", password = "123" };

        var response = await _client.PostAsJsonAsync("/api/v1/users", request);

        Assert.Equal(HttpStatusCode.BadRequest, response.StatusCode);
    }

    [Fact]
    public async Task Update_WhenUserExists_ShouldReturn200()
    {
        var id = Guid.NewGuid();
        var request = new UpdateUserRequest("Updated Name", null, null);
        var updated = new UserResponse(id, "Updated Name", "john@mail.com", DateTime.UtcNow, DateTime.UtcNow);

        _serviceMock
            .Setup(s => s.UpdateAsync(id, It.IsAny<UpdateUserRequest>(), It.IsAny<CancellationToken>()))
            .ReturnsAsync(updated);

        var response = await _client.PutAsJsonAsync($"/api/v1/users/{id}", request);
        var json = await response.Content.ReadAsStringAsync();
        var body = JsonSerializer.Deserialize<ApiResponse<UserResponse>>(json, JsonOptions);

        Assert.Equal(HttpStatusCode.OK, response.StatusCode);
        Assert.Equal("Updated Name", body!.Data!.Name);
    }

    [Fact]
    public async Task Update_WhenUserNotFound_ShouldReturn404()
    {
        var randomId = Guid.NewGuid();

        _serviceMock
            .Setup(s => s.UpdateAsync(randomId, It.IsAny<UpdateUserRequest>(), It.IsAny<CancellationToken>()))
            .ThrowsAsync(new NotFoundException($"User with id {randomId} not found."));

        var response = await _client.PutAsJsonAsync($"/api/v1/users/{randomId}",
            new UpdateUserRequest(null, null, null));

        Assert.Equal(HttpStatusCode.NotFound, response.StatusCode);
    }

    [Fact]
    public async Task Delete_WhenUserExists_ShouldReturn200()
    {
        var id = Guid.NewGuid();

        _serviceMock
            .Setup(s => s.DeleteAsync(id, It.IsAny<CancellationToken>()))
            .Returns(Task.CompletedTask);

        var response = await _client.DeleteAsync($"/api/v1/users/{id}");

        Assert.Equal(HttpStatusCode.OK, response.StatusCode);
    }

    [Fact]
    public async Task Delete_WhenUserNotFound_ShouldReturn404()
    {
        var randomId = Guid.NewGuid();

        _serviceMock
            .Setup(s => s.DeleteAsync(randomId, It.IsAny<CancellationToken>()))
            .ThrowsAsync(new NotFoundException($"User with id {randomId} not found."));

        var response = await _client.DeleteAsync($"/api/v1/users/{randomId}");

        Assert.Equal(HttpStatusCode.NotFound, response.StatusCode);
    }

    [Fact]
    public async Task GetAll_WithoutToken_ShouldReturn401()
    {
        var realFactory = new WebApplicationFactory<Program>();
        var unauthClient = realFactory.CreateClient();

        var response = await unauthClient.GetAsync("/api/v1/users");

        Assert.Equal(HttpStatusCode.Unauthorized, response.StatusCode);

        unauthClient.Dispose();
        realFactory.Dispose();
    }
}