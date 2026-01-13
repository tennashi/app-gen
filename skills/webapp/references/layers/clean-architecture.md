# Clean Architecture

By Robert C. Martin (2017).

## Layer Definition

| Layer | Responsibility | Depends On |
|-------|---------------|------------|
| Entity | Enterprise business rules, domain logic | - |
| UseCase | Application-specific business rules | Entity |
| InterfaceAdapter | Data conversion (Handler, Repository, Gateway) | UseCase, Entity |
| Framework | External tools (DB, Web, UI) | InterfaceAdapter |

## Dependency Direction

```
[Framework] → [InterfaceAdapter] → [UseCase] → [Entity]
```

Outer layers depend on inner layers. Never the reverse.

## Rules

- Inner layers define interfaces, outer layers implement (Dependency Inversion)
- Data is converted to a form convenient for inner layers when crossing boundaries
- Frameworks and databases are replaceable details

## Growing Layers

Not all layers need to be separated from the start. Grow them as complexity increases.

| Complexity | Layers | When |
|------------|--------|------|
| Minimal | Entity + InterfaceAdapter | Small projects, UseCase is inline in Handler |
| Standard | Entity + UseCase + InterfaceAdapter | Medium projects, UseCase logic grows |
| Full | + Framework (explicit) | Large projects, need DB abstraction |

Start minimal. Separate when:
- UseCase logic becomes complex
- Multiple handlers share same business logic
- Need to swap database implementations

YAGNI: Don't separate until you need to.

## Layer Structure Template

Clean Architecture uses `analyze-layers` skill to derive layers dynamically based on:
- Domain logic complexity
- Number of external interfaces
- External dependencies

Example output (varies by project):

```markdown
## Layer Structure

### Entity (feature-bound)
Encapsulates business rules for domain objects.

### UseCase (feature-bound)
Application-specific business logic.

### InterfaceAdapter (feature-bound)
Handles external input/output.

**Components:**
- Handler (input): Processes HTTP requests
- Repository (output): Persists data
- Gateway (output): Communicates with external services

### Framework (cross-feature)
Technical infrastructure.

**Components:**
- DB: Database connections
- HTTP: Router, middleware
```

Note: Not all layers are always needed. See "Growing Layers" section above.
