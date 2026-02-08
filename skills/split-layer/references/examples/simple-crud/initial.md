# Example: Simple Bookmark Manager (No Split)

## Input (CLAUDE.md)

```markdown
## Application

Simple bookmark manager. Stores and retrieves bookmarks.

## External Interfaces

- HTTP API: CRUD endpoints

## External Dependencies

- SQLite: Bookmark storage
```

## Analysis

1. **Externals**: HTTP API, SQLite
2. **Judgments**: None — API schema, internal model, and DB schema are essentially the same structure
3. **Result**: No judgments → single Application layer

## Output

```markdown
## Layer Structure

### Application (cross-feature)

**Called by:** External
```
