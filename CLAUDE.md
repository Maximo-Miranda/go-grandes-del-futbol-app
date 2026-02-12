# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Football/soccer tournament management app ("Grandes del Fútbol") built with **Go (Goravel framework)** backend and **Vue 3 + Vuetify + Inertia.js** frontend. Uses PostgreSQL via GORM.

## Development Commands

```bash
# Start PostgreSQL (required first)
docker-compose up -d

# Start backend (port 3000)
go run .

# Start frontend dev server (port 5173, run in separate terminal)
pnpm dev

# Build frontend (includes TypeScript check)
pnpm build

# Run database migrations
go run . artisan migrate

# Seed database
go run . artisan db:seed

# Run tests
go test ./tests/...
```

Tool versions managed via mise: Go 1.25, Node 22. Package manager is **pnpm**.

## Architecture

**Backend (Goravel / Laravel-style Go)**:
- `app/http/controllers/` — HTTP controllers, one per resource. Auth controllers in `auth/` subdirectory.
- `app/http/requests/` — Form request validation classes (e.g., `StorePlayerRequest`, `UpdatePlayerRequest`)
- `app/http/middleware/` — Custom middleware: `Authenticate`, `Guest`, `CSRF`, `SecurityHeaders`, `RateLimitAuth`
- `app/models/` — GORM models. Key relationships: Tournament has Teams (M2M via `TournamentTeam`), Team has Players (M2M via `TeamPlayer`), Match has Events and Lineups.
- `app/services/` — Business logic (auth service handles login/registration/password hashing)
- `app/inertia/` — Custom Inertia.js server-side adapter for Goravel
- `routes/web.go` — All web route definitions with middleware groups
- `database/migrations/` — Go-based migration files
- `config/` — Framework configuration (database, auth, HTTP, etc.)

**Frontend (Vue 3 + Inertia.js)**:
- `resources/js/pages/` — Page components, directory structure maps to routes (e.g., `Players/Index.vue`, `Players/Create.vue`)
- `resources/js/layouts/` — `AppLayout` (authenticated pages) and `AuthLayout` (login/register). Layout is auto-selected in `app.ts` based on route path.
- `resources/js/components/` — Reusable Vue components
- `resources/js/plugins/` — Vuetify configuration
- Built assets go to `public/build/` with Vite manifest

**Inertia.js bridge**: Controllers render pages via `inertia.Render(ctx, "PageName", props)` which sends data to Vue components. Form submissions use Inertia's router for SPA-like navigation without full page reloads.

**File uploads**: Player photos use POST with method spoofing (extra `POST /players/{id}` route alongside `PUT`) because multipart forms can't use PUT. Photos stored in `storage/app/public/players/`.

## Environment Setup

Copy `.env.example` to `.env`. Key settings:
- `DB_CONNECTION=postgres`, default port 5432
- `APP_PORT=3000` (backend)
- `JWT_SECRET` required for authentication
- Docker Compose provides PostgreSQL 17 on localhost:5432 (user: `postgres`, password: `password`, db: `grandes_del_futbol_development_db`)

## Local Documentation

All framework docs are available locally under `docs/`. Consult these before guessing API usage.

**Goravel v1.17** — `docs/Goravel-v1.17/`:
- `the-basics/validation.md` — Form request validation, rules, custom messages, filters
- `the-basics/request.md` — Request input, file uploads, binding
- `the-basics/response.md` — Response types, redirects, cookies
- `the-basics/session.md` — Session management, flash data
- `the-basics/controllers.md` — Controller patterns
- `the-basics/routing.md` — Route definitions, middleware groups

**Vuetify 3** — `docs/Vuetify-3/`:
- `components/` — All Vuetify component docs (data tables, forms, dialogs, etc.)
- `styles/` — Theming, colors, spacing, typography
- `features/` — Built-in features (internationalization, accessibility, etc.)
- `directives/` — Vue directives provided by Vuetify
- `api/` — Component API reference

**Inertia.js v2** — `docs/InertiaJS-v2/`:
- `the-basics/` — Pages, responses, routing, redirects
- `core-concepts/` — How Inertia works, the protocol
- `getting-started/` — Installation and setup
- `advanced/` — Advanced usage patterns
- `security/` — CSRF, authorization

## Conventions

- Validation errors use the type `map[string]map[string]string` (field -> rule -> message)
- Controllers follow resource pattern: `Index`, `Create`, `Store`, `Show`, `Edit`, `Update`, `Destroy`
- After mutations (create/update/delete), use `router.visit()` for full page refresh to ensure reactivity
- Go module path: `grandesdelfutbol`
