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

**Features:**
- Order: Confirmation rules, cancellation rules, status transitions
- Product: Inventory rules, pricing
- User: Account management

### UseCase (feature-bound)
Application-specific business logic shared across interfaces.

**Features:**
- Order: ConfirmOrder, CancelOrder, GetOrderStatus
- Product: ListProducts, GetProductDetail
- User: Authenticate, GetUserProfile

### InterfaceAdapter (feature-bound)
Handles external input/output.

**Components:**
- Handler (input): Processes HTTP requests
- Repository (output): Persists data
- Gateway (output): Communicates with external services

**Features:**
- Order: OrderHandler, OrderRepository
- Product: ProductHandler, ProductRepository
- User: UserHandler, UserRepository, PaymentGateway

### Framework (cross-feature)
Technical infrastructure.

**Components:**
- DB: PostgreSQL connection, Redis connection
- HTTP: Router, middleware
```
