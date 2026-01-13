# Packages

Separation by directories/packages with explicit imports.

## Characteristics

| Aspect | Value |
|--------|-------|
| Boundary | Package/Module |
| Enforcement | Strong (import required, visibility rules) |
| Suitable for | Medium to large projects |

## Examples

### Feature first (packages)

Each package contains all Layers for one Feature:

```
user/
  entity.go      // Entity×User
  handler.go     // Handler×User
  repository.go  // Repository×User
project/
  entity.go      // Entity×Project
  handler.go     // Handler×Project
  repository.go  // Repository×Project
```

Internal separation: files, functions, or inline

### Layer first (packages)

Each package contains all Features for one Layer:

```
entity/
  user.go        // Entity×User
  project.go     // Entity×Project
  task.go        // Entity×Task
handler/
  user.go        // Handler×User
  project.go     // Handler×Project
  task.go        // Handler×Task
repository/
  user.go
  project.go
  task.go
```

Internal separation: files, functions, or inline

## When to use

- Clear architectural boundaries needed
- Enforce dependency direction via imports
- Multiple teams working on different areas
- Reusable packages
- Code Units > 30
