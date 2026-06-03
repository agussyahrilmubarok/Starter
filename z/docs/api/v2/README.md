# API Documentation

## Base URL

```txt
http://localhost:8080/api/v2
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
  "message": "Successfully",
  "data": {}
}
```

## Error Response

```json
{
  "message": "Invalid request",
  "errors": {
    "error": "Something went wrong"
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
  "message": "Signed up successfully",
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
  "message": "Invalid request",
  "errors": {
    "email": "Email already exists"
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
  "message": "Signed in successfully",
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
  "message": "Invalid credentials",
  "errors": {
    "password": "Wrong password"
  }
}
```

---

# USERS

> All Users endpoints require Bearer Token authentication.

---

## Get All Users

Retrieve users with advanced pagination, sorting, and global search.

### Endpoint

```http
GET /users?search=jane&sort=name,asc&page=1&size=10
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
  "message": "Users fetched successfully",
  "data": [
    {
      "id": "usr_1",
      "name": "John Doe",
      "email": "john@example.com",
      "created_at": "2026-05-21T10:00:00Z",
      "updated_at": "2026-05-21T10:00:00Z"
    }
  ],
  "pagination": {
    "current_page": 2,
    "page_size": 5,
    "total_items": 42,
    "total_pages": 9,
    "has_next": true,
    "has_previous": true
  }
}
```

### Error Response

```http
400 Bad Request
```

```json
{
  "message": "Invalid request",
  "errors": {
    "page": "Must be a positive integer",
    "size": "Must be between 1 and 100"
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
  "message": "User fetched successfully",
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
  "message": "User not found",
  "errors": {
    "error": "User does not exist"
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
  "message": "User created successfully",
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
  "message": "Invalid request",
  "errors": {
    "email": "Email already exists"
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
  "message": "User updated successfully",
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
  "message": "User not found",
  "errors": {
    "error": "User does not exist"
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
  "message": "user not found",
  "errors": {
    "error": "user does not exist"
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
