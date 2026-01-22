# Project Architecture

This project follows the **Hexagonal Architecture** (also known as Ports and Adapters) pattern to ensure a clear separation of concerns, testability, and maintainability.

## Directory Structure Overview

The project is organized into the following key directories:

### `cmd/`
Contains the application entry points.
- `app/`: The main application entry point (likely `main.go`).
- `migrate/`: Entry point for database migration tools.

### `config/`
Handles configuration loading, typically from environment variables (`.env`).

### `internal/`
Contains the private application code, hidden from external import. This is where the core architecture lives.

#### `internal/ports/` (Interfaces)
Defines the contracts (interfaces) that separate the core logic from external dependencies.
- **`service/`**: Interfaces for the **Business Logic**. This is the core API of the application.
- **`secondary/`**: Interfaces for **Driven Adapters** (dependencies) such as Databases (`db`), Caching (`cache`), and external clients (`grpc`, `http`).
- **`primary/`**: Interfaces for **Driving Adapters**, typically related to the transport layer (e.g., gRPC contracts).

#### `internal/adapters/` (Implementations)
Contains the concrete implementations of the interfaces defined in `ports`.
- **`service/`**: Implementation of the **Business Logic** (implements `ports/service`).
- **`primary/`**: **Driving Adapters** that handle incoming requests.
    - `grpc/`: gRPC handlers/servers.
    - `http/`: HTTP handlers (REST API) using Echo.
- **`secondary/`**: **Driven Adapters** that interact with external infrastructure.
    - `repository/`: Database implementations (using GORM).
    - `grpc/`, `http/`: Clients for external services.
- **`infrastructure/`**: Data Access Objects (DAO) and ORM models (GORM structs) that map to database tables.
- **`errors/`**: Custom application error types.
- **`types/`**: Common types and constants.
- **`util/`**: Helper utility functions.

### `migrations/`
SQL migration files (`.up.sql` and `.down.sql`) for managing database schema changes using `golang-migrate`.

### `proto/`
Protocol Buffer (`.proto`) definitions and generated Go code for gRPC services.

## Architecture Layers

### 1. The Core (Business Logic)
- **Location**: `internal/adapters/service` (Logic) & `internal/ports/service` (Interface).
- **Responsibility**: Contains the business rules and domain logic. It does not depend on implementation details like the database or HTTP framework.

### 2. Primary Adapters (Driving)
- **Location**: `internal/adapters/primary` (GRPC/HTTP).
- **Responsibility**: These adapters "drive" the application. They receive requests from the outside world (e.g., REST API calls, gRPC requests) and call the Core Service.

### 3. Secondary Adapters (Driven)
- **Location**: `internal/adapters/secondary` (Repository, Cache, etc.).
- **Responsibility**: These adapters are "driven" by the application. The Core Service calls them through the ports (interfaces) to perform operations like saving data to a database or calling an external API.

## Data Flow

1.  **Request**: Enters through a **Primary Adapter** (e.g., `internal/adapters/primary/http`).
2.  **Logic**: The adapter converts the request and calls the **Core Service** (`internal/ports/service`).
3.  **Persistence**: The Core Service processes the logic and may interact with a **Secondary Adapter** (e.g., `internal/ports/secondary/repository`).
4.  **Result**: The Secondary Adapter returns data (using `internal/adapters/infrastructure` models), which flows back through the Service to the Primary Adapter and finally to the user.

## Key Technologies
- **Frameworks**: Echo (HTTP), gRPC (RPC).
- **Database**: PostgreSQL with GORM (ORM).
- **Observability**: OpenTelemetry, Jaeger, Prometheus.
- **Migrations**: Golang-Migrate.
