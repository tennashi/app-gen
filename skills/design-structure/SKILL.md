---
name: design-structure
description: Design directory structure from layer structure. Derives structure through iterative separation decisions. Use before code generation.
---

# Directory Structure Designer

## Overview

Designs directory structure by iteratively deciding how to separate code. Each separation decision involves two choices:
- **Axis**: Layer or Feature (direction of separation)
- **Stage**: How far to separate (granularity)

## Definitions

### Layer
Horizontal axis. Partitions by technical responsibility with defined dependency direction.

**Types:**
- **Feature-bound**: Has Code Units per Feature. Example: Entity, InterfaceAdapter
- **Cross-feature**: Independent of Features. Example: Framework, Config, Middleware

### Feature
Vertical axis. Partitions by domain/business concern. Orthogonal to Layer.
Example: Task, Project, User, Order

### Component
Technical subdivision within a Layer. No dependency direction between Components (parallel).
Example: Handler, Repository, Gateway (all within InterfaceAdapter Layer)

### Code Unit
Intersection of Layer (or Component) × Feature. The actual code to be organized.
Example: TaskEntity, TaskHandler, TaskRepository

```
Feature-bound Layers:
                          |  Task  | Project | Comment |
--------------------------|--------|---------|---------|
Entity                    |   ●    |    ●    |    ●    |
InterfaceAdapter/Handler  |   ●    |    ●    |    ●    |
InterfaceAdapter/Repository|  ●    |    ●    |    ●    |

Cross-feature Layers:
                          |   DB   |  HTTP   |  Logger |
--------------------------|--------|---------|---------|
Framework                 |   ●    |    ●    |    ●    |
```

Note:
- Handler and Repository are Components within InterfaceAdapter Layer
- Cross-feature Layers have their own subdivision (not Features)

### Axis
Direction of separation: by-layer or by-feature.

- **By Layer**: Group by technical responsibility (entity, handler, ...)
- **By Feature**: Group by domain concept (user, project, ...)

### Stage
Granularity of separation:

| Stage | Description | Example |
|-------|-------------|---------|
| inline | No separation | all in main |
| functions | Split into functions | newUser(), saveUser() |
| files | Split into files | user, handler |
| packages | Split into directories | user/, handler/ |
| services | Split into services | user-service/ |

Reference:
- [Inline](references/stages/inline.md)
- [Functions](references/stages/functions.md)
- [Files](references/stages/files.md)
- [Packages](references/stages/packages.md)
- [Services](references/stages/services.md)

## Workflow

1. **Read Layer Structure**
   - Parse CLAUDE.md for `## Layer Structure` section
   - Count Layers, Components (horizontal) and Features (vertical)
   - Identify the (Layer/Component) × Feature matrix

2. **Analyze Git History** (for existing codebases)
   - Read Git log for scale metrics (total lines, commits/month, contributors)
   - Use metrics to inform Stage selection

3. **Derive Separation Decisions**
   - For each separation point, decide Axis and Stage
   - Based on Code Unit count, Git metrics, and estimated complexity
   - **Do NOT consider current directory structure** (derive from requirements only)
   - **Do NOT pre-consider language-specific constraints** (e.g., circular dependencies)
     - If implementation fails due to constraints, adjust then

4. **Write to CLAUDE.md**
   - Write designed directory structure to project's CLAUDE.md
   - If differs from current structure, note as proposed change

---

## Separation Flow

Each separation decision = **Axis** (direction) × **Stage** (granularity)

The (Layer/Component) × Feature matrix is sliced along one axis, then optionally subdivided along the other.

**Note:** When counting for Axis selection, Components count as separate rows (like Layers).

### Step 0: Single File

Everything starts in one file (Stage: inline).

```
main   // all Code Units in one file
```

### Step 1: Initial Separation

Choose Axis based on matrix shape:

| Condition | Axis |
|-----------|------|
| Feature count > (Layer + Component) count | Feature (slice columns) |
| Feature count ≤ (Layer + Component) count | Layer (slice rows) |

Choose Stage based on Code Unit count (see Decision Criteria for details):

| Code Units | Stage |
|------------|-------|
| ≤ 30 | files (default) |
| 31+ | packages |

**Example: Feature axis + files stage** (slice by columns)
```
task        // Entity×Task, Handler×Task, Repository×Task
project     // Entity×Project, Handler×Project, Repository×Project
comment     // ...
```

**Example: Layer axis + files stage** (slice by rows)
```
entity      // Entity×Task, Entity×Project, Entity×Comment, ...
handler     // Handler×Task, Handler×Project, ...
repository  // Repository×Task, Repository×Project, ...
```

**Example: Layer axis + packages stage**
```
entity/
interface_adapter/
  handler/
  repository/
```
(Internal structure decided in Step 2)

### Step 2: Internal Separation

After initial separation, each unit may need further separation using the other Axis.

**CRITICAL: Step 2 Stage must be LOWER than Step 1 Stage.**

| Step 1 Stage | Step 2 Options |
|--------------|----------------|
| services | packages, files, functions, inline |
| packages | files, functions, inline |
| files | functions, inline |
| functions | inline |
| inline | (none) |

**Example of INVALID derivation:**
- Step 1: files stage
- Step 2: files stage ← WRONG (must be functions or inline)
- Result: packages ← WRONG (contradicts Step 1)

See [references/stages/](references/stages/) for examples of each stage.

### Step 3: Extract Cross-feature Layers

Cross-feature Layers are separated independently from Feature-bound Layers.

**Layer classification:**
```
Feature-bound:              Cross-feature:
- Entity                    - Framework
- InterfaceAdapter
```

- Feature-bound Layers share Features → can group by Feature
- Cross-feature Layers have their own subdivision → extracted independently

**Result:**
```
task/                      // Feature-bound, grouped by Feature
  entity
  handler
  repository
project/
  ...
framework/                 // Cross-feature, separated independently
  db
  http
```

### Grouping Rules

Grouping is the inverse of separation.

**Can group?**

| Condition | Can Group? |
|-----------|------------|
| Feature-bound Layers/Components | Yes (by Feature) |
| Cross-feature Layers | No (separate independently) |

**Should group?**

| Condition | Action |
|-----------|--------|
| Feature changes > Layer changes | Group by Feature |
| Team ownership by Feature | Group by Feature |
| Feature count growing | Group by Feature |
| Need to see Layer across Features | Group by Layer |

---

## Decision Criteria

### Axis Selection

| Condition | Axis |
|-----------|------|
| Feature count > (Layer + Component) count | Feature first (slice columns) |
| Feature count ≤ (Layer + Component) count | Layer first (slice rows) |
| Team owns features | Feature |
| Team owns layers | Layer |

### Stage Selection

**By Code Unit count:**

| Code Units | Stage |
|------------|-------|
| 1-3 | inline or functions |
| 4-10 | files |
| 11-30 | files (consider packages if complex) |
| 31+ | packages |

**By estimated lines per Code Unit:**

| Lines/Unit | Stage |
|------------|-------|
| < 50 | files (combine in single file per axis) |
| 50-150 | files (separate files) |
| 150+ | packages |

**Principle: Start minimal, grow as needed.**
- Default to files stage
- Only use packages when files become unwieldy

### Git-based Scale Analysis

For existing codebases (skip for new projects):

| Metric | Threshold | Implication |
|--------|-----------|-------------|
| Total lines | < 1000 | files stage sufficient |
| Total lines | 1000-5000 | files or packages |
| Total lines | 5000+ | packages likely needed |
| Commits/month | < 10 | Low churn, files sufficient |
| Commits/month | 10+ | Higher churn, consider packages |
| Contributors | 1-2 | files sufficient |
| Contributors | 3+ | packages for parallel work |

### When to Separate Further

| Signal | Action |
|--------|--------|
| File > ~300 lines | Consider next stage |
| Multiple Code Units in one file | Separate by other axis |
| Cross-feature Layer exists | Extract independently |
| Layer has multiple Components | Separate Components within Layer |

---

## Output Format

**Output must be faithful to the analysis results.**
- If Step 1 selected files stage, output must be files (not packages)
- Do NOT contradict the derived Stage in the final structure

Write to project's CLAUDE.md:

```markdown
## Directory Structure

```
{directory or file tree based on separation decisions}
```
```
