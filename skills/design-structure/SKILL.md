---
name: design-structure
description: Design directory structure from layer structure. Derives structure through separation decisions. Use before code generation.
---

# Directory Structure Designer

## Overview

Designs directory structure by deciding how to separate code.

## Definitions

### Layer
Horizontal separation by technical responsibility with defined dependency direction.

**Types:**
- **Feature-bound**: Has Code Units per Feature (e.g., Entity, Handler, Repository)
- **Cross-feature**: Independent of Features (e.g., Framework, Config)

### Feature
Vertical separation by domain/business concern.

Examples: User, Project, Order, Task

### Component
Subdivision within a Layer. No dependency direction between Components.

Example: Handler, Repository (within Adapter Layer)

### SubFeature
Subdivision within a Feature.

Example: Admin, Customer (within User Feature)

### Code Unit
Intersection of Layer (or Component) × Feature (or SubFeature).

### Stage
Granularity of separation:

| Stage | Description |
|-------|-------------|
| inline | All in one file |
| functions | Split into functions |
| files | Split into files |
| packages | Split into directories |
| services | Split into services |

---

## Workflow

### Step 1: Identify what exists

From analyze-layers:
- Layers and Components
- Cross-feature Layers

From domain code (the source code provided as input):
- **Features**: Identify from file structure and naming
  - Separate files (user.go, project.go) → separate Features (User, Project)
  - Single file with multiple entities → single Feature or SubFeatures
- **SubFeatures**: Identify from structure within Feature
  - Separate structs/types in one file → SubFeature candidates
  - Nested or grouped code → SubFeature candidates

**Important:** Domain code is used only to identify Features/SubFeatures. The final directory structure is derived from requirements, not from the original code structure.

### Step 2: Choose primary axis

| Condition | Axis |
|-----------|------|
| Feature count > Layer count | Feature |
| Feature count ≤ Layer count | Layer |
| Team owns by Feature | Feature |
| Team owns by Layer | Layer |

**IMPORTANT: After selecting the axis, follow ONLY the corresponding reference file:**

- If **Feature axis** selected → Follow [Feature Axis Structure](references/feature-axis.md)
- If **Layer axis** selected → Follow [Layer Axis Structure](references/layer-axis.md)

Do NOT read or reference the other file.

---

## Separation Priority

1. **Business concern** (Feature → SubFeature)
2. **Technical responsibility** (Layer → Component)

Separate by business concern first, then by technical responsibility.

---

## Output Format

Write to project's CLAUDE.md:

**IMPORTANT: Output ONLY the format specified below. The Primary axis MUST match the axis selected in Step 2.**

```markdown
## Directory Structure

Primary axis: {Feature|Layer}
Stage: {packages|files}

```
{directory or file tree}
```
```
