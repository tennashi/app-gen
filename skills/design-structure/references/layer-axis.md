# Layer Axis Structure

Use this when **Layer axis** was selected in Step 2.

## Step 3: Decide whether to expand Components

For each Layer, decide whether to expand its Components to top level.

| Condition | Decision |
|-----------|----------|
| Component count is small and unlikely to grow | Expand |
| Component count is large or likely to grow | Don't expand |

Example (expand):
```
handler/       # expanded from adapter/handler
repository/    # expanded from adapter/repository
entity/
```

Example (don't expand):
```
adapter/       # handler, repository inside
entity/
```

## Step 4: Choose stage combination

| Want inner separation? | Outer stage | Inner stage |
|------------------------|-------------|-------------|
| Yes, by files | packages | files |
| Yes, by functions | files | functions |
| No | files | inline |

**Layer axis + packages/files:**
```
entity/            # Layer (packages)
  user.go          # Feature (files)
  project.go
handler/
  user.go
  project.go
repository/
  user.go
  project.go
```

**Layer axis + files/inline:**
```
entity.go          # Layer (files), Feature is inline
handler.go
repository.go
```

**Principle: Start minimal, grow as needed.**

| Code Unit count | Recommended |
|-----------------|-------------|
| 1-3 | files/inline |
| 4-30 | files/inline or packages/files |
| 31+ | packages/files |

## Step 5: Handle Cross-feature Layers

Cross-feature Layers (e.g., Framework) are placed at the same level as other Layers.

**Only if analyze-layers identified Cross-feature Layers.**

Do NOT add Cross-feature Layers that were not in analyze-layers output.

Example:
```
entity/
  user.go
  project.go
handler/
  user.go
  project.go
repository/
  user.go
  project.go
framework/         # Cross-feature Layer
  db.go
  http.go
```

## Output Format

```markdown
## Directory Structure

Primary axis: Layer
Stage: {packages|files}

```
{directory tree with Layer as outer, Feature as inner}
```
```

## Examples

### packages/files (typical)

```
entity/
  user.go
  project.go
  task.go
handler/
  user.go
  project.go
  task.go
repository/
  user.go
  project.go
  task.go
main.go
```

### files/inline (small projects)

```
entity.go
handler.go
repository.go
main.go
```

### With Component expansion

```
entity/
  user.go
  project.go
handler/           # adapter/handler expanded
  user.go
  project.go
repository/        # adapter/repository expanded
  user.go
  project.go
main.go
```
