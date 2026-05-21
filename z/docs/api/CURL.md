# CURLs

---

# AUTH

## Sign Up

Create new user account.

```bash
curl -X POST http://localhost:8080/api/v1/auth/sign-up \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "P@ssw0rd"
  }'
```

## Sign In

Authenticate user account.

```bash
curl -X POST http://localhost:8080/api/v1/auth/sign-in \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "P@ssw0rd"
  }'
```