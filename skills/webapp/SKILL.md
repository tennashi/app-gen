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

2. **Write Layer Structure**
   - Clean Architecture: Use `analyze-layers` skill (derives from requirements)
   - Other architectures: Write static layer structure (see "Static Layer Structures" below)
   - Writes `## Layer Structure` to CLAUDE.md

3. **Design Directory Structure**
   - Use the `design-structure` skill
   - Reads `## Layer Structure` from CLAUDE.md
   - Writes `## Directory Structure` to CLAUDE.md

4. **Generate Code**
   - Read CLAUDE.md for configuration, layer structure, and directory structure
   - Output to `dist/` directory (keeps source clean)
   - dist/ should be a complete, runnable application

5. **Verify**
   - Ensure generated code compiles
   - Check for proper error handling

---

## Layer Definition

| Condition | Layer Definition |
|-----------|------------------|
| Default, widely understood | Clean Architecture |
| Traditional enterprise | Layered |
| Emphasis on ports/adapters | Hexagonal |
| Domain model centric | Onion |
| Simple, fewer abstractions | Three-Tier |

Reference:
- [Clean Architecture](references/layers/clean-architecture.md) (default)
- [Layered](references/layers/layered.md)
- [Hexagonal](references/layers/hexagonal.md)
- [Onion](references/layers/onion.md)
- [Three-Tier](references/layers/three-tier.md)

---

## Static Layer Structures

For non-Clean Architecture styles, write these static structures to CLAUDE.md.

### Layered

```markdown
## Layer Structure

### Domain (feature-bound)
Business concepts and business rules.

### Application (feature-bound)
Coordinate tasks, delegate to domain.

### Presentation (feature-bound)
Display information, interpret user commands.

**Components:**
- Handler (input): HTTP request handling

### Infrastructure (feature-bound)
Technical capabilities.

**Components:**
- Repository (output): Persistence
```

### Hexagonal

```markdown
## Layer Structure

### Application (feature-bound)
Business logic and use cases.

### Port (feature-bound)
Interfaces defining how to interact with Application.

### Adapter (feature-bound)
Implementations connecting external systems to Ports.

**Components:**
- DrivingAdapter (input): HTTP Controller
- DrivenAdapter (output): Repository, External API Client
```

### Onion

```markdown
## Layer Structure

### DomainModel (feature-bound)
Entities and business rules.

### DomainServices (feature-bound)
Domain logic spanning multiple entities.

### ApplicationServices (feature-bound)
Use case orchestration.

### Infrastructure (feature-bound)
Persistence, external services, UI.

**Components:**
- Handler (input): HTTP request handling
- Repository (output): Persistence
```

### Three-Tier

```markdown
## Layer Structure

### Presentation (feature-bound)
User interface, display, input handling.

**Components:**
- Handler (input): HTTP request handling

### BusinessLogic (feature-bound)
Business rules, validation, processing.

### DataAccess (feature-bound)
Database operations, persistence.

**Components:**
- Repository (output): Database operations
```

---

## Defaults

These defaults apply unless overridden in project's CLAUDE.md.

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

Projects specify (human writes):

```markdown
## Application

{Description of the application}

## External Interfaces

- {Name}: {Description}

## External Dependencies

- {Name}: {Description}

## Tech Stack

- Language: Go 1.21+
- HTTP Router: chi
- Database: SQLite with sqlx
```

Generated by skills (can be edited by human):

```markdown
## Layer Structure

(Clean Architecture: written by analyze-layers)
(Other architectures: written by webapp from static templates)

## Directory Structure

(Written by design-structure)
```
