---
name: split-layer
description: Determine whether the innermost layer should split. Use when designing architecture, reviewing code structure, or before generating application code.
---

# Split Layer

## Overview

Analyzes whether the current innermost layer should split into two layers.

## Concepts

### What is a Layer?

A layer is a boundary where **modeling target** changes.

Every application has an inside and an outside. The outside is I/O — files, network, stdin/stdout, databases. Between the outside and the inside, OS, runtime, and libraries already provide a layer of abstraction. The application developer's code begins on top of these existing layers.

Within the application's own code, a new layer emerges when a **new modeling target** appears — when the code begins to model something different from what it was modeling before.

### Layer vs. Module

Extracting code into a function or module does not create a layer. A new layer only emerges when the extracted code carries **a different modeling target**. If it models the same thing, it is a module within the same layer.

### Layer vs. Feature

- **Layer**: Horizontal separation. Boundaries where modeling target changes.
- **Feature**: Vertical separation by domain/business concern.

These are orthogonal.

### How Layers Emerge

Layers are not chosen from a menu of architectural patterns. They emerge through a single recurring question: **does the innermost layer contain judgments that do not depend on the layer one level outside?**

If yes, those judgments split off as a new inner layer. The same principle applies recursively — in the TCP/IP protocol stack, each layer's judgments hold independently of adjacent layers.

### Boundary Between Layers

At each layer boundary, what changes is the modeling target. Common patterns:

**Between I/O and application logic**: The **owner of representation** switches. On the I/O side, the shape of data is dictated by external contracts (API schema, wire protocol, DB schema). On the application side, the shape is determined by the application developer.

**Between application logic and domain logic**: The **structure of models** changes. Application models represent operation inputs and outputs — essentially data. Domain models represent business concepts with identity, invariants, state transitions, and behavior.

## Definitions

### Layer

A boundary where modeling target changes, with defined dependency direction (outer → inner).

**Types:**
- **Feature-bound**: Has code units per Feature (e.g., domain layer, adapter layer)
- **Cross-feature**: Independent of Features (e.g., framework, config, middleware)

### Feature

Vertical separation by domain/business concern. Orthogonal to Layer.

Examples: User, Project, Order, Task

## Workflow

First, check whether `## Layer Structure` exists in the project's CLAUDE.md.

- If it does **not** exist → read and follow **only** `references/initial.md`
- If it **does** exist → read and follow **only** `references/subsequent.md`

Do NOT read both files. Read only the one that applies.

---

## Input

From project's CLAUDE.md:

- Application description
- External Interfaces
- External Dependencies
- Layer Structure (if already exists)

From codebase (if exists):

- Domain logic in code
- Current code structure

---

## Output

**IMPORTANT: Output ONLY the format specified below. Do NOT add any extra sections or fields. Keep the output minimal and strictly follow the format.**

### Layer

```markdown
### {LayerName} (feature-bound|cross-feature)

**Called by:** {caller layer, or External}
```

### Feature (if multiple domain concerns exist)

```markdown
**Features:** {Feature1}, {Feature2}, ...
```

### Full structure

Combine the above. List layers from inner to outer. Wrap in `## Layer Structure`.

```markdown
## Layer Structure

{layers}
```

---

## Notes

- OS, runtime, and libraries already provide layers below your application code — layer analysis focuses on layers **within your own code**
- The number of layers is not fixed by any pattern — it is derived from the modeling targets your application actually needs
