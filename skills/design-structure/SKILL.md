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

Grouping is the inverse of separation.

**Can group?**

| Condition | Can Group? |
|-----------|------------|
| Layers share common Features | Yes |
| Layers have different Features | No (remain independent) |

**Should group?**

| Condition | Action |
|-----------|--------|
| Feature changes > Layer changes | Group |
| Team ownership by Feature | Group |
| Feature count growing | Group |
| Need to see Layer across Features | Don't group |

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
| Multiple items, still small code | functions |
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

## Output Format

Write to project's CLAUDE.md:

```markdown
## Directory Structure

```
{directory or file tree based on separation decisions}
```
```
