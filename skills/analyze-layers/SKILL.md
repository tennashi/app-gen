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

## Workflow

1. **Read Requirements**
   - Parse CLAUDE.md for application description, external interfaces, and dependencies

2. **Analyze Domain Logic**
   - Search codebase for domain logic (not limited to specific directories)
   - Domain logic = validation rules, state transitions, business constraints
   - Look for methods like `CanX()`, `IsValid()`, `Validate()`, state machine patterns

3. **Derive Layer Structure**
   - Apply derivation logic based on requirements and domain analysis

4. **Compare with Current Structure**
   - Read current code structure
   - Optionally read Git log for context (committer count, change frequency)

5. **Write to CLAUDE.md**
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

### {LayerName}
{Responsibility description}

**Features:**
- {FeatureName} (input|output): {Description}
  - Implementations: {impl1}, {impl2}
```

---

## Derivation Logic

### Layer Separation

| Condition | Decision | Rationale |
|-----------|----------|-----------|
| Domain logic exists (validation, state transitions) | Separate Entity layer | Testability benefit |
| No domain logic (data-only structures) | Entity layer unnecessary | No benefit to separate |
| Multiple external interfaces sharing logic | Derive UseCase layer | Reusability benefit |
| Single external interface | UseCase unnecessary | No benefit to separate |

### Feature Separation

| Condition | Decision | Rationale |
|-----------|----------|-----------|
| Multiple external dependencies of same type | List as implementations | Swappability benefit |
| Single dependency per type | No implementations list | No benefit |

### Git-based Adjustments (Optional)

| Condition | Adjustment |
|-----------|------------|
| Many committers | Prefer clearer boundaries for parallel development |
| High change frequency in specific area | Prioritize separating that area |
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
