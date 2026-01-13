---
name: webapp
description: Generate web application code including HTTP handlers, repository interfaces, and implementations. Use when building REST APIs, creating CRUD endpoints, scaffolding backend services, or implementing web applications.
---

# Web Application Generator

## Overview

This skill generates web application code from domain models using Clean Architecture. It applies sensible defaults that can be overridden by project-specific CLAUDE.md.

## Workflow

1. **Analyze Relationships**
   - Use the `analyze-relations` skill to analyze domain models
   - Review the inferred relationships before proceeding

2. **Read Project Configuration**
   - Check CLAUDE.md for tech stack selection and customizations
   - Use defaults (below) for anything not specified

3. **Analyze Layer Structure** (Clean Architecture only)
   - Use the `analyze-layers` skill to derive layer structure from requirements
   - Review the proposed layer separation before proceeding

4. **Design Directory Structure**
   - Apply Stage/Axis rules to layers from `analyze-layers` output
   - Only include layers that were derived (skip absent layers)
   - See "Structure Design" section below

5. **Generate Code**
   - Output to `dist/` directory (keeps source clean)
   - Treat source domain/ as specification, generate equivalent implementation in dist/
   - Apply designed structure from step 4
   - dist/ should be a complete, runnable application

6. **Verify**
   - Ensure generated code compiles
   - Check for proper error handling

---

## Auto Selection

If not specified in CLAUDE.md, select based on project characteristics.

### Stage per Layer

Stage is selected **per layer**, applied only to layers present in `analyze-layers` output.

| Layer (if present) | Small (1-3 entities) | Medium (4-10) | Large (10+) |
|--------------------|---------------------|---------------|-------------|
| Entity | files | packages | packages |
| UseCase | inline | inline | packages |
| InterfaceAdapter | files | packages | packages |

### Axis

| Condition | Axis |
|-----------|------|
| Single team, find-by-layer useful | by-layer |
| Multiple teams, feature ownership | by-feature |
| Mixed (layer at top, feature inside) | by-layer + nested by-feature |

### Layer Definition

| Condition | Layer Definition |
|-----------|------------------|
| Default, widely understood | Clean Architecture |
| Traditional enterprise | Layered |
| Emphasis on ports/adapters | Hexagonal |
| Domain model centric | Onion |
| Simple, fewer abstractions | Three-Tier |

---

## Defaults

These defaults apply unless overridden in project's CLAUDE.md.

### Structure

Structure is specified hierarchically: LayerDefinition → Layers → SubLayers

```markdown
## Structure
LayerDefinition: <architecture style>
TopLevel: <stage>, <axis>
  <Layer>: <stage>, [axis], [expand]
    <SubLayer>: <stage>, [axis]
```

**Stage** (how far to separate):
- [inline](references/structures/stages/inline.md) - embed in adjacent layer
- [functions](references/structures/stages/functions.md) - extract functions
- [files](references/structures/stages/files.md) - split files
- [packages](references/structures/stages/packages.md) - directories with imports
- [services](references/structures/stages/services.md) - microservices

**Axis** (what to separate by):
- [by-layer](references/structures/axes/by-layer.md) - technical responsibility
- [by-feature](references/structures/axes/by-feature.md) - business capability

**expand**: Expand sublayers to top level instead of grouping under parent directory.

### Layer Definition (Reference)

- [Clean Architecture](references/layers/clean-architecture.md) (default)
- [Layered](references/layers/layered.md)
- [Hexagonal](references/layers/hexagonal.md)
- [Onion](references/layers/onion.md)
- [Three-Tier](references/layers/three-tier.md)

### Structure Design (Clean Architecture)

Design directory structure based on `analyze-layers` output.

**Example 1: Entity + InterfaceAdapter** (no UseCase)

analyze-layers output:
```markdown
### Entity
### InterfaceAdapter
  - Handler (input)
  - Repository (output)
```

→ Structure (3 entities, small):
```
dist/
├── domain/           # Entity: files
│   ├── user.go
│   └── task.go
├── handler/          # InterfaceAdapter: files, expand
│   ├── user.go
│   └── task.go
└── repository/
    └── sqlite.go
```

**Example 2: Entity + UseCase + InterfaceAdapter** (full)

analyze-layers output:
```markdown
### Entity
### UseCase
### InterfaceAdapter
  - Handler (input)
  - Repository (output)
  - Gateway (output)
```

→ Structure (8 entities, medium):
```
dist/
├── domain/           # Entity: packages
│   ├── order/
│   └── product/
├── usecase/          # UseCase: inline → embedded in handler
├── handler/          # InterfaceAdapter: packages, expand
│   ├── web/
│   └── admin/
├── repository/
│   └── postgres.go
└── gateway/
    └── payment.go
```

**Example 3: InterfaceAdapter only** (no domain logic)

analyze-layers output:
```markdown
### InterfaceAdapter
  - Handler (input)
  - Repository (output)
```

→ Structure:
```
dist/
├── handler/
│   └── bookmark.go
└── repository/
    └── sqlite.go
```

### API Design

| Relationship | Route Pattern |
|-------------|---------------|
| Top-level entity | `/{entities}`, `/{entities}/{id}` |
| belongs_to | `/{parents}/{parentID}/{children}` |
| Self-reference | `/{entities}/{id}/sub{entities}` |
| Many-to-many | `/{entities}/{id}/{related}`, `/{entities}/{id}/{related}/{relatedID}` |
| Polymorphic | Routes on each target (`/{targets}/{id}/attachments`) |

### Conventions

- Use `context.Context` for all repository methods
- Return domain errors, not database-specific errors
- Use domain methods for business logic (e.g., `entity.CanDelete(userID)`)
- Handler signature: `func(w http.ResponseWriter, r *http.Request)`
- JSON for request/response bodies
- User ID from `X-User-ID` header (for authorization checks)

### Authorization

Infer authorization rules from domain methods:
- `CanDelete(userID)` → check before delete
- `CanEdit(userID)` → check before update
- `IsOwner(userID)` → owner-only operations
- `IsMember(userID)` → member-only access

### Schema Generation

- Generate `initSchema()` function in main.go
- Include foreign key constraints based on relationships
- Add indexes for foreign key columns
- Use appropriate types per database

---

## Project CLAUDE.md

Projects only need to specify:

```markdown
## Tech Stack
- Language: Go 1.21+
- HTTP Router: chi
- Database: SQLite with sqlx
```

Optional overrides:
- Custom directory structure
- Custom route patterns
- Additional conventions
