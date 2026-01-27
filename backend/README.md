📌 Project Architecture Context (MVC + Clean Architecture)
This project follows a pragmatic Clean Architecture combined with MVC-style delivery.
The goal is:

- Keep business logic isolated
- Allow multiple transports in the future
- Keep code generation scoped and consistent

🧱 High-level Architecture
delivery (transport / protocol)
  → handler (controller / orchestration)
    → usecase (business rules)
      → repository (database access)
        → db/models (GORM structs)
Dependency rule is strict and one-directional.

📁 Project Structure
internal/
├── infrastructure/
│   ├── database/
│   │   ├── postgres/
│   │   │   └── gorm.go
│   │   └── mongo/
│   │       └── client.go
│   │
│   ├── cache/
│   │   └── redis/
│   │       └── client.go
│   │
│   ├── messaging/
│   │   └── kafka/
│   │       ├── producer.go
│   │       └── consumer.go
│   │
│   └── external/
│       └── linkly/
│           ├── client.go
│           └── dto.go
│
├── config/                     # Application configuration (env-based)
│   ├── config.go               # Config structs
│   └── load.go                 # Load from env
│
├── db/
│   └── models/                 # Shared GORM models
│       ├── user.go
│       └── order.go
│
├── user/
│   ├── delivery/
│   │   └── http/
│   │       ├── router.go        # HTTP routing & middleware
│   │       ├── request.go       # HTTP request DTO
│   │       └── response.go      # HTTP response DTO
│   │
│   ├── handler/                # Controller / orchestration
│   │   └── user_handler.go
│   │
│   ├── usecase/                # Business logic
│   │   └── user_uc.go
│   │
│   └── repository/             # Database access (GORM)
│       └── user_repo.go
│
├── order/
│   ├── delivery/
│   │   └── http/
│   │       ├── router.go
│   │       ├── request.go
│   │       └── response.go
│   │
│   ├── handler/
│   │   └── order_handler.go
│   │
│   ├── usecase/
│   │   └── order_uc.go
│   │
│   └── repository/
│       └── order_repo.go
│
main.go                         # Composition root

🟦 Layer Responsibilities (IMPORTANT)
infrastructure/*
- Depends on external libraries and frameworks
- Implements interfaces defined by inner layers (usecase)
- Can be replaced without affecting business logic
- Contains no domain or business rules

delivery/*
- Transport-specific (HTTP)
- Routing, middleware, request binding, response mapping
- No business logic
- Can depend on handler

handler/*
- Controller / orchestration layer
- Calls usecase
- Does NOT depend on HTTP framework
- Accepts context.Context and domain models

usecase/*
- Business rules
- Validation related to domain
- No HTTP, no SQL, no framework imports

repository/*
- Database access only
- Uses GORM
- No business logic

db/models/*
- Shared GORM structs
- Can be reused across modules
- Not allowed to depend on any module

config/*
- Application configuration
- Loaded once in main.go
- Must NOT be imported by usecase

main.go
- Composition root
- Wiring dependencies only
- No business logic

🚫 Forbidden Dependencies (DO NOT GENERATE)
delivery → repository
delivery → usecase
handler → delivery
usecase → delivery
usecase → config
repository → handler
cross-module usecase calls (user → order usecase)

✅ Allowed Dependencies
delivery → handler → usecase → repository → db/models

🧠 Code Generation Rules for AI
When generating code:
1. Always ask which module (user, order, etc.)
2. Respect layer responsibility
3. New HTTP endpoints:
    - router.go (route)
    - request.go / response.go (DTO)
    - handler method
    - usecase method
    - repository method (if DB involved)
4. Never place business logic in:
    - router
    - delivery
    - repository
5. Prefer explicit, readable code over abstraction

🎯 Architectural Philosophy
This is Clean Architecture (pragmatic):
- Focus on dependency direction
- Avoid over-abstraction
- Optimize for team readability and maintainability

📝 Short Explanation (for humans)
This project separates transport, orchestration, business logic, and persistence.
It follows Clean Architecture principles while remaining practical for real-world development.