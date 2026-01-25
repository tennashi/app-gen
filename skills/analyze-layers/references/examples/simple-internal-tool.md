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

**Called by:** InterfaceAdapter

**Behavior:**
- Precondition: Input satisfies type constraints
- Postcondition: Returns correct decision based on business rules
- Invariant: Entity always satisfies business rules

### InterfaceAdapter (feature-bound)

**Components:**

#### Handler (input)

**Called by:** HTTP server

**Behavior:**
- Precondition: HTTP request received
- Postcondition: Valid request → correct Entity call → correct HTTP response; Invalid request → error response
- Invariant: Stateless

#### Repository (output)

**Called by:** Handler

**Behavior:**
- Precondition: Entity satisfies invariants
- Postcondition: Save then retrieve → equivalent data returned; Retrieve non-existent → not-found indication
- Invariant: Persisted data integrity maintained
```
