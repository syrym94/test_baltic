# MyApp

A Go-based application for handling user transactions and managing balances.

## Prerequisites

- **Go**: Version 1.20+
- **Docker** & **Docker Compose**
- **golang-migrate**: For database migrations ([Install Instructions](https://github.com/golang-migrate/migrate))

---

## Setup Instructions

[//]: # (### Clone the Repository)

[//]: # ()
[//]: # (```bash)

[//]: # (git clone https://github.com/syrym94/test_baltic.git)

[//]: # (cd myapp)

[//]: # (```)

### Build and Run

#### Using Go

Build the application:
```bash
make build
```

Run the application:
```bash
make run
```

#### Using Docker

Build and run the application in Docker:
```bash
make docker-build
make docker-run
```

Stop the Docker containers:
```bash
make docker-stop
```

---

## Database Migrations

### Apply Migrations

Run all pending migrations:
```bash
make migrate-up
```

### Rollback Migrations

Rollback the last migration:
```bash
make migrate-down
```

### Create a New Migration

Generate new migration files:
```bash
make migrate-new
```
You will be prompted to enter the name of the migration. Two files (`.up.sql` and `.down.sql`) will be created in the `migrations` directory.

---

## Linting and Formatting

### Lint the Code

Run linters to check for code issues:
```bash
make lint
```

### Format the Code

Format the codebase:
```bash
make format
```

---

## Project Structure

```plaintext
.
├── cmd/                # Entry point of the application
├── internal/           # Application code (handlers, services, repositories)
├── pkg/                # Shared libraries and utilities
├── migrations/         # Database migration files
├── Dockerfile          # Docker build configuration
├── docker-compose.yml  # Docker Compose configuration
├── Makefile            # Automation commands
└── README.md           # Project documentation
```

---

## Troubleshooting

1. **Database Connection Issues**:
    - Ensure the database is running and accessible.
    - Verify `DB_URL` in the `Makefile` or environment variables.

2. **Migrations Not Found**:
    - Check that the `migrations` folder exists and contains valid `.sql` files.
    - Ensure `golang-migrate` is installed and available in your `PATH`.

3. **Docker Issues**:
    - Run `make docker-logs` to view Docker container logs.
    - Use `make docker-clean` to clean up stopped containers and dangling images.

---

## Contribution

Feel free to fork and create pull requests. For issues, please file a ticket on the GitHub repository.
