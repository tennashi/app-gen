---
name: design-structure
description: Design directory structure from layer structure. Applies Stage/Axis rules to create concrete directory layouts. Use before code generation.
---

# Directory Structure Designer

## Overview

Transforms layer structure into concrete directory structure by applying Stage and Axis rules. Reads from and writes to project's CLAUDE.md.

## Workflow

1. **Read Layer Structure**
   - Parse CLAUDE.md for `## Layer Structure` section

2. **Count Entities**
   - Search codebase for entity count to determine stage

3. **Apply Stage/Axis Rules**
   - Select stage based on entity count and layer
   - Select axis based on team structure

4. **Write to CLAUDE.md**
   - Write designed directory structure to project's CLAUDE.md

## Input

Read from project's CLAUDE.md:

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

Write to project's CLAUDE.md:

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

CLAUDE.md (Layer Structure):
```markdown
### Entity
### InterfaceAdapter
  - Handler (input)
  - Repository (output)
```

→ Write to CLAUDE.md (3 entities, small):
```markdown
## Directory Structure

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
```

**Example 2: Entity + UseCase + InterfaceAdapter** (full)

CLAUDE.md (Layer Structure):
```markdown
### Entity
### UseCase
### InterfaceAdapter
  - Handler (input)
  - Repository (output)
  - Gateway (output)
```

→ Write to CLAUDE.md (8 entities, medium):
```markdown
## Directory Structure

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
```

**Example 3: InterfaceAdapter only** (no domain logic)

CLAUDE.md (Layer Structure):
```markdown
### InterfaceAdapter
  - Handler (input)
  - Repository (output)
```

→ Write to CLAUDE.md:
```markdown
## Directory Structure

```
dist/
├── handler/
│   └── bookmark.go
└── repository/
    └── sqlite.go
```
```
