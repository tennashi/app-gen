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
