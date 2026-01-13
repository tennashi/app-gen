# Inline

No separation - code is embedded directly without function or file boundaries.

## Characteristics

| Aspect | Value |
|--------|-------|
| Boundary | None |
| Enforcement | Convention (comments, ordering) |
| Suitable for | Very small projects, prototypes |

## Examples

### All Code Units inline (single file)

```
// main file

function main() {
    // --- Entity ---
    user = { id: "1", name: "alice" }

    // --- Repository ---
    db.execute("INSERT INTO users ...")

    // --- Handler ---
    response.write(user)
}
```

### Layer embedding - UseCase in Handler

```
function UserHandler.create(request) {
    data = request.body

    // UseCase logic (inline, not in separate function)
    if data.name == "" {
        return error("name required")
    }
    user = newUser(generateId(), data.name)

    repo.create(user)
    return success(user)
}
```

### Layer embedding - Framework in Repository

```
function UserRepository.create(user) {
    // Framework (db) inline within Repository
    db.execute("INSERT INTO users ...", user.id, user.name)
}
```

## When to use

- Exploring an idea / prototyping
- Single-file scripts
- Layers that are simple enough to not need separation
- Internal separation within functions stage
