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

### Entity
Encapsulates business rules for Order, Product, User.

### UseCase
Application-specific business logic shared across interfaces.
- Order confirmation, cancellation, status transitions

### InterfaceAdapter
Handles external input/output.

**Features:**
- Handler (input): Processes HTTP requests
  - Implementations: WebHandler, AdminHandler, BatchJob
- Repository (output): Persists data
  - Implementations: PostgresRepository, RedisCache
- Gateway (output): Communicates with external services
  - Implementations: PaymentGateway
```
