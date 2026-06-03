namespace Backend.Domain.User;

public class User
{
    public Guid Id { get; private set; }
    public string Name { get; private set; } = null!;
    public string Email { get; private set; } = null!;
    public string Password { get; private set; } = null!;
    public DateTime CreatedAt { get; private set; }
    public DateTime UpdatedAt { get; private set; }

    private User() { }

    public static User Create(string name, string email, string hashedPassword)
    {
        if (string.IsNullOrWhiteSpace(name))
            throw new ArgumentException("Name is required.", nameof(name));

        if (string.IsNullOrWhiteSpace(email))
            throw new ArgumentException("Email is required.", nameof(email));

        if (string.IsNullOrWhiteSpace(hashedPassword))
            throw new ArgumentException("Password is required.", nameof(hashedPassword));

        var now = DateTime.UtcNow;

        return new User
        {
            Id = Guid.NewGuid(),
            Name = name,
            Email = email,
            Password = hashedPassword,
            CreatedAt = now,
            UpdatedAt = now
        };
    }

    public void UpdateName(string name)
    {
        if (string.IsNullOrWhiteSpace(name))
            throw new ArgumentException("Name is required.", nameof(name));

        Name = name;
        UpdatedAt = DateTime.UtcNow;
    }

    public void UpdateEmail(string email)
    {
        if (string.IsNullOrWhiteSpace(email))
            throw new ArgumentException("Email is required.", nameof(email));

        Email = email;
        UpdatedAt = DateTime.UtcNow;
    }

    public void UpdatePassword(string hashedPassword)
    {
        if (string.IsNullOrWhiteSpace(hashedPassword))
            throw new ArgumentException("Password is required.", nameof(hashedPassword));

        Password = hashedPassword;
        UpdatedAt = DateTime.UtcNow;
    }
}