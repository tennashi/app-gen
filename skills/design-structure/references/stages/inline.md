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

```go
// main.go

func main() {
    // --- Entity ---
    type User struct {
        ID   string
        Name string
    }
    user := User{ID: "1", Name: "alice"}

    // --- Repository ---
    db.Exec("INSERT INTO users (id, name) VALUES (?, ?)", user.ID, user.Name)

    // --- Handler ---
    json.NewEncoder(w).Encode(user)
}
```

### Layer embedding - UseCase in Handler

```go
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
    var req CreateUserRequest
    json.NewDecoder(r.Body).Decode(&req)

    // UseCase logic (inline, not in separate function)
    if req.Name == "" {
        http.Error(w, "name required", http.StatusBadRequest)
        return
    }
    user := domain.NewUser(uuid.New().String(), req.Name)

    h.repo.Create(r.Context(), user)
    respondJSON(w, http.StatusCreated, user)
}
```

### Layer embedding - Framework in Repository

```go
func (r *SQLiteUserRepository) Create(ctx context.Context, u *domain.User) error {
    // Framework (sqlite) inline within Repository
    _, err := r.db.ExecContext(ctx,
        "INSERT INTO users (id, name) VALUES (?, ?)",
        u.ID, u.Name)
    return err
}
```

## When to use

- Exploring an idea / prototyping
- Single-file scripts
- Layers that are simple enough to not need separation
- Internal separation within functions stage
