# Feature Axis Structure

Use this when **Feature axis** was selected in Step 2.

## Step 3: Decide whether to expand SubFeatures

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

## Step 4: Choose stage combination

| Want inner separation? | Outer stage | Inner stage |
|------------------------|-------------|-------------|
| Yes, by files | packages | files |
| Yes, by functions | files | functions |
| No | files | inline |

**Feature axis + packages/files:**
```
user/              # Feature (packages)
  entity.go        # Layer (files)
  handler.go
  repository.go
project/
  entity.go
  handler.go
  repository.go
```

**Feature axis + files/inline:**
```
user.go            # Feature (files), Layer is inline
project.go
```

**Principle: Start minimal, grow as needed.**

| Code Unit count | Recommended |
|-----------------|-------------|
| 1-3 | files/inline |
| 4-30 | files/inline or packages/files |
| 31+ | packages/files |

## Step 5: Handle Cross-feature Layers

**Only if analyze-layers identified Cross-feature Layers.**

Do NOT add Cross-feature Layers that were not in analyze-layers output.

If Cross-feature Layer has Components (e.g., DB, HTTP, Logger), decide whether to expand:

| Condition | Decision |
|-----------|----------|
| Component count is small and unlikely to grow | Expand |
| Component count is large or likely to grow | Don't expand |

Example (don't expand):
```
user/
  entity.go
  handler.go
  repository.go
project/
  entity.go
  handler.go
  repository.go
framework/         # Cross-feature, not expanded
  db.go
  http.go
```

Example (expand):
```
user/
  entity.go
  handler.go
  repository.go
db/                # Cross-feature, expanded
  connection.go
http/              # Cross-feature, expanded
  router.go
```

## Output Format

```markdown
## Directory Structure

Primary axis: Feature
Stage: {packages|files}

```
{directory tree with Feature as outer, Layer as inner}
```
```

## Examples

### packages/files (typical)

```
user/
  entity.go
  handler.go
  repository.go
project/
  entity.go
  handler.go
  repository.go
task/
  entity.go
  handler.go
  repository.go
main.go
```

### files/inline (small projects)

```
user.go
project.go
task.go
main.go
```

### With SubFeature expansion

```
admin/             # user/admin expanded
  entity.go
  handler.go
  repository.go
customer/          # user/customer expanded
  entity.go
  handler.go
  repository.go
product/
  entity.go
  handler.go
  repository.go
main.go
```
