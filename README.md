# Toppira — Backend

> Backend service for **Toppira**, a reminder and todo management platform built with **Go**. Designed for production-grade reliability and future **AI** integration.

---

## 📋 Overview

Toppira is a reminder and task management API with a modular, dependency-injected architecture. It currently provides authentication (email/password and Google OAuth), user profile management, and a full domain model for reminders. The codebase is structured for scalability and ready for advanced features such as AI-powered scheduling and smart notifications.

---

## ✨ Features

| Area              | Description                                                                                                |
| ----------------- | ---------------------------------------------------------------------------------------------------------- |
| 🔐 **Auth**       | Sign-up and login with email/password; Google OAuth redirect URL for client-side flow                      |
| 👤 **User**       | Get and update profile (JWT-protected endpoints)                                                           |
| 📌 **Reminders**  | Domain model with status (`pending`, `completed`, `missed`), scheduling, priority, and JSON reminder times |
| 📚 **API docs**   | Swagger/OpenAPI generated from code annotations                                                            |
| 🗄️ **Database**   | SQLite with WAL, foreign keys, and auto-migrations (GORM)                                                  |
| 🛡️ **Resilience** | Circuit breaker on critical use cases (e.g. user creation) to avoid cascading failures                     |

---

## 🏗️ Technical

### 🛡️ Resilience & Error Handling

- **Circuit breaker (gobreaker v2)**  
  Critical flows (e.g. user creation) are wrapped in a circuit breaker. After a configurable number of consecutive failures, the circuit opens and the API returns `503 Service Unavailable` instead of overloading the database. Custom `IsSuccessful` logic treats expected cases (e.g. duplicate key) as non-failures so the circuit does not trip on business-rule violations.

- **Centralized error handling**  
  Application errors use a **typed `ErrCode`** and **`AppError`** type. Handlers and use cases return `AppError`; a **global error-handler middleware** runs after `c.Next()`, uses `errors.As` to detect `AppError`, and responds with the correct HTTP status via a **single mapping** (`ErrCode` → status). Client responses expose only a safe **`ClientError`** (e.g. `{ "error": "USER_ALREADY_EXISTS" }`) without leaking internal details.

- **Consistent error responses**  
  Auth guard and all handlers use the same error type and status mapping, so clients get predictable, stable error payloads across the API.

### 📊 Observability & Middleware

- **Structured logging (Zap)**  
  Production vs development logger configuration; all application errors and unhandled panics are logged with context.

- **Request logging & panic recovery**  
  **Ginzap** for request/response logging and **RecoveryWithZap** to recover from panics and log stack traces without crashing the server.

- **Graceful shutdown**  
  Uber **FX** lifecycle hooks start the HTTP server in a goroutine and register **`Shutdown(ctx)`** on stop, ensuring in-flight requests complete before exit.

### 🧩 Architecture & Code Quality

- **Dependency injection (Uber FX)**  
  Modules register providers and route registration via `Invoke`; the error handler and JWT guard are injected by name. Clear lifecycle for DB, logger, and HTTP server.

- **Layered design**  
  HTTP **handlers** → **use cases** → **repositories**; domain entities and shared errors live in a common package. Use cases depend on interfaces where appropriate.

- **Type-safe data access**  
  Repositories are generated with **GORM Gen**, providing type-safe queries and reducing boilerplate and runtime errors.

- **Database configuration**  
  SQLite is tuned with **WAL**, **foreign_keys**, and **busy_timeout**; **prepared statements** are enabled for safety and performance.

---

## 📦 Prerequisites

- **Go** (version specified in `go.mod`)
- **swag** CLI: `go install github.com/swaggo/swag/cmd/swag@latest`
- **golangci-lint** (recommended for CI and pre-commit)

---

## 🚀 Running the Project

- **Load env:** Variables are loaded from `.env` in `main` via `configs.LoadEnvironmentsFromEnvFile()`.
- **Start server** (from project root):
    ```bash
    ./scripts/start.sh
    ```
    Or:
    ```bash
    go run ./cmd/http
    ```
- **Swagger UI:** After the server is up: `http://localhost:<PORT>/swagger/index.html`.

---

## 📜 Developer Commands

| Command                                          | Purpose                                  |
| ------------------------------------------------ | ---------------------------------------- |
| `swag init -o ./docs -g ./cmd/http/main.go --pd` | Regenerate Swagger docs from annotations |
| `go run ./cmd/gen`                               | Regenerate repositories (GORM Gen)       |
| `golangci-lint run`                              | Run linters                              |

Run all commands from the project root. Do not edit generated files (e.g. under `docs/` or generated repository code) by hand.

---

## 🔒 API Security

Protected routes require the `Authorization: Bearer <token>` header. JWT is validated by the guard middleware; invalid or missing tokens receive a consistent error response and correct HTTP status. See Swagger for `BearerAuth` usage.

---

## 📐 Conventions

- Execute scripts and commands from the **project root**.
- Do **not** manually edit generated artifacts (Swagger docs, Gen repositories).
- Use the shared **error package** (`ErrCode`, `AppError`, `E()`) and the central **HTTP status mapping** for all API errors.
- Migrations are automatic via GORM on the entities registered in configs.

---

## 🗺️ Roadmap

- Full **Reminder** and **Todo** CRUD modules and public APIs.
- **AI** integration for smart scheduling, summarization, and categorization.
- Notifications and reminder delivery based on `reminder_times` and `scheduled_at`.

---

For questions or contributions, please reach out to the project maintainer.
