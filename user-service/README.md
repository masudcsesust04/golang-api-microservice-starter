# User Service

This service provides user management functionality, including CRUD operations for users.

## Tech Details

### Dependencies

- **Go**: 1.23.2
- **Router**: `github.com/gorilla/mux`
- **Database Driver**: `github.com/jackc/pgx/v5`
- **Configuration**: `github.com/spf13/viper`
- **Validation**: `github.com/go-playground/validator/v10`
- **JWT**: `github.com/golang-jwt/jwt/v5`
- **Testing**: `github.com/stretchr/testify`
- **Cryptography**: `golang.org/x/crypto`

### Running the Server

To start the server, run:

```bash
go run cmd/server/main.go
```

The server will start on port `8081`.

---

## API Documentation

### Calling Mechanism

-   **Authenticated Routes:** For endpoints that require authentication, you must include an `Authorization` header in your request with a valid JWT access token from the `auth-service`:
    ```
    Authorization: Bearer <your_access_token>
    ```
-   **Request/Response Format:** All request and response bodies are in JSON format. Ensure your requests have the `Content-Type: application/json` header.

---

### User Management APIs

These endpoints handle CRUD operations for users.

#### 1. Get All Users

-   **Description:** Retrieves a list of all users.
-   **Method:** `GET`
-   **Path:** `/users`
-   **Authentication:** **Required**.
-   **Success Response (200 OK):** An array of user objects.

#### 2. Get User by ID

-   **Description:** Retrieves a single user by their ID.
-   **Method:** `GET`
-   **Path:** `/users/{id}` (e.g., `/users/1`)
-   **Authentication:** **Required**.
-   **Success Response (200 OK):** A single user object.

#### 3. Update User

-   **Description:** Updates an existing user's information.
-   **Method:** `PUT`
-   **Path:** `/users/{id}` (e.g., `/users/1`)
-   **Authentication:** **Required**.
-   **Request Body:**
    ```json
    {
      "first_name": "John",
      "last_name": "Doe",
      "phone_number": "1234567890",
      "email": "john.doe@example.com",
      "status": "active"
    }
    ```
-   **Success Response (200 OK):** The updated user object.

#### 4. Delete User

-   **Description:** Deletes a user by their ID.
-   **Method:** `DELETE`
-   **Path:** `/users/{id}` (e.g., `/users/1`)
-   **Authentication:** **Required**.
-   **Success Response:** `204 No Content`
