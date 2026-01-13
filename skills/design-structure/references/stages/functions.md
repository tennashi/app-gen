# Functions

Separation by extracting functions.

## Characteristics

| Aspect | Value |
|--------|-------|
| Boundary | Function call |
| Enforcement | Weak (still in same scope) |
| Suitable for | Small projects, single-file applications |

## Example

```go
func main() {
    user := createUser("alice")
    saveUser(user)
    respondWithJSON(w, user)
}

func createUser(name string) User {
    return User{Name: name}
}

func saveUser(u User) {
    db.Save(u)
}

func respondWithJSON(w io.Writer, v any) {
    json.NewEncoder(w).Encode(v)
}
```

## When to use

- Small applications
- Clear responsibility per function
- Not yet needing file separation
