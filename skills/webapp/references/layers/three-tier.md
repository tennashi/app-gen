# Three-Tier Architecture

Traditional enterprise architecture.

## Layer Definition

| Layer | Responsibility | Depends On |
|-------|---------------|------------|
| Presentation | User interface, display, input handling | Business Logic |
| Business Logic | Business rules, validation, processing | Data Access |
| Data Access | Database operations, persistence | - |

## Dependency Direction

```
[Presentation] → [Business Logic] → [Data Access]
```

Top-down dependency. Each layer only knows the layer directly below.

## Rules

- Layers communicate only with adjacent layers
- Presentation never accesses Data Access directly
- Each layer can be deployed independently (in theory)
- Often maps to physical tiers (web server, app server, database server)

## Layer Structure Template

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
