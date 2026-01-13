---
name: design-structure
description: Design directory structure from layer structure. Derives structure through iterative separation decisions. Use before code generation.
---

# Directory Structure Designer

## Overview

Designs directory structure by iteratively deciding how to separate code. Starts from a single file and grows through Feature/Layer separation decisions.

## Definitions

### Layer
Horizontal separation. Partitions by technical responsibility with defined dependency direction.

### Feature
Vertical separation within a Layer. Each Layer defines its own Features independently.

## Workflow

1. **Read Layer Structure**
   - Parse CLAUDE.md for `## Layer Structure` section
   - Count Features (domain entities) and Layers

2. **Derive Initial Separation**
   - Compare Feature count vs Layer count
   - Decide first axis (by-feature or by-layer)

3. **Apply Recursive Separation**
   - For each unit, decide if internal separation is needed
   - Based on implementation volume within the unit

4. **Extract Shared Layers**
   - Identify implementation not tied to specific domain entities (e.g., infrastructure)
   - Separate as shared layer

5. **Write to CLAUDE.md**
   - Write designed directory structure to project's CLAUDE.md

---

## Separation Flow

### Stage 0: Single File

Everything starts in one file.

```
main.go
```

### Stage 1: Initial Separation

| Condition | Decision | Rationale |
|-----------|----------|-----------|
| Feature count > Layer count | Feature separation first | Avoid large layer files |
| Feature count ≤ Layer count | Layer separation first | Simpler structure sufficient |

**Feature separation first:**
```
user.go      // all layers for User
project.go   // all layers for Project
task.go      // all layers for Task
```

**Layer separation first:**
```
model.go     // all entities
handler.go   // all handlers
repository.go // all repositories
```

### Stage 2: Internal Separation

After initial separation, each unit may need further separation.

**Feature → Layer separation:**
When implementation volume per Layer grows within a Feature.

```
user/
  model.go      // Entity layer
  handler.go    // Handler layer
  repository.go // Repository layer (interface)
project/
  ...
```

**Layer → Feature separation:**
When implementation volume per Feature grows within a Layer.

```
handler/
  user.go
  project.go
  task.go
repository/
  user.go
  project.go
  task.go
```

### Stage 3: Extract Shared Layers

Implementation layers (e.g., DB, external API) can be extracted. Same choice applies: Layer or Feature separation.

**Layer separation (shared infrastructure):**
```
user/
  model.go
  handler.go
  repository.go  // interface
project/
  ...
infrastructure/   // shared across domain entities
  db.go           // all DB implementations
  http_client.go  // all external API clients
```

**Feature separation (per-entity infrastructure):**
```
user/
  model.go
  handler.go
  repository.go
  infrastructure/
    mysql.go      // User's DB implementation
project/
  ...
  infrastructure/
    mysql.go      // Project's DB implementation
```

---

## Decision Criteria

### When to Separate Further

| Condition | Action |
|-----------|--------|
| File > ~300 lines | Consider separation |
| Multiple distinct responsibilities in file | Separate |
| Implementation has external dependency | Extract to shared layer |

### Shared Layer Candidates

| Pattern | Extract to |
|---------|------------|
| DB connection, queries | `infrastructure/db.go` |
| HTTP client for external API | `infrastructure/client.go` |
| Shared utilities | `pkg/` or `internal/` |

---

## Examples

**Example 1: Few Features, Few Layers**

Input: 2 Features (User, Task), 2 Layers (Entity, Repository)

→ Feature ≤ Layer, start with Layer separation:
```
model.go       // User, Task entities
repository.go  // save/get for all
```

**Example 2: Many Features, Few Layers**

Input: 5 Features (User, Project, Task, Comment, Tag), 2 Layers

→ Feature > Layer, start with Feature separation:
```
user.go
project.go
task.go
comment.go
tag.go
```

**Example 3: Growth with Shared Layer**

Input: 3 Features, 3 Layers, DB + external API

→ Feature separation → Layer separation → Extract shared:
```
user/
  model.go
  handler.go
  repository.go
project/
  ...
task/
  ...
infrastructure/
  mysql.go
  external_api.go
```

---

## Output Format

Write to project's CLAUDE.md:

```markdown
## Directory Structure

```
dist/
├── user/
│   ├── model.go
│   ├── handler.go
│   └── repository.go
├── project/
│   └── ...
└── infrastructure/
    └── mysql.go
```
```
