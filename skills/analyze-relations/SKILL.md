---
name: analyze-relations
description: Analyze domain models and infer entity relationships from naming conventions. Language-agnostic. Use when understanding data models, designing database schemas, or before generating API code.
---

# Relationship Analyzer

## Overview

Analyzes domain model definitions and infers entity relationships from field naming conventions. Works with any programming language.

## Workflow

1. Read domain model files
2. Parse type/class/struct definitions
3. Infer relationships from field names
4. Output structured relationship information

## Inference Rules

| Pattern (normalized) | Inference |
|---------------------|-----------|
| `{entity}_id` / `{entity}Id` / `{Entity}ID` | belongs to Entity |
| `{entity}_ids` / `{entity}Ids` / `{Entity}IDs` | has many Entities |

### Role-based User References

| Pattern | Relationship |
|---------|-------------|
| `author`, `creator`, `created_by` | → User (as author/creator) |
| `owner` | → User (as owner) |
| `assignee`, `assigned_to` | → User (as assignee) |
| `members`, `participants` | → []User (as members) |
| `reviewer`, `approver` | → User (as reviewer/approver) |

## Output Format

```
## Entity Relationships

### {EntityName}
- belongs to {OtherEntity} via {fieldName}
- has many {OtherEntity} via {fieldName}
- belongs to User as {role} via {fieldName}
```
