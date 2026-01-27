📌 Project Architecture Context (MVC + Clean Architecture)
This project follows a pragmatic Clean Architecture combined with MVC-style delivery, mirroring the backend architecture.
The goal is:

- Keep business logic isolated from UI
- Allow multiple UI frameworks in the future (React Native, Vue, etc.)
- Keep code generation scoped and consistent
- Maintain clear separation of concerns

🧱 High-level Architecture
presentation (UI components / pages)
  → application/hooks (orchestration / React hooks)
    → application/usecase (business rules)
      → application/repository (API calls)
        → domain/models (TypeScript types/interfaces)
Dependency rule is strict and one-directional.

📁 Project Structure
src/
├── infrastructure/
│   ├── api/
│   │   ├── client.ts              # Base API client (axios/fetch)
│   │   └── interceptors.ts         # Request/response interceptors
│   │
│   ├── storage/
│   │   ├── localStorage.ts         # LocalStorage wrapper
│   │   └── sessionStorage.ts       # SessionStorage wrapper
│   │
│   ├── cache/
│   │   └── memoryCache.ts          # In-memory cache
│   │
│   └── external/
│       └── analytics.ts             # External services (analytics, etc.)
│
├── config/                         # Application configuration
│   ├── config.ts                   # Config constants
│   └── env.ts                      # Environment variables
│
├── domain/
│   └── models/                     # Shared domain models/types
│       ├── user.ts                 # User type definitions
│       └── order.ts                # Order type definitions
│
├── user/
│   ├── presentation/               # UI layer
│   │   ├── pages/
│   │   │   ├── UserListPage.tsx
│   │   │   └── UserDetailPage.tsx
│   │   └── components/
│   │       ├── UserCard.tsx
│   │       └── UserForm.tsx
│   │
│   ├── application/               # Business logic layer
│   │   ├── hooks/
│   │   │   └── useUser.ts         # React hooks (orchestration)
│   │   ├── usecase/
│   │   │   └── userUsecase.ts     # Business rules
│   │   └── repository/
│   │       └── userRepository.ts  # API calls
│   │
│   └── domain/
│       └── types.ts                # User-specific types
│
├── order/
│   ├── presentation/
│   │   ├── pages/
│   │   │   ├── OrderListPage.tsx
│   │   │   └── OrderDetailPage.tsx
│   │   └── components/
│   │       ├── OrderCard.tsx
│   │       └── OrderForm.tsx
│   │
│   ├── application/
│   │   ├── hooks/
│   │   │   └── useOrder.ts
│   │   ├── usecase/
│   │   │   └── orderUsecase.ts
│   │   └── repository/
│   │       └── orderRepository.ts
│   │
│   └── domain/
│       └── types.ts
│
├── shared/                         # Shared utilities
│   ├── components/                 # Reusable UI components
│   │   ├── Button.tsx
│   │   ├── Input.tsx
│   │   └── Loading.tsx
│   ├── hooks/                      # Shared React hooks
│   │   ├── useAuth.ts
│   │   └── useToast.ts
│   └── utils/                      # Utility functions
│       ├── validation.ts
│       └── format.ts
│
├── app/                            # Application setup
│   ├── App.tsx                     # Root App component
│   ├── router.tsx                   # Routing configuration
│   └── providers.tsx                # Context providers
│
└── main.tsx                        # Entry point

🟦 Layer Responsibilities (IMPORTANT)

infrastructure/*
- Depends on external libraries (axios, localStorage API, etc.)
- Implements interfaces defined by inner layers
- Can be replaced without affecting business logic
- Contains no domain or business rules
- Examples: API client setup, storage wrappers, external service integrations

presentation/*
- UI components and pages only
- React-specific code (JSX, hooks for UI state)
- No business logic
- Can depend on application/hooks
- Handles user interactions and displays data
- Examples: Pages, form components, card components

application/hooks/*
- React hooks that orchestrate usecases
- Connects presentation to usecase layer
- Manages React state (useState, useEffect)
- Calls usecase methods
- Does NOT contain business logic
- Examples: useUser, useOrder

application/usecase/*
- Business rules and validation
- Pure functions (no React dependencies)
- No API calls, no UI code
- Can be tested independently
- Examples: User validation, order calculation logic

application/repository/*
- API calls only
- Data fetching and mutation
- Uses infrastructure/api client
- No business logic
- Returns domain models
- Examples: fetchUsers, createOrder

domain/models/*
- Shared TypeScript types/interfaces
- Can be reused across modules
- Not allowed to depend on any module
- Pure data structures
- Examples: User interface, Order interface

config/*
- Application configuration
- Environment variables
- Must NOT be imported by usecase
- Loaded once at app startup

shared/*
- Reusable utilities and components
- Cross-cutting concerns
- Can be used by any module
- Examples: Button component, validation utils

app/*
- Application composition root
- Routing setup
- Provider setup
- No business logic

main.tsx
- Application entry point
- Initializes React app
- Wires dependencies

🚫 Forbidden Dependencies (DO NOT GENERATE)
presentation → repository
presentation → usecase
presentation → infrastructure
hooks → presentation
usecase → presentation
usecase → infrastructure (except through repository)
usecase → config
repository → hooks
repository → usecase
cross-module usecase calls (user → order usecase)
domain/models → any module

✅ Allowed Dependencies
presentation → hooks → usecase → repository → infrastructure
presentation → shared
hooks → usecase → repository
usecase → domain/models
repository → domain/models
repository → infrastructure/api
all → domain/models (shared types)

🧠 Code Generation Rules for AI
When generating code:
1. Always ask which module (user, order, etc.)
2. Respect layer responsibility strictly
3. New feature flow:
   - presentation: Create page/component
   - hooks: Create or update hook (useUser, useOrder)
   - usecase: Add business logic method
   - repository: Add API call method
   - domain: Add/update types if needed
4. Never place business logic in:
   - presentation components
   - hooks (only orchestration)
   - repository (only API calls)
5. Prefer explicit, readable code over abstraction
6. Use TypeScript strictly - define types for everything

📋 Example: Creating a New Feature

Let's say we want to add "Create User" functionality:

1. **domain/models/user.ts** - Define User type (if not exists)
```typescript
export interface User {
  id: string;
  name: string;
  email: string;
}
```

2. **user/application/repository/userRepository.ts** - Add API call
```typescript
export const createUser = async (data: CreateUserRequest): Promise<User> => {
  // API call using infrastructure/api client
}
```

3. **user/application/usecase/userUsecase.ts** - Add business logic
```typescript
export const createUserUsecase = (data: CreateUserRequest): ValidationResult => {
  // Validate email format
  // Validate name length
  // Return validation result
}
```

4. **user/application/hooks/useUser.ts** - Add React hook
```typescript
export const useCreateUser = () => {
  const [loading, setLoading] = useState(false);
  // Call usecase for validation
  // Call repository for API
  // Return { createUser, loading, error }
}
```

5. **user/presentation/components/UserForm.tsx** - Add UI
```typescript
export const UserForm = () => {
  const { createUser, loading } = useCreateUser();
  // Form UI
  // Call createUser on submit
}
```

🎯 Architectural Philosophy
This is Clean Architecture (pragmatic) for React:
- Focus on dependency direction
- Avoid over-abstraction
- Optimize for team readability and maintainability
- Keep React-specific code (hooks, JSX) in presentation and hooks layers
- Keep business logic pure and testable

📝 Short Explanation (for humans)
This project separates UI, orchestration, business logic, and data access.
It follows Clean Architecture principles while remaining practical for React development.
Business logic is isolated and can be tested without React or API dependencies.

🔧 Technology Stack Recommendations
- **React 18+** with TypeScript
- **React Router** for routing
- **Axios** or **Fetch API** for HTTP client
- **React Query** or **SWR** for data fetching (optional, can wrap repository layer)
- **Zustand** or **Context API** for global state (if needed)
- **React Hook Form** for form management (in presentation layer)
- **Tailwind CSS** or **CSS Modules** for styling

⚠️ Important Notes
- Hooks layer (application/hooks) is React-specific and acts as the bridge between UI and business logic
- Usecase layer should be framework-agnostic (no React imports)
- Repository layer handles all external data fetching
- Presentation layer should be "dumb" components that receive props and call hooks
- Always define TypeScript types in domain/models or module/domain
