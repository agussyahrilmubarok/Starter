# Register

```bash
curl -X POST http://localhost:3000/api/v2/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "username": "john.doe",
    "email": "john.doe@mail.com",
    "password": "P@ssw0rd!"
  }'
```

# Login

```bash
curl -X POST http://localhost:3000/api/v2/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "john.doe",
    "password": "P@ssw0rd!"
  }'
```

# Find Users

```bash
curl -X GET http://localhost:3000/api/v2/users \
  -H "Content-Type: application/json" 
```

# Create User

```bash
curl -X POST http://localhost:3000/api/v2/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "username": "john.doe",
    "email": "john.doe@mail.com",
    "password": "P@ssw0rd!"
  }'
```

# Find User By ID

```bash
curl -X GET http://localhost:3000/api/v2/users/1 \
  -H "Content-Type: application/json" 
```

# Update User

```bash
curl -X PUT http://localhost:3000/api/v2/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe Update",
    "username": "john.doe",
    "email": "john.doe@mail.com",
    "password": "P@ssw0rd!"
  }'
```

# Delete User By ID

```bash
curl -X DELETE http://localhost:3000/api/v1/users/1 \
  -H "Content-Type: application/json" 
```
