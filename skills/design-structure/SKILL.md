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

Reference:
- [By Layer](references/axes/by-layer.md)
- [By Feature](references/axes/by-feature.md)

### Stage
Granularity of separation:

| Stage | Description | Example |
|-------|-------------|---------|
| inline | No separation | all in `main()` |
| functions | Split into functions | `newUser()`, `saveUser()` |
| files | Split into files | `user.go`, `handler.go` |
| packages | Split into directories | `user/`, `handler/` |
| services | Split into services | `user-service/` |

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

2. **Derive Separation Decisions**
   - For each separation point, decide Axis and Stage
   - Based on counts and implementation volume

3. **Write to CLAUDE.md**
   - Write designed directory structure to project's CLAUDE.md

---

## Separation Flow

Each separation decision = **Axis** (direction) × **Stage** (granularity)

The (Layer/Component) × Feature matrix is sliced along one axis, then optionally subdivided along the other.

**Note:** When counting for Axis selection, Components count as separate rows (like Layers).

### Step 0: Single File

Everything starts in one file (Stage: inline).

```
main.go   // all Code Units in one file
```

### Step 1: Initial Separation

Choose Axis based on matrix shape:

| Condition | Axis |
|-----------|------|
| Feature count > (Layer + Component) count | Feature (slice columns) |
| Feature count ≤ (Layer + Component) count | Layer (slice rows) |

Choose Stage based on total volume:

| Volume | Stage |
|--------|-------|
| Small (fits in files) | files |
| Medium/Large | packages |

**Example: Feature axis + files stage** (slice by columns)
```
task.go      // Entity×Task, Handler×Task, Repository×Task
project.go   // Entity×Project, Handler×Project, Repository×Project
comment.go   // ...
```

**Example: Layer axis + files stage** (slice by rows)
```
entity.go     // Entity×Task, Entity×Project, Entity×Comment, ...
handler.go    // Handler×Task, Handler×Project, ...
repository.go // Repository×Task, Repository×Project, ...
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

**Important:** Each unit independently chooses its Stage based on its own volume.

| Volume | Stage |
|--------|-------|
| Small | inline (no separation) |
| Medium | files |
| Large | packages |

**Feature first → then Layer/Component (files stage):**
```
task/
  entity.go      // Entity×Task
  handler.go     // InterfaceAdapter/Handler×Task
  repository.go  // InterfaceAdapter/Repository×Task
project/
  entity.go
  handler.go
  repository.go
```

**Layer first → then Feature (files stage):**
```
entity/
  task.go
  project.go
  comment.go
interface_adapter/
  handler/
    task.go        // Handler×Task
    project.go     // Handler×Project
  repository/
    task.go
    project.go
```

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
  entity.go
  handler.go
  repository.go
project/
  ...
framework/                 // Cross-feature, separated independently
  db.go
  http.go
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

| Condition | Stage |
|-----------|-------|
| Few items, small code | inline |
| Multiple items, still small code | functions |
| Multiple items, moderate code | files |
| Many items or large code | packages |
| Independent deployment needed | services |

### When to Separate Further

| Signal | Action |
|--------|--------|
| File > ~300 lines | Consider next stage |
| Multiple Code Units in one file | Separate by other axis |
| Cross-feature Layer exists | Extract independently |
| Layer has multiple Components | Separate Components within Layer |

---

## Output Format

Write to project's CLAUDE.md:

```markdown
## Directory Structure

```
{directory or file tree based on separation decisions}
```
```
