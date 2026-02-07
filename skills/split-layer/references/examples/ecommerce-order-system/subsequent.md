# Example: Order Management System — Subsequent (Split)

## Input (CLAUDE.md)

```markdown
## Application

Order management system with multiple entry points.

## External Interfaces

- Web API: Customer-facing order endpoints
- Admin API: Back-office management
- Batch: Nightly order processing

## External Dependencies

- PostgreSQL: Order persistence
- Redis: Session cache
- Payment API: External payment gateway

## Layer Structure

### Application (feature-bound)

**Called by:** Adapter

### Adapter (feature-bound)

**Called by:** HTTP server (Web API, Admin API), Batch scheduler
```

## Analysis

1. **Innermost layer**: Application
2. **One level outside**: Adapter
3. **Judgments in Application**: `Confirm()`, `CanCancel()`, "confirmed order cannot be modified", pricing calculation, order workflow orchestration (place order flow, cancel order flow)
4. **Judgments independent of Adapter?**: Yes — "confirmed order cannot be modified" and `CanCancel()` hold regardless of whether triggered via Web API, Admin API, or Batch. These domain invariants represent a different modeling target (business rules with identity and state transitions) than the workflow orchestration (operation sequencing).

**Result:** Split. Domain invariants form a new inner layer (Domain). Workflow orchestration remains (UseCase).

## Output

```markdown
## Layer Structure

### Domain (feature-bound)

**Called by:** UseCase

### UseCase (feature-bound)

**Called by:** Adapter

### Adapter (feature-bound)

**Called by:** HTTP server (Web API, Admin API), Batch scheduler
```
