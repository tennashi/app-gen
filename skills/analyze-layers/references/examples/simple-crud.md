# Example: Simple CRUD (No Layer Separation)

## Input (CLAUDE.md)

```markdown
## Application

Simple bookmark manager. No business logic.

## External Interfaces

- HTTP API: CRUD endpoints

## External Dependencies

- SQLite: Bookmark storage
```

## Domain Analysis

- `Bookmark` has only data fields, no methods → no domain logic

## Output

```markdown
## Layer Structure

Single layer. No separation.

**Rationale:** Simple CRUD with no business logic to protect. DB schema matches data model directly.

**Called by:** HTTP server, Database

**Behavior:**
- Precondition: HTTP request received
- Postcondition: Correct database operation performed; Correct HTTP response returned
- Invariant: CRUD operations maintain data consistency
```
