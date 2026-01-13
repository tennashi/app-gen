# Onion Architecture

By Jeffrey Palermo (2008).

## Layer Definition

| Layer | Responsibility | Depends On |
|-------|---------------|------------|
| Domain Model | Entities, business rules | - |
| Domain Services | Domain logic spanning multiple entities | Domain Model |
| Application Services | Use case orchestration, transaction boundary | Domain Services, Domain Model |
| Infrastructure | Persistence, external services, UI | Application Services |

## Dependency Direction

```
[Infrastructure] → [Application Services] → [Domain Services] → [Domain Model]
```

Outer layers depend on inner layers. Domain Model is at the center.

## Rules

- Domain Model has no outward dependencies
- All coupling is toward the center
- Infrastructure concerns are pushed to the outer layer
- Interfaces are defined in inner layers, implemented in outer layers

## Layer Structure Template

```markdown
## Layer Structure

### DomainModel (feature-bound)
Entities and business rules.

### DomainServices (feature-bound)
Domain logic spanning multiple entities.

### ApplicationServices (feature-bound)
Use case orchestration.

### Infrastructure (feature-bound)
Persistence, external services, UI.

**Components:**
- Handler (input): HTTP request handling
- Repository (output): Persistence
```
