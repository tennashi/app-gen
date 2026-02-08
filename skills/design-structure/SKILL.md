---
name: design-structure
description: Design code structure from layer structure. Derives structure through separation decisions. Use before code generation.
---

# Structure Designer

## Overview

Designs code structure by building a Structure Tree. Each node declares how it is separated from its siblings.

Structure emerges from assessing what separation is needed, not from upfront design.

## Definitions

### Stage

How a node is separated from its siblings. Each stage provides a specific type of independence:

| Stage | Boundary | Independence |
|-------|----------|--------------|
| inline | None | - |
| functions | Function call | Testability |
| files | File | Identifiability |
| packages | Package/Directory | Enforceability |
| services | Process/Network | Deployability |

### Valid Parent-Child Stages

A child's stage is constrained by its parent's stage (physical containment):

| Parent Stage | Valid Child Stages |
|--------------|-------------------|
| services | services, packages, files |
| packages | packages, files, functions, inline |
| files | functions, inline |
| functions | functions, inline |
| inline | inline |

### Structure Tree

A tree of Feature and Layer nodes:

- Top level: Features + Cross-feature Layers
- Under each Feature: Feature-bound Layers
- Under Layers with Components: Components
- Under Features with SubFeatures: SubFeatures
- Each node declares its Stage

Feature is always the outer grouping.

---

## Workflow

### Step 1: Build logical tree

From split-layer output, domain code, and CLAUDE.md, build the tree structure (without Stages):

- Top level: Features (from domain code) + Cross-feature Layers (from split-layer)
- Under each Feature: Feature-bound Layers (from split-layer)
- Under Layers with Components: Components (from split-layer, or from External Interfaces / External Dependencies for Cross-feature Layers)
  - Layers with Components can be marked as omitted (e.g., `Adapter (omit)`), promoting Components to the parent level in the tree
- Under Features with SubFeatures: SubFeatures

This step determines all nodes and their parent-child relationships. The logical structure is fixed here.

### Step 2: Assign Stages top-down

Starting from the top-level sibling group, assign Stages downward. Parent's Stage constrains child's choices (see Valid Parent-Child Stages).

For each sibling group, start from inline and increase only when a specific independence is needed:

| Need | Stage |
|------|-------|
| No separation needed | inline |
| Independent testing of each sibling | functions |
| Human needs to see/navigate structure | files |
| Dependency rules must be enforced by tooling | packages |
| Independent deployment | services |

Notes:
- LLM-as-developer rarely needs packages. Convention via files is usually sufficient.
- files is primarily for human comprehension (table of contents for the codebase).
- functions is for when siblings have independent judgments that need isolated testing.

### Step 3: Derive physical structure

Convert Structure Tree to physical file/directory tree:

| Stage | Physical |
|-------|----------|
| services | Separate project directory (with own entry point) |
| packages | Directory |
| files | File |
| functions | Functions within parent file |
| inline | No new boundary |

Entry point (main.go etc.) is always at the project root.

---

## Output Format

Write to project's CLAUDE.md:

```markdown
## Directory Structure

### Structure Tree

```
{structure tree with stages}
```

### Physical

```
{file/directory tree}
```
```

---

## Examples

### Single Feature, Single Layer

split-layer: Application (feature-bound)
Domain: Task

```
Task (files)
  Application (inline)
```

```
main.go
task.go
```

### Multiple Features, Multiple Layers (small)

split-layer: Entity (feature-bound), UseCase (feature-bound), Adapter[Handler, Repository] (feature-bound), Framework (cross-feature)
Domain: User, Project

```
User (files)
  Entity (inline)
  UseCase (inline)
  Adapter (omit)
    Handler (inline)
    Repository (inline)
Project (files)
  Entity (inline)
  UseCase (inline)
  Adapter (omit)
    Handler (inline)
    Repository (inline)
Framework (files)
```

```
main.go
user.go
project.go
framework.go
```

### Multiple Features, Multiple Layers (enforced)

Same input, but dependency enforcement needed between Features (inner separation not yet needed):

```
User (packages)
  Entity (inline)
  UseCase (inline)
  Adapter (omit)
    Handler (inline)
    Repository (inline)
Project (packages)
  Entity (inline)
  UseCase (inline)
  Adapter (omit)
    Handler (inline)
    Repository (inline)
Framework (packages)
```

```
user/
  user.go
project/
  project.go
framework/
  framework.go
main.go
```

### Multiple Features, Multiple Layers (growing)

Same input, but Features need inner visibility for human navigation:

```
User (packages)
  Entity (files)
  UseCase (files)
  Adapter (omit)
    Handler (files)
    Repository (files)
Project (packages)
  Entity (files)
  UseCase (files)
  Adapter (omit)
    Handler (files)
    Repository (files)
Framework (packages)
  DB (files)
  HTTP (files)
```

```
user/
  entity.go
  usecase.go
  handler.go
  repository.go
project/
  entity.go
  usecase.go
  handler.go
  repository.go
framework/
  db.go
  http.go
main.go
```
