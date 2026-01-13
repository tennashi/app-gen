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
  entity       // Entity×User
  handler      // Handler×User
  repository   // Repository×User
project/
  entity       // Entity×Project
  handler      // Handler×Project
  repository   // Repository×Project
```

Internal separation: files, functions, or inline

### Layer first (packages)

Each package contains all Features for one Layer:

```
entity/
  user         // Entity×User
  project      // Entity×Project
  task         // Entity×Task
handler/
  user         // Handler×User
  project      // Handler×Project
  task         // Handler×Task
repository/
  user
  project
  task
```

Internal separation: files, functions, or inline

## When to use

- Clear architectural boundaries needed
- Enforce dependency direction via imports
- Multiple teams working on different areas
- Reusable packages
- Code Units > 30
