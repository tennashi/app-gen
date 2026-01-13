# Example: Simple CRUD (No Invariants)

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

### InterfaceAdapter (feature-bound)
Handles external input/output. Entity logic is embedded here (no separate Entity layer needed).

**Components:**
- Handler (input): Processes HTTP requests, contains data structures
- Repository (output): Persists to SQLite
```
