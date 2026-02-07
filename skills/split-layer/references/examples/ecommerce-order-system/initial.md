# Example: Order Management System — Initial

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
```

## Analysis

1. **Externals**: HTTP (Web API, Admin API), Batch scheduler, PostgreSQL, Redis, Payment API
2. **Judgments**: `Confirm()`, `CanCancel()`, "confirmed order cannot be modified", pricing calculation, order workflow orchestration
3. **Result**: Judgments exist → Application (inner) + Adapter (outer)

## Output

```markdown
## Layer Structure

### Application (feature-bound)

**Called by:** Adapter

### Adapter (feature-bound)

**Called by:** HTTP server (Web API, Admin API), Batch scheduler
```
