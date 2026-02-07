# Example: Internal TODO App — Initial

## Input (CLAUDE.md)

```markdown
## Application

Internal TODO app for a small team.

## External Interfaces

- HTTP API: RESTful endpoints for task management

## External Dependencies

- SQLite: Task and user persistence
```

## Analysis

1. **Externals**: HTTP API, SQLite
2. **Judgments**: `CanTransitionTo()`, `IsOverdue()` — task state transitions and temporal rules
3. **Result**: Judgments exist → Application (inner) + Adapter (outer)

## Output

```markdown
## Layer Structure

### Application (feature-bound)

**Called by:** Adapter

### Adapter (feature-bound)

**Called by:** HTTP server
```
