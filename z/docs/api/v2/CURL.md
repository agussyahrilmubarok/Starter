# CURLs

---

## AUTH

### SIGN UP

```bash
curl -X POST http://localhost:8080/api/v2/auth/sign-up \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "P@ssw0rd"
  }'
```

### SIGN IN

```bash
curl -X POST http://localhost:8080/api/v2/auth/sign-in \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "P@ssw0rd"
  }'
```

---

## USERS

### GET ALL USERS

```bash
curl -X GET http://localhost:8080/api/v2/users \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json"
```

### GET USER BY ID

```bash
curl -X GET http://localhost:8080/api/v2/users/{id} \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json"
```

### CREATE USER

```bash
curl -X POST http://localhost:8080/api/v2/users \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "P@ssw0rd"
  }'
```

### UPDATE USER BY ID

```bash
curl -X PUT http://localhost:8080/api/v2/users/{id}
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe Updated"
  }'
```

### DELETE USER BY ID

```bash
curl -X DELETE http://localhost:8080/api/v2/users/{id} \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json"
```
