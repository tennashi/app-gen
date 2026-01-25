---
name: webapp
description: Generate web application code including HTTP handlers, repository interfaces, implementations, and tests. Use when building REST APIs, creating CRUD endpoints, scaffolding backend services, or implementing web applications.
---

# Web Application Generator

## Overview

This skill generates web application code from domain models using Clean Architecture. It applies sensible defaults that can be overridden by project-specific CLAUDE.md.

## Workflow

**Execute all steps in sequence without stopping for user confirmation.**

1. **Analyze Relationships**
   - Follow `analyze-relations` skill to analyze domain models
   - Do NOT write to CLAUDE.md, do NOT stop after this step

2. **Write Layer Structure**
   - Follow `analyze-layers` skill to derive layers from requirements
   - Write `## Layer Structure` with Behaviors to CLAUDE.md, do NOT stop after this step

3. **Design Directory Structure**
   - Follow `design-structure` skill to derive directory structure
   - Read `## Layer Structure` from CLAUDE.md
   - Write `## Directory Structure` to CLAUDE.md, do NOT stop after this step

4. **Generate Code**
   - Read CLAUDE.md for configuration, layer structure, and directory structure
   - Output to `dist/` directory (keeps source clean)
   - dist/ should be a complete, runnable application

5. **Generate Tests**
   - Read Behaviors from `## Layer Structure` in CLAUDE.md
   - Generate test code based on Behaviors (see Test Generation section)
   - Output to `dist/` directory alongside implementation

6. **Verify**
   - Ensure generated code compiles
   - Ensure tests pass
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

Each reference contains layer concepts and Layer Structure Template for CLAUDE.md:
- [Clean Architecture](references/layers/clean-architecture.md) (default) - use `analyze-layers` skill
- [Layered](references/layers/layered.md)
- [Hexagonal](references/layers/hexagonal.md)
- [Onion](references/layers/onion.md)
- [Three-Tier](references/layers/three-tier.md)

For non-Clean Architecture styles, copy Layer Structure Template from the reference file to CLAUDE.md.

---

## Test Generation

Tests verify that each layer fulfills its Behavior (Precondition/Postcondition/Invariant).

### Test Strategy by Layer Type

| Layer/Component | What to test | How to test |
|-----------------|--------------|-------------|
| Entity (inner) | Postcondition: correct decisions based on rules | Unit test with various inputs |
| Entity (inner) | Invariant: business rules always satisfied | Unit test that invariant holds after any operation |
| UseCase | Postcondition: goal achieved by coordinating dependencies | Unit test with mocked dependencies |
| UseCase | Invariant: consistency maintained | Unit test that checks consistency after operations |
| Handler (input) | Postcondition: correct response for valid/invalid requests | Unit test with mocked inner layer |
| Repository (output) | Postcondition: save then retrieve returns equivalent data | Integration test with real or in-memory DB |
| Gateway (output) | Postcondition: correct external call made | Unit test with mocked external service |

### Test File Structure

Place test files alongside implementation:

```
dist/
  entity/
    task.go
    task_test.go      # Entity tests
  handler/
    task_handler.go
    task_handler_test.go  # Handler tests
  repository/
    task_repository.go
    task_repository_test.go  # Repository tests
```

### Test Patterns

#### Entity Test (Invariant + Postcondition)

```go
func TestTask_CanTransitionTo(t *testing.T) {
    // Postcondition: Returns correct decision based on business rules
    task := NewTask("title", StatusTodo)

    // Valid transition
    if !task.CanTransitionTo(StatusInProgress) {
        t.Error("should allow Todo -> InProgress")
    }

    // Invalid transition
    if task.CanTransitionTo(StatusDone) {
        t.Error("should not allow Todo -> Done directly")
    }
}

func TestTask_Invariant(t *testing.T) {
    // Invariant: Entity always satisfies business rules
    task := NewTask("title", StatusTodo)
    task.Complete()

    if task.Status != StatusDone {
        t.Error("completed task should have Done status")
    }
    if task.CompletedAt == nil {
        t.Error("completed task should have CompletedAt set")
    }
}
```

#### Handler Test (Postcondition)

```go
func TestTaskHandler_Create(t *testing.T) {
    // Postcondition: Valid request → correct inner call → correct response
    mockUseCase := &MockTaskUseCase{}
    handler := NewTaskHandler(mockUseCase)

    req := httptest.NewRequest("POST", "/tasks", strings.NewReader(`{"title":"test"}`))
    rec := httptest.NewRecorder()

    handler.Create(rec, req)

    if rec.Code != http.StatusCreated {
        t.Errorf("expected 201, got %d", rec.Code)
    }
}

func TestTaskHandler_Create_InvalidRequest(t *testing.T) {
    // Postcondition: Invalid request → error response
    handler := NewTaskHandler(&MockTaskUseCase{})

    req := httptest.NewRequest("POST", "/tasks", strings.NewReader(`{invalid}`))
    rec := httptest.NewRecorder()

    handler.Create(rec, req)

    if rec.Code != http.StatusBadRequest {
        t.Errorf("expected 400, got %d", rec.Code)
    }
}
```

#### Repository Test (Postcondition)

```go
func TestTaskRepository_SaveAndFind(t *testing.T) {
    // Postcondition: Save then retrieve → equivalent data returned
    db := setupTestDB(t)
    repo := NewTaskRepository(db)

    task := &Task{Title: "test", Status: StatusTodo}
    err := repo.Save(context.Background(), task)
    if err != nil {
        t.Fatal(err)
    }

    found, err := repo.FindByID(context.Background(), task.ID)
    if err != nil {
        t.Fatal(err)
    }

    if found.Title != task.Title || found.Status != task.Status {
        t.Error("retrieved task should be equivalent to saved task")
    }
}

func TestTaskRepository_FindByID_NotFound(t *testing.T) {
    // Postcondition: Retrieve non-existent → not-found indication
    db := setupTestDB(t)
    repo := NewTaskRepository(db)

    _, err := repo.FindByID(context.Background(), "nonexistent")
    if err != ErrNotFound {
        t.Errorf("expected ErrNotFound, got %v", err)
    }
}
```

### Test Conventions

- Use table-driven tests for multiple input scenarios
- Name tests as `Test{Component}_{Method}` or `Test{Component}_{Behavior}`
- Comments should reference which Behavior (Precondition/Postcondition/Invariant) is being tested
- Use `t.Helper()` for test helper functions
- Use `t.Parallel()` where safe

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

(Includes Behaviors for each layer/component)

## Directory Structure

(Derived by design-structure skill)
```
