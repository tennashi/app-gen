---
name: analyze-layers
description: Analyze whether layer separation benefits the project, and derive layer structure with behaviors. Use when designing architecture, reviewing code structure, or before generating application code.
---

# Layer Structure Analyzer

## Overview

Analyzes whether layer separation benefits the project. If separation is warranted, derives layer structure with behaviors (preconditions, postconditions, invariants) from application requirements.

## What is a Layer?

- **Horizontal separation** by technical responsibility
- **Dependency direction is defined** (outer → inner)
- Orthogonal to Feature (vertical separation by business concern)

Layer separation is **not always necessary**. It has costs (complexity). Only separate when benefits outweigh costs.

## When Layer Separation is Beneficial

The core question: **Do you want to keep domain logic independent from external implementations?**

| Situation | Why separation helps |
|-----------|---------------------|
| Multiple persistence targets (RDB + MongoDB) | Domain doesn't need to know persistence details |
| Multiple input interfaces (HTTP + gRPC + CLI) | Domain doesn't need to know input format |
| Multiple external integrations (Payment A + Payment B) | Domain doesn't need to know API details |

→ When **multiple implementations exist**, and you want to keep domain independent from them.

## When Layer Separation is Unnecessary

| Situation | Why |
|-----------|-----|
| Simple CRUD | No domain logic to protect |
| Single persistence target, unlikely to change | Little benefit from separation |
| DB schema ≒ Domain model | No divergence to manage |
| Prototype / PoC | Speed over structure |

## Definitions

### Layer

Horizontal separation with defined dependency direction.

**Types:**
- **Feature-bound**: Has Code Units per Feature (e.g., domain layer, adapter layer)
- **Cross-feature**: Independent of Features (e.g., framework, config, middleware)

### Feature

Vertical separation by domain/business concern. Orthogonal to Layer.

Examples: User, Project, Order, Task

### Component

Technical subdivision within a Layer. No dependency direction between Components (parallel).

Examples: Handler, Repository, Gateway, Presenter

### Behavior

A layer's behavior defines what it does from the outside. Described using:

- **Precondition**: What must be true before the operation
- **Postcondition**: What will be true after the operation
- **Invariant**: What is always maintained across all operations

Behaviors are **domain-independent** at the layer level. They define the pattern of behavior, not specific domain rules.

## Workflow

1. **Read Requirements**
   - Parse CLAUDE.md for application description, external interfaces, and dependencies

2. **Analyze Whether Separation is Beneficial**
   - Check: Are there multiple implementations for persistence, input, or external integrations?
   - Check: Is there domain logic worth protecting from external dependencies?
   - If no clear benefit, recommend no layer separation

3. **Derive Layer Structure** (if separation is beneficial)
   - Identify boundaries where dependency direction matters
   - Name layers by their responsibility (naming is flexible)
   - Define behavior for each layer/component

4. **Compare with Current Structure** (for existing codebases)
   - Read current code structure
   - Propose changes if structure differs

5. **Write to CLAUDE.md**
   - Write derived layer structure with behaviors to project's CLAUDE.md
   - If no separation needed, document that decision with rationale

---

## Input

Human writes in CLAUDE.md:

```markdown
## Application

{Description of the application}

## External Interfaces

- {Name}: {Description}

## External Dependencies

- {Name}: {Description}
```

Automatically gathered:
- Domain logic from codebase
- Current code structure

---

## Output

Write to project's CLAUDE.md:

**If layer separation is beneficial:**

```markdown
## Layer Structure

### {LayerName} (feature-bound|cross-feature)

**Called by:** {who calls this layer}

**Behavior:**
- Precondition: {precondition}
- Postcondition: {postcondition}
- Invariant: {invariant}

**Components:** (if Layer has Components)

#### {ComponentName} (input|output)

**Called by:** {who calls this component}

**Behavior:**
- Precondition: {precondition}
- Postcondition: {postcondition}
- Invariant: {invariant}
```

**If layer separation is not needed:**

```markdown
## Layer Structure

Single layer. No separation.

**Rationale:** {Why separation is unnecessary for this project}

**Behavior:**
- Precondition: {precondition}
- Postcondition: {postcondition}
- Invariant: {invariant}
```

---

## Derivation Logic

### Should We Separate Layers?

| Question | If Yes | If No |
|----------|--------|-------|
| Multiple persistence targets? | Separate domain from persistence | No separation needed for this |
| Multiple input interfaces? | Separate domain from input handling | No separation needed for this |
| Multiple external integrations? | Separate domain from external APIs | No separation needed for this |
| Domain logic worth protecting? | Separate domain layer | No domain layer needed |

If all answers are "No", layer separation is likely unnecessary.

### What Layers to Create?

Layers are defined by **responsibility**, not fixed names.

| Responsibility | Common names (examples) |
|---------------|------------------------|
| Holds domain rules and logic | Domain, Entity, Model, Core |
| Coordinates use cases across interfaces | UseCase, Application, Service |
| Handles external input/output | Adapter, Infrastructure, Gateway |
| Provides shared infrastructure | Framework, Platform |

Choose names that fit your project. Existing patterns (MVC, Layered, Clean Architecture, Hexagonal) are references, not requirements.

### Component Separation (within a Layer)

| Condition | Decision |
|-----------|----------|
| Multiple input interfaces (HTTP, gRPC, CLI) | Separate input Components |
| Multiple output dependencies (DB, cache, API) | Separate output Components |
| Single input/output | No Component separation needed |

### Behavior Patterns by Layer Type

| Layer/Component type | Behavior pattern |
|---------------------|------------------|
| Inner layer (domain) | Precondition: valid input / Postcondition: correct decision based on rules / Invariant: business rules always satisfied |
| Orchestration layer (usecase) | Precondition: valid request / Postcondition: goal achieved by coordinating dependencies / Invariant: consistency maintained |
| Input component | Precondition: external request received / Postcondition: correct internal call made, correct response returned / Invariant: stateless |
| Output component | Precondition: valid internal request / Postcondition: external operation completed, result translated / Invariant: data integrity maintained |

---

## Examples

### Example 1: Simple Internal Tool

**Input (CLAUDE.md):**

```markdown
## Application

Internal TODO app for a small team.

## External Interfaces

- HTTP API: RESTful endpoints for task management

## External Dependencies

- SQLite: Task and user persistence
```

**Domain Analysis:**

- `Task` has `CanTransitionTo()`, `IsOverdue()` → domain logic exists

**Output:**

```markdown
## Layer Structure

### Entity (feature-bound)

**Called by:** InterfaceAdapter

**Behavior:**
- Precondition: Input satisfies type constraints
- Postcondition: Returns correct decision based on business rules
- Invariant: Entity always satisfies business rules

### InterfaceAdapter (feature-bound)

**Components:**

#### Handler (input)

**Called by:** HTTP server

**Behavior:**
- Precondition: HTTP request received
- Postcondition: Valid request → correct Entity call → correct HTTP response; Invalid request → error response
- Invariant: Stateless

#### Repository (output)

**Called by:** Handler

**Behavior:**
- Precondition: Entity satisfies invariants
- Postcondition: Save then retrieve → equivalent data returned; Retrieve non-existent → not-found indication
- Invariant: Persisted data integrity maintained
```

### Example 2: E-commerce Order System

**Input (CLAUDE.md):**

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

**Domain Analysis:**

- `Order` has `Confirm()`, `CanCancel()` → domain logic exists
- Multiple interfaces need order confirmation logic

**Output:**

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

### Example 3: Simple CRUD (No Layer Separation)

**Input (CLAUDE.md):**

```markdown
## Application

Simple bookmark manager. No business logic.

## External Interfaces

- HTTP API: CRUD endpoints

## External Dependencies

- SQLite: Bookmark storage
```

**Domain Analysis:**

- `Bookmark` has only data fields, no methods → no domain logic

**Output:**

```markdown
## Layer Structure

Single layer. No separation.

**Rationale:** Simple CRUD with no business logic to protect. DB schema matches data model directly.

**Called by:** HTTP server, Database

**Behavior:**
- Precondition: HTTP request received
- Postcondition: Correct database operation performed; Correct HTTP response returned
- Invariant: CRUD operations maintain data consistency
```

---

## Change Proposal Format

When current structure differs from derived structure:

```markdown
## Proposed Changes

### {Change description}

**Current:** {What exists now}

**Proposed:** {What should be}

**Rationale:** {Why this change benefits the project}

**Steps:**
1. {Concrete step}
2. {Concrete step}
```

---

## Notes

- Behaviors define **what** each layer does from the outside, which becomes **what tests verify**
- Behaviors are domain-independent patterns; specific conditions come from domain knowledge
- Combine layer behaviors with domain details to write actual test cases
