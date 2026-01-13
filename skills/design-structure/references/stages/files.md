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
user.go      // Entity×User, Handler×User, Repository×User
project.go   // Entity×Project, Handler×Project, Repository×Project
task.go      // Entity×Task, Handler×Task, Repository×Task
```

Internal separation: functions or inline (see [functions.md](functions.md), [inline.md](inline.md))

### Layer first (files)

Each file contains all Features for one Layer:

```
entity.go     // User, Project, Task entities
handler.go    // User, Project, Task handlers
repository.go // User, Project, Task repositories
```

Internal separation: functions or inline

## When to use

- Project growing beyond single file
- Multiple developers
- Want logical grouping without package overhead
- Code Units < 30
