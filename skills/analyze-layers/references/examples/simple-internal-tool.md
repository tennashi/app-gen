# Example: Simple Internal Tool

## Input (CLAUDE.md)

```markdown
## Application

Internal TODO app for a small team.

## External Interfaces

- HTTP API: RESTful endpoints for task management

## External Dependencies

- SQLite: Task and user persistence
```

## Domain Analysis

- `Task` has `CanTransitionTo()`, `IsOverdue()` → domain logic exists

## Output

```markdown
## Layer Structure

### Entity
Encapsulates business rules for Task and User.

### InterfaceAdapter
Handles external input/output.

**Features:**
- Handler (input): Processes HTTP requests
- Repository (output): Persists to SQLite
```
