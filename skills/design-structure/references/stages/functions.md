# Functions

Separation by extracting functions within a file.

## Characteristics

| Aspect | Value |
|--------|-------|
| Boundary | Function call |
| Enforcement | Weak (still in same scope) |
| Suitable for | Small projects, single-file applications |

## Examples

### Feature first (functions)

All Code Units for one Feature, separated by functions:

```
// user.{ext}

// Entity
class User { id, name }
function User.validate() { ... }

// Handler
function newUserHandler(repo) { ... }
function UserHandler.create(request) { ... }

// Repository
function newUserRepository(db) { ... }
function UserRepository.save(user) { ... }
```

### Layer first (functions)

All Code Units for one Layer, separated by functions:

```
// handler.{ext}

// User Handler
function newUserHandler(repo) { ... }
function UserHandler.create(request) { ... }

// Project Handler
function newProjectHandler(repo) { ... }
function ProjectHandler.create(request) { ... }
```

## When to use

- Small applications
- Clear responsibility per function
- Not yet needing file separation
- Internal separation within files stage
