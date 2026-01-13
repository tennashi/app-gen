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

```go
// user.go

// Entity
type User struct {
    ID   string
    Name string
}

func (u *User) Validate() error {
    if u.Name == "" {
        return errors.New("name required")
    }
    return nil
}

// Handler
func NewUserHandler(repo UserRepository) *UserHandler { ... }

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) { ... }

// Repository
func NewUserRepository(db *sql.DB) *UserRepository { ... }

func (r *UserRepository) Save(ctx context.Context, u *User) error { ... }
```

### Layer first (functions)

All Code Units for one Layer, separated by functions:

```go
// handler.go

// User Handler
func NewUserHandler(repo UserRepository) *UserHandler { ... }
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) { ... }

// Project Handler
func NewProjectHandler(repo ProjectRepository) *ProjectHandler { ... }
func (h *ProjectHandler) Create(w http.ResponseWriter, r *http.Request) { ... }
```

## When to use

- Small applications
- Clear responsibility per function
- Not yet needing file separation
- Internal separation within files stage
