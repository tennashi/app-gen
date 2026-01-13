# Files

Separation by splitting into multiple files.

## Characteristics

| Aspect | Value |
|--------|-------|
| Boundary | File |
| Enforcement | Medium (separate compilation units) |
| Suitable for | Small to medium projects |

## Examples

### Feature first (files)

Each file contains all Layers for one Feature:

```
user.{ext}        // Entity×User, Handler×User, Repository×User
project.{ext}     // Entity×Project, Handler×Project, Repository×Project
task.{ext}        // Entity×Task, Handler×Task, Repository×Task
```

Internal separation: functions or inline (see [functions.md](functions.md), [inline.md](inline.md))

### Layer first (files)

Each file contains all Features for one Layer:

```
entity.{ext}      // User, Project, Task entities
handler.{ext}     // User, Project, Task handlers
repository.{ext}  // User, Project, Task repositories
```

Internal separation: functions or inline

## When to use

- Project growing beyond single file
- Multiple developers
- Want logical grouping without package overhead
- Code Units < 30
