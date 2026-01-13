# Clean Architecture

By Robert C. Martin (2017).

## Layer Definition

| Layer | Responsibility | Depends On |
|-------|---------------|------------|
| Entity | Enterprise business rules, domain logic | - |
| UseCase | Application-specific business rules | Entity |
| Interface Adapter | Data conversion (Controller, Presenter, Gateway) | UseCase, Entity |
| Framework & Driver | External tools (DB, Web, UI) | Interface Adapter |

## Dependency Direction

```
[Framework & Driver] → [Interface Adapter] → [UseCase] → [Entity]
```

Outer layers depend on inner layers. Never the reverse.

## Rules

- Inner layers define interfaces, outer layers implement (Dependency Inversion)
- Data is converted to a form convenient for inner layers when crossing boundaries
- Frameworks and databases are replaceable details

## Growing Layers

Not all layers need to be separated from the start. Grow them as complexity increases.

| Stage | Layers | When |
|-------|--------|------|
| Minimal | Entity + InterfaceAdapter | Small projects, UseCase is inline in Handler |
| Standard | Entity + UseCase + InterfaceAdapter | Medium projects, UseCase logic grows |
| Full | + Framework (explicit) | Large projects, need DB abstraction |

Start minimal. Separate when:
- UseCase logic becomes complex
- Multiple handlers share same business logic
- Need to swap database implementations

YAGNI: Don't separate until you need to.
