# Layered Architecture

Traditional layered architecture, commonly used with DDD.

## Layer Definition

| Layer | Responsibility | Depends On |
|-------|---------------|------------|
| Presentation | Display information, interpret user commands | Application |
| Application | Coordinate tasks, delegate to domain, thin layer | Domain |
| Domain | Business concepts, business rules, business state | - |
| Infrastructure | Technical capabilities (persistence, messaging) | Domain |

## Dependency Direction

```
[Presentation] → [Application] → [Domain] ← [Infrastructure]
```

All layers depend on Domain. Infrastructure implements interfaces defined in Domain.

## Rules

- Domain layer has no dependencies on other layers
- Application layer orchestrates but contains no business rules
- Infrastructure implements repository interfaces defined in Domain
