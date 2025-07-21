# Golang API Microservice Starter

This project is a starter kit for building microservice applications using Golang. It includes two services: an `auth-service` for handling authentication and a `user-service` for managing user data.

## Architecture

The application is designed with a microservices architecture. Each service is a standalone application with its own API and database schema. They communicate with each other via API calls.

-   **`auth-service`**: Provides user registration, login, and JWT token generation.
-   **`user-service`**: Provides CRUD operations for user management and is protected by JWT authentication.

## Tech Stack

-   **Backend**: Golang
-   **Database**: PostgreSQL
-   **Containerization**: Docker, Docker Compose

## Services

For detailed information about each service, including API documentation and setup instructions, please refer to their respective `README.md` files:

-   [auth-service/readme.md](./auth-service/readme.md)
-   [user-service/readme.md](./user-service/readme.md)

## Getting Started

To run the application, you will need Docker and Docker Compose installed.

1.  **Clone the repository:**
    ```bash
    git clone <repository-url>
    cd golang-api-microservice-starter
    ```

2.  **Run the application:**
    ```bash
    docker-compose up --build
    ```

This will start the `auth-service`, `user-service`, and a PostgreSQL database. The services will be available at:

-   `auth-service`: http://localhost:8080
-   `user-service`: http://localhost:8081
