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
Horizontal separation. Partitions by technical responsibility with defined dependency direction.

### Feature
Vertical separation within a Layer. Each Layer defines its own Features independently.

### Axis
Direction of separation: by-layer or by-feature.

### Stage
Granularity of separation:

| Stage | Description | Example |
|-------|-------------|---------|
| inline | No separation | all in `main.go` |
| files | Split into files | `user.go`, `handler.go` |
| packages | Split into directories | `user/`, `handler/` |
| services | Split into services | `user-service/` |

## Workflow

1. **Read Layer Structure**
   - Parse CLAUDE.md for `## Layer Structure` section
   - Count Features (domain entities) and Layers

2. **Derive Separation Decisions**
   - For each separation point, decide Axis and Stage
   - Based on counts and implementation volume

3. **Write to CLAUDE.md**
   - Write designed directory structure to project's CLAUDE.md

---

## Separation Flow

Each separation decision = **Axis** (direction) × **Stage** (granularity)

### Step 0: Single File

Everything starts in one file (Stage: inline).

```
main.go
```

### Step 1: Initial Separation

Choose Axis based on counts:

| Condition | Axis |
|-----------|------|
| Feature count > Layer count | Feature |
| Feature count ≤ Layer count | Layer |

Choose Stage based on total volume:

| Volume | Stage |
|--------|-------|
| Small (fits in files) | files |
| Medium/Large | packages |

**Example: Feature axis + files stage**
```
user.go      // all layers for User
project.go   // all layers for Project
task.go      // all layers for Task
```

**Example: Layer axis + files stage**
```
model.go     // all entities
handler.go   // all handlers
repository.go // all repositories
```

**Example: Layer axis + packages stage**
```
entity/
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

**Feature → Layer (files stage):**
```
user/
  model.go
  handler.go
  repository.go
project/
  ...
```

**Layer → Feature (files stage):**
```
handler/
  user.go
  project.go
  task.go
```

### Step 3: Extract Shared Layers

Layers not sharing Features with domain Layers can be extracted independently.

**Key insight:** Layers that share common Features (e.g., Entity, Handler, Repository all have User, Project, Task) can be grouped by Feature. Layers with different Features (e.g., Framework with DB, HTTP) remain independent.

**Example analysis:**
```
entity.go      // Features: User, Project, Task
handler.go     // Features: User, Project, Task
repository.go  // Features: User, Project, Task
framework.go   // Features: DB, HTTP (different!)
```

Entity, Handler, Repository share Features → can group by Feature
Framework has different Features → remains independent

**After grouping by Feature:**
```
user/
  model.go
  handler.go
  repository.go
project/
  ...
framework.go   // stays independent (different Features)
```

**Or with packages stage for Framework:**
```
user/
  ...
project/
  ...
framework/
  db.go
  http.go
```

### Grouping Rules

Grouping is the inverse of separation. Layers can be grouped when:

| Condition | Can Group? |
|-----------|------------|
| Layers share common Features | Yes |
| Layers have different Features | No (remain independent) |

**Grouped (common Features):**
```
user/           // Entity, Handler, Repository for User
  model.go
  handler.go
  repository.go
```

**Not grouped (different Features):**
```
user/
  ...
infrastructure.go   // DB, HTTP - different Features
```

---

## Decision Criteria

### Axis Selection

| Condition | Axis |
|-----------|------|
| Feature count > Layer count | Feature first |
| Feature count ≤ Layer count | Layer first |
| Team owns features | Feature |
| Team owns layers | Layer |

### Stage Selection

| Condition | Stage |
|-----------|-------|
| Few items, small code | inline |
| Multiple items, moderate code | files |
| Many items or large code | packages |
| Independent deployment needed | services |

### When to Separate Further

| Signal | Action |
|--------|--------|
| File > ~300 lines | Consider next stage |
| Multiple responsibilities in unit | Separate by other axis |
| External dependency | Extract to shared layer |

---

## Examples

**Example 1: Small (2 Features, 2 Layers)**

Axis: Layer (2 ≥ 2), Stage: files

```
model.go
repository.go
```

**Example 2: Feature-heavy (5 Features, 2 Layers)**

Axis: Feature (5 > 2), Stage: files

```
user.go
project.go
task.go
comment.go
tag.go
```

**Example 3: Large with Independent Layer**

Layers: Entity, Handler, Repository (share Features: User, Project), Infrastructure (Features: DB, HTTP)

Step 1: Layer axis + files stage
```
entity.go      // User, Project
handler.go     // User, Project
repository.go  // User, Project
infrastructure.go  // DB, HTTP
```

Step 2: Group layers with common Features by Feature
```
user/
  model.go
  handler.go
  repository.go
project/
  ...
infrastructure.go  // different Features, stays independent
```

Step 3: Infrastructure chooses packages stage independently
```
user/
  model.go
  handler.go
  repository.go
project/
  ...
infrastructure/
  db.go
  http.go
```

---

## Output Format

Write to project's CLAUDE.md:

```markdown
## Directory Structure

```
user/
├── model.go
├── handler.go
└── repository.go
project/
└── ...
infrastructure/
└── db.go
```
```
