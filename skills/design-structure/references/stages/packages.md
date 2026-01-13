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
  entity.{ext}       // Entity×User
  handler.{ext}      // Handler×User
  repository.{ext}   // Repository×User
project/
  entity.{ext}       // Entity×Project
  handler.{ext}      // Handler×Project
  repository.{ext}   // Repository×Project
```

Internal separation: files, functions, or inline

### Layer first (packages)

Each package contains all Features for one Layer:

```
entity/
  user.{ext}         // Entity×User
  project.{ext}      // Entity×Project
  task.{ext}         // Entity×Task
handler/
  user.{ext}         // Handler×User
  project.{ext}      // Handler×Project
  task.{ext}         // Handler×Task
repository/
  user.{ext}
  project.{ext}
  task.{ext}
```

Internal separation: files, functions, or inline

## When to use

- Clear architectural boundaries needed
- Enforce dependency direction via imports
- Multiple teams working on different areas
- Reusable packages
- Code Units > 30
