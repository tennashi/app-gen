# Hexagonal Architecture (Ports and Adapters)

By Alistair Cockburn (2005).

## Layer Definition

| Layer | Responsibility | Depends On |
|-------|---------------|------------|
| Application | Business logic, use cases | - |
| Port | Interface defining how to interact with Application | Application |
| Adapter | Implementation connecting external systems to Ports | Port |

## Dependency Direction

```
[Adapter] → [Port] → [Application]
```

Adapters depend on Ports. Application defines Ports but doesn't know Adapters.

## Port Types

| Type | Direction | Example |
|------|-----------|---------|
| Driving (Primary) | Outside → Application | HTTP Controller, CLI |
| Driven (Secondary) | Application → Outside | Repository, External API Client |

## Rules

- Application is isolated from external concerns
- Ports are interfaces owned by Application
- Adapters are interchangeable (test doubles, different implementations)
- No distinction between "front" and "back" - all external systems are equal

## Layer Structure Template

```markdown
## Layer Structure

### Application (feature-bound)
Business logic and use cases.

### Port (feature-bound)
Interfaces defining how to interact with Application.

### Adapter (feature-bound)
Implementations connecting external systems to Ports.

**Components:**
- DrivingAdapter (input): HTTP Controller
- DrivenAdapter (output): Repository, External API Client
```
