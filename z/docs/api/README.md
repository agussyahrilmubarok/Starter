# API Documentation

## Base URL

```txt
http://localhost:8080/api/v1
```

## Content Type

```http
Content-Type: application/json
```

## Authorization

Use Bearer Token authentication for protected endpoints.

```http
Authorization: Bearer <token>
```

---

# Response Format

## Success Response

```json
{
  "message": "success message",
  "data": {}
}
```

## Error Response

```json
{
  "success": false,
  "message": "invalid request",
  "errors": {
    "error": "error"
  }
}
```

---

# AUTH

## Sign Up

Create a new user account.

### Endpoint

```http
POST /auth/sign-up
```

### Request Body

```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "P@ssw0rd"
}
```

### Success Response

```http
201 Created
```

```json
{
  "message": "signup success",
  "data": {
    "id": "usr_xxx",
    "name": "John Doe",
    "email": "john@example.com",
    "created_at": "2026-05-21T10:00:00Z",
    "updated_at": "2026-05-21T10:00:00Z"
  }
}
```

### Error Response

```http
400 Bad Request
```

```json
{
  "success": false,
  "message": "invalid request",
  "errors": {
    "email": "email already exists"
  }
}
```

---

## Sign In

Authenticate user and return access token.

### Endpoint

```http
POST /auth/sign-in
```

### Request Body

```json
{
  "email": "john@example.com",
  "password": "P@ssw0rd"
}
```

### Success Response

```http
200 OK
```

```json
{
  "message": "signin success",
  "data": {
    "access_token": "your-jwt-token",
    "user": {
      "id": "usr_xxx",
      "name": "John Doe",
      "email": "john@example.com",
      "created_at": "2026-05-21T10:00:00Z",
      "updated_at": "2026-05-21T10:00:00Z"
    }
  }
}
```

### Error Response

```http
401 Unauthorized
```

```json
{
  "success": false,
  "message": "invalid credentials",
  "errors": {
    "email": "email or password is incorrect"
  }
}
```

---

# USERS

> All Users endpoints require Bearer Token authentication.

---

## Get All Users

Retrieve all users.

### Endpoint

```http
GET /users
```

### Headers

```http
Authorization: Bearer <token>
```

### Success Response

```http
200 OK
```

```json
{
  "message": "get users success",
  "data": [
    {
      "id": "usr_1",
      "name": "John Doe",
      "email": "john@example.com",
      "created_at": "2026-05-21T10:00:00Z",
      "updated_at": "2026-05-21T10:00:00Z"
    }
  ]
}
```

### Error Response

```http
401 Unauthorized
```

```json
{
  "success": false,
  "message": "unauthorized",
  "errors": {
    "token": "invalid token"
  }
}
```

---

## Get User By ID

Retrieve a user by ID.

### Endpoint

```http
GET /users/{id}
```

### Example

```http
GET /users/usr_1
```

### Headers

```http
Authorization: Bearer <token>
```

### Success Response

```http
200 OK
```

```json
{
  "message": "get user success",
  "data": {
    "id": "usr_1",
    "name": "John Doe",
    "email": "john@example.com",
    "created_at": "2026-05-21T10:00:00Z",
    "updated_at": "2026-05-21T10:00:00Z"
  }
}
```

### Error Response

```http
404 Not Found
```

```json
{
  "success": false,
  "message": "user not found",
  "errors": {
    "id": "user does not exist"
  }
}
```

---

## Create User

Create a new user.

### Endpoint

```http
POST /users
```

### Headers

```http
Authorization: Bearer <token>
```

### Request Body

```json
{
  "name": "Jane Doe",
  "email": "jane@example.com",
  "password": "P@ssw0rd"
}
```

### Success Response

```http
201 Created
```

```json
{
  "message": "create user success",
  "data": {
    "id": "usr_2",
    "name": "Jane Doe",
    "email": "jane@example.com",
    "created_at": "2026-05-21T10:00:00Z",
    "updated_at": "2026-05-21T10:00:00Z"
  }
}
```

### Error Response

```http
400 Bad Request
```

```json
{
  "success": false,
  "message": "invalid request",
  "errors": {
    "email": "email already exists"
  }
}
```

---

## Update User By ID

Update a user by ID.

### Endpoint

```http
PUT /users/{id}
```

### Headers

```http
Authorization: Bearer <token>
```

### Request Body

```json
{
  "name": "John Doe Updated"
}
```

### Success Response

```http
200 OK
```

```json
{
  "message": "update user success",
  "data": {
    "id": "usr_1",
    "name": "John Doe Updated",
    "email": "john@example.com",
    "created_at": "2026-05-21T10:00:00Z",
    "updated_at": "2026-05-21T11:00:00Z"
  }
}
```

### Error Response

```http
404 Not Found
```

```json
{
  "success": false,
  "message": "user not found",
  "errors": {
    "id": "user does not exist"
  }
}
```

---

## Delete User By ID

Delete a user by ID.

### Endpoint

```http
DELETE /users/{id}
```

### Headers

```http
Authorization: Bearer <token>
```

### Success Response

```http
200 OK
```

```json
{
  "message": "delete user success",
  "data": null
}
```

### Error Response

```http
404 Not Found
```

```json
{
  "success": false,
  "message": "user not found",
  "errors": {
    "id": "user does not exist"
  }
}
```

---

# References

## User Model

```json
{
  "id": "string",
  "name": "string",
  "email": "string",
  "created_at": "datetime",
  "updated_at": "datetime"
}
```

## HTTP Status Codes

| Code | Description           |
| ---- | --------------------- |
| 200  | OK                    |
| 201  | Created               |
| 400  | Bad Request           |
| 401  | Unauthorized          |
| 404  | Not Found             |
| 500  | Internal Server Error |
