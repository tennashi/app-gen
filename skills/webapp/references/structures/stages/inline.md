# Inline

Embed code within adjacent layer instead of separating.

## Characteristics

| Aspect | Value |
|--------|-------|
| Boundary | None |
| Enforcement | Convention |
| Suitable for | Small projects, or layers that don't need separation |

## Use Cases

### 1. Small projects - all in one function

```go
func main() {
    // Domain
    user := User{Name: "alice"}

    // Persistence
    db.Save(user)

    // Presentation
    json.NewEncoder(w).Encode(user)
}
```

### 2. Layer embedding - UseCase in Handler

```go
// UseCase logic is inline within Handler
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
    var req CreateUserRequest
    json.NewDecoder(r.Body).Decode(&req)

    // UseCase logic (inline, not in separate file)
    if req.Name == "" {
        http.Error(w, "name required", http.StatusBadRequest)
        return
    }
    user := domain.NewUser(uuid.New().String(), req.Name)

    // Repository call
    h.repo.Create(r.Context(), user)
    respondJSON(w, http.StatusCreated, user)
}
```

### 3. Layer embedding - Framework in Repository

```go
// Framework (sqlite) is inline within Repository implementation
func (r *SQLiteUserRepository) Create(ctx context.Context, u *domain.User) error {
    // sqlite-specific code inline
    _, err := r.db.ExecContext(ctx,
        "INSERT INTO users (id, name) VALUES (?, ?)",
        u.ID, u.Name)
    return err
}
```

## When to use

- Exploring an idea
- Single-file scripts
- Layers that are simple enough to not need separation
- UseCase layer in small/medium projects
- Framework implementations within Repository
