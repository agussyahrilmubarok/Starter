# API

Base URL

```txt
/api/v1
````

Content-Type

```http
Content-Type: application/json
```

---

# AUTH

## Sign Up

Create new user account.

### Endpoint

```http
POST /auth/signup
```

### Request Body

```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "secret123"
}
```

### Success Response

```http
201 Created
```

```json
{
  "success": true,
  "message": "user created",
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
    "error": "error"
  }
}
```

---

## Sign In

Authenticate user and return access token.

### Endpoint

```http
POST /auth/signin
```

### Request Body

```json
{
  "email": "john@example.com",
  "password": "secret123"
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
    "access_token": "jwt_token_here",
    "token_type": "Bearer"
  }
}
```

### Error Response

```http
401 Unauthorized
```

```json
{
  "message": "invalid credentials"
}
```

---

# USERS

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

---

## Get User By ID

Retrieve single user by id.

### Endpoint

```http
GET /users/{id}
```

### Example

```http
GET /users/usr_1
```

### Success Response

```http
200 OK
```

```json
{
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
  "message": "user not found"
}
```

---

## Create User

Create new user.

### Endpoint

```http
POST /users
```

### Request Body

```json
{
  "name": "Jane Doe",
  "email": "jane@example.com",
  "password": "secret123"
}
```

### Success Response

```http
201 Created
```

```json
{
  "message": "user created",
  "data": {
    "id": "usr_2",
    "name": "Jane Doe",
    "email": "jane@example.com"
  }
}
```

---

## Update User

Update existing user.

### Endpoint

```http
PUT /users/{id}
```

### Request Body

```json
{
  "name": "Jane Updated",
  "email": "janeupdated@example.com"
}
```

### Success Response

```http
200 OK
```

```json
{
  "message": "user updated"
}
```

---

## Delete User

Delete user by id.

### Endpoint

```http
DELETE /users/{id}
```

### Success Response

```http
204 No Content
```

---

# References

## User Model

```json
{
  "id": "string",
  "name": "string",
  "email": "string",
  "password": "string",
  "created_at": "datetime",
  "updated_at": "datetime"
}
```

## HTTP Status Codes

| Code | Description           |
| ---- | --------------------- |
| 200  | OK                    |
| 201  | Created               |
| 204  | No Content            |
| 400  | Bad Request           |
| 401  | Unauthorized          |
| 404  | Not Found             |
| 500  | Internal Server Error |

## Authorization

Use Bearer Token authentication.

```http
Authorization: Bearer <token>
```

