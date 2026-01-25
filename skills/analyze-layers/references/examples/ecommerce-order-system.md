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

**Called by:** UseCase

**Behavior:**
- Precondition: Input satisfies type constraints
- Postcondition: Returns correct decision based on business rules
- Invariant: Entity always satisfies business rules (e.g., confirmed order cannot be modified)

### UseCase (feature-bound)

**Called by:** Handler

**Behavior:**
- Precondition: Valid use case request
- Postcondition: Goal achieved by coordinating Entity, Repository, Gateway
- Invariant: Consistency maintained across operations (e.g., order saved only if payment succeeds)

### InterfaceAdapter (feature-bound)

**Components:**

#### Handler (input)

**Called by:** HTTP server (Web API, Admin API), Batch scheduler

**Behavior:**
- Precondition: External request received
- Postcondition: Valid request → correct UseCase call → correct response; Invalid request → error response
- Invariant: Stateless

#### Repository (output)

**Called by:** UseCase

**Behavior:**
- Precondition: Entity satisfies invariants
- Postcondition: Save then retrieve → equivalent data returned; Retrieve non-existent → not-found indication
- Invariant: Persisted data integrity maintained

#### Gateway (output)

**Called by:** UseCase

**Behavior:**
- Precondition: Valid internal request
- Postcondition: Correct external API call made; External response translated to internal format
- Invariant: External errors propagated appropriately

### Framework (cross-feature)

**Called by:** All layers

**Behavior:**
- Precondition: Configuration provided
- Postcondition: Infrastructure services available (DB connection, HTTP router, etc.)
- Invariant: Infrastructure stability maintained
```
