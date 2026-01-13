---
name: design-structure
description: Design directory structure from layer structure. Applies Stage/Axis rules to create concrete directory layouts. Use after analyze-layers, before code generation.
---

# Directory Structure Designer

## Overview

Transforms layer structure (from `analyze-layers`) into concrete directory structure by applying Stage and Axis rules.

## Input

Output from `analyze-layers`:

```markdown
## Layer Structure

### Entity
{Responsibility}

### UseCase
{Responsibility}

### InterfaceAdapter
{Responsibility}

**Features:**
- Handler (input): {Description}
- Repository (output): {Description}
```

## Output

```markdown
## Directory Structure

```
dist/
├── domain/
│   ├── user.go
│   └── task.go
├── handler/
│   └── user.go
└── repository/
    └── sqlite.go
```
```

---

## Design Rules

### Stage (how far to separate)

| Stage | Description |
|-------|-------------|
| inline | Embed in adjacent layer |
| functions | Extract as functions |
| files | Split into files |
| packages | Directories with imports |
| services | Microservices |

Reference:
- [inline](references/stages/inline.md)
- [functions](references/stages/functions.md)
- [files](references/stages/files.md)
- [packages](references/stages/packages.md)
- [services](references/stages/services.md)

### Axis (what to separate by)

| Axis | Description |
|------|-------------|
| by-layer | Technical responsibility (handler/, repository/) |
| by-feature | Business capability (user/, order/) |

Reference:
- [by-layer](references/axes/by-layer.md)
- [by-feature](references/axes/by-feature.md)

### expand

Expand sublayers to top level instead of grouping under parent directory.

---

## Stage Selection

Apply stage based on entity count, only to layers present in input.

| Layer (if present) | Small (1-3 entities) | Medium (4-10) | Large (10+) |
|--------------------|---------------------|---------------|-------------|
| Entity | files | packages | packages |
| UseCase | inline | inline | packages |
| InterfaceAdapter | files | packages | packages |

---

## Axis Selection

| Condition | Axis |
|-----------|------|
| Single team, find-by-layer useful | by-layer |
| Multiple teams, feature ownership | by-feature |
| Mixed (layer at top, feature inside) | by-layer + nested by-feature |

---

## Examples

**Example 1: Entity + InterfaceAdapter** (no UseCase)

Input (analyze-layers output):
```markdown
### Entity
### InterfaceAdapter
  - Handler (input)
  - Repository (output)
```

Output (3 entities, small):
```
dist/
├── domain/           # Entity: files
│   ├── user.go
│   └── task.go
├── handler/          # InterfaceAdapter: files, expand
│   ├── user.go
│   └── task.go
└── repository/
    └── sqlite.go
```

**Example 2: Entity + UseCase + InterfaceAdapter** (full)

Input (analyze-layers output):
```markdown
### Entity
### UseCase
### InterfaceAdapter
  - Handler (input)
  - Repository (output)
  - Gateway (output)
```

Output (8 entities, medium):
```
dist/
├── domain/           # Entity: packages
│   ├── order/
│   └── product/
├── usecase/          # UseCase: inline → embedded in handler
├── handler/          # InterfaceAdapter: packages, expand
│   ├── web/
│   └── admin/
├── repository/
│   └── postgres.go
└── gateway/
    └── payment.go
```

**Example 3: InterfaceAdapter only** (no domain logic)

Input (analyze-layers output):
```markdown
### InterfaceAdapter
  - Handler (input)
  - Repository (output)
```

Output:
```
dist/
├── handler/
│   └── bookmark.go
└── repository/
    └── sqlite.go
```
