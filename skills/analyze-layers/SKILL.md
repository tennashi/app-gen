---
name: analyze-layers
description: Analyze whether layer separation benefits the project, and derive layer structure if needed. Use when designing architecture, reviewing code structure, or before generating application code.
---

# Layer Structure Analyzer

## Overview

Analyzes whether layer separation benefits the project. If separation is warranted, derives layer structure from application requirements.

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

4. **Compare with Current Structure** (for existing codebases)
   - Read current code structure
   - Propose changes if structure differs

5. **Write to CLAUDE.md**
   - Write derived layer structure to project's CLAUDE.md
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
{Responsibility description}

**Components:** (if Layer has Components)
- {ComponentName} (input|output): {Description}
```

**If layer separation is not needed:**

```markdown
## Layer Structure

Single layer. No separation.

**Rationale:** {Why separation is unnecessary for this project}
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

---

## Examples

- [Simple Internal Tool](references/examples/simple-internal-tool.md) - Domain + Adapter
- [E-commerce Order System](references/examples/ecommerce-order-system.md) - Domain + UseCase + Adapter with multiple implementations
- [Simple CRUD](references/examples/simple-crud.md) - No layer separation (single layer)

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
