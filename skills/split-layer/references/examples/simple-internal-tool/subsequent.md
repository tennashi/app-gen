# Example: Internal TODO App — Subsequent (No Split)

## Input (CLAUDE.md)

```markdown
## Application

Internal TODO app for a small team.

## External Interfaces

- HTTP API: RESTful endpoints for task management

## External Dependencies

- SQLite: Task and user persistence

## Layer Structure

### Application (feature-bound)

**Called by:** Adapter

### Adapter (feature-bound)

**Called by:** External
```

## Analysis

1. **Innermost layer**: Application
2. **One level outside**: Adapter
3. **Judgments in Application**: `CanTransitionTo()`, `IsOverdue()` — task state transitions and temporal rules
4. **Judgments independent of Adapter?**: No — the rules and orchestration share a single modeling target ("task management"). No distinct inner modeling target emerges.

**Result:** No split.

## Output

No change to Layer Structure.
