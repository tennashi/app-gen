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

### Entity (feature-bound)
Encapsulates business rules for Task and User.

### InterfaceAdapter (feature-bound)
Handles external input/output.

**Components:**
- Handler (input): Processes HTTP requests
- Repository (output): Persists to SQLite
```
