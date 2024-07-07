# Building an API

## Running the Program

To run the program, use the following command:

```bash
go run cmd/api/main.go
```

#### API POST Route

This API POST route requires an `Authorization` header for authentication.

###### Example

For the username `alex`, the Authorization header should be `ABC123`.

```
POST http://localhost:8000/account/coins/?username=alex
Authorization: ABC123
```

You can find dummy data in the /internal/tools/mockdb.go file.