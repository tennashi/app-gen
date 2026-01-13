# Example: E-commerce Order System

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

## Domain Analysis

- `Order` has `Confirm()`, `CanCancel()` → domain logic exists
- Multiple interfaces need order confirmation logic

## Output

```markdown
## Layer Structure

### Entity (feature-bound)
Encapsulates business rules for Order, Product, User.

### UseCase (feature-bound)
Application-specific business logic shared across interfaces.

### InterfaceAdapter (feature-bound)
Handles external input/output.

**Components:**
- Handler (input): Processes HTTP requests
- Repository (output): Persists data
- Gateway (output): Communicates with external services

### Framework (cross-feature)
Technical infrastructure.

**Components:**
- DB: PostgreSQL connection, Redis connection
- HTTP: Router, middleware
```
