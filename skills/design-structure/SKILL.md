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
- **Feature-bound**: Has Code Units per Feature (e.g., Domain, Adapter)
- **Cross-feature**: Independent of Features (e.g., Framework, Config)

### Feature
Vertical separation by domain/business concern.

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

### Step 3: Decide whether to expand

**Feature axis → Expand SubFeatures?**

For each Feature, decide whether to expand its SubFeatures to top level.

| Condition | Decision |
|-----------|----------|
| SubFeature count is small and unlikely to grow | Expand |
| SubFeature count is large or likely to grow | Don't expand |
| Domain code has separate files per SubFeature | Expand |
| Domain code has SubFeatures in one file | Don't expand |

Example (expand):
```
admin/         # expanded from user/admin
customer/      # expanded from user/customer
product/
```

Example (don't expand):
```
user/          # admin, customer inside
product/
```

**Layer axis → Expand Components?**

For each Layer, decide whether to expand its Components to top level.

| Condition | Decision |
|-----------|----------|
| Component count is small and unlikely to grow | Expand |
| Component count is large or likely to grow | Don't expand |

Example (expand):
```
handler/       # expanded from adapter/handler
repository/    # expanded from adapter/repository
domain/
```

Example (don't expand):
```
adapter/       # handler, repository inside
domain/
```

### Step 4: Choose stage combination

Decide outer stage (primary axis) and inner stage (secondary axis) together.

**Constraint: Inner stage must be lower than outer stage.**

| Outer | Inner options |
|-------|---------------|
| services | packages, files, functions, inline |
| packages | files, functions, inline |
| files | functions, inline |
| functions | inline |

**Decision table:**

| Want inner separation? | Outer stage | Inner stage |
|------------------------|-------------|-------------|
| Yes, by files | packages | files |
| Yes, by functions | files | functions |
| No | files | inline |

**Examples:**

Feature axis + packages/files:
```
user/              # Feature (packages)
  domain.go        # Layer (files)
  handler.go
  repository.go
```

Feature axis + files/inline:
```
user.go            # Feature (files), Layer is inline
project.go
```

Layer axis + packages/files:
```
domain/            # Layer (packages)
  user.go          # Feature (files)
  project.go
```

Layer axis + files/inline:
```
domain.go          # Layer (files), Feature is inline
handler.go
```

**Principle: Start minimal, grow as needed.**

| Code Unit count | Recommended |
|-----------------|-------------|
| 1-3 | files/inline |
| 4-30 | files/inline or packages/files |
| 31+ | packages/files |

### Step 5: Handle Cross-feature Layers (Feature axis only)

**Only if analyze-layers identified Cross-feature Layers.**

Do NOT add Cross-feature Layers that were not in analyze-layers output.

When using Feature axis and Cross-feature Layers exist:

**Expand Components or not?**

If Cross-feature Layer has Components (e.g., DB, HTTP, Logger), decide whether to expand:

| Condition | Decision |
|-----------|----------|
| Component count is small and unlikely to grow | Expand |
| Component count is large or likely to grow | Don't expand |

Example (don't expand):
```
user/
  ...
framework/         # Cross-feature, not expanded
  db.go
  http.go
  logger.go
```

Example (expand):
```
user/
  ...
db/                # Cross-feature, expanded
  connection.go
  migration.go
http/              # Cross-feature, expanded
  router.go
  middleware.go
```

---

## Separation Priority

1. **Business concern** (Feature → SubFeature)
2. **Technical responsibility** (Layer → Component)

Separate by business concern first, then by technical responsibility.

---

## Examples

### Feature axis + packages stage

```
user/
  domain.go
  handler.go
  repository.go
project/
  domain.go
  handler.go
  repository.go
framework/
  db.go
  http.go
```

### Feature axis + files stage

```
user.go            # domain + handler + repository
project.go
framework.go
```

### Layer axis + packages stage

```
domain/
  user.go
  project.go
handler/
  user.go
  project.go
repository/
  user.go
  project.go
framework/
  db.go
  http.go
```

### Layer axis + files stage

```
domain.go          # user + project
handler.go
repository.go
framework.go
```

### With SubFeature expansion

```
admin/             # user/admin expanded
  domain.go
  handler.go
  repository.go
customer/          # user/customer expanded
  domain.go
  handler.go
  repository.go
product/
  domain.go
  handler.go
  repository.go
```

### With Component expansion

```
domain/
  user.go
  project.go
handler/           # adapter/handler expanded
  user.go
  project.go
repository/        # adapter/repository expanded
  user.go
  project.go
```

---

## Output Format

Write to project's CLAUDE.md:

```markdown
## Directory Structure

Primary axis: {Feature|Layer}
Stage: {inline|functions|files|packages|services}

```
{directory or file tree}
```
```
