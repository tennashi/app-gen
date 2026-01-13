---
name: analyze-layers
description: Derive layer structure from application requirements and propose changes if the current code structure differs. Use when designing architecture, reviewing code structure, or before generating application code.
---

# Layer Structure Analyzer

## Overview

Derives layer structure from application requirements and compares it with the current code structure. Proposes changes when there are discrepancies.

## Principles

- **Dependency Rule**: Dependencies point inward only (outer layers depend on inner layers)
- **Layers are derived**: Not fixed to 4 layers; structure emerges from requirements
- **YAGNI**: Only separate when benefits outweigh costs

## Definitions

### Layer

Horizontal separation. Partitions by technical responsibility with defined dependency direction.

**Types:**
- **Feature-bound**: Has Code Units per Feature. Example: Entity, UseCase, InterfaceAdapter
- **Cross-feature**: Independent of Features. Example: Framework, Config, Middleware

### Feature

Vertical separation by domain/business concern. Orthogonal to Layer (shared across Layers).

Examples: User, Project, Order, Task

### Component

Technical subdivision within a Layer. No dependency direction between Components (parallel).

Examples: Handler, Repository, Gateway, Presenter

```
Layer: Entity
  └─ Features: User, Project, Order

Layer: UseCase
  └─ Features: CreateUser, CreateProject, ...

Layer: InterfaceAdapter
  └─ Components: Handler (input), Repository (output), Gateway (output)
  └─ Features: User, Project, Order (shared with Entity)
```

## Workflow

1. **Read Requirements**
   - Parse CLAUDE.md for application description, external interfaces, and dependencies

2. **Analyze Domain Logic**
   - Search codebase for domain logic (not limited to specific directories)
   - Domain logic = validation rules, state transitions, business constraints
   - Look for methods like `CanX()`, `IsValid()`, `Validate()`, state machine patterns

3. **Derive Layer Structure**
   - Apply derivation logic based on requirements and domain analysis

4. **Analyze Git History** (for existing codebases)
   - Read Git log for scale metrics (committer count, change frequency, total lines)
   - Apply Git-based adjustments to derivation

5. **Compare with Current Structure**
   - Read current code structure

6. **Write to CLAUDE.md**
   - Write derived layer structure to project's CLAUDE.md
   - If structure differs from current, include proposed changes

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
- Entities and domain logic from codebase
- Current code structure
- Git information (optional)

---

## Output

Write to project's CLAUDE.md:

```markdown
## Layer Structure

### {LayerName} (feature-bound|cross-feature)
{Responsibility description}

**Components:** (if Layer has Components)
- {ComponentName} (input|output): {Description}
```

Note: Features are not listed (inferred from source code).

---

## Derivation Logic

### Layer Separation

| Condition | Decision | Rationale |
|-----------|----------|-----------|
| Domain logic exists (validation, state transitions) | Separate Entity layer (feature-bound) | Testability benefit |
| No domain logic (data-only structures) | Entity layer unnecessary | No benefit to separate |
| Multiple external interfaces sharing logic | Derive UseCase layer (feature-bound) | Reusability benefit |
| Single external interface | UseCase unnecessary | No benefit to separate |
| External dependencies exist (DB, HTTP, etc.) | Derive InterfaceAdapter layer (feature-bound) | Dependency isolation |
| Shared infrastructure needed (connection pools, routers) | Derive Framework layer (cross-feature) | Reusability across Features |

### Feature Separation

| Condition | Decision | Rationale |
|-----------|----------|-----------|
| Multiple domain entities with distinct rules | Separate Features per entity | Cohesion benefit |
| Single domain entity | Single Feature | No benefit to separate |
| Features shared across Layers | List same Features in each Layer | Consistency |

### Component Separation

| Condition | Decision | Rationale |
|-----------|----------|-----------|
| Multiple input interfaces (HTTP, gRPC, CLI) | Separate Handler Components | Interface isolation |
| Multiple output dependencies (DB, cache, API) | Separate Repository/Gateway Components | Dependency isolation |
| Single input/output | No Component separation | No benefit |

### Git-based Adjustments

For existing codebases (skip for new projects):

| Condition | Adjustment |
|-----------|------------|
| Many committers (3+) | Prefer clearer Layer/Component boundaries |
| High change frequency in specific area | Prioritize separating that area |
| Large codebase (5000+ lines) | More granular Layer separation |
| Long-lived repository | Favor stability-oriented separation |

---

## Examples

- [Simple Internal Tool](references/examples/simple-internal-tool.md) - Entity + InterfaceAdapter
- [E-commerce Order System](references/examples/ecommerce-order-system.md) - Entity + UseCase + InterfaceAdapter with multiple implementations
- [Simple CRUD](references/examples/simple-crud.md) - InterfaceAdapter only (no domain logic)

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
