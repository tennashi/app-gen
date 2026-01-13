# By Layer

Horizontal separation. Partition by Layer.

## Axis

Group code by Layer (technical responsibility with defined dependency direction).

## Example

```
entity/
  user.go
  project.go
usecase/
  user.go
  project.go
handler/
  user.go
  project.go
repository/
  user.go
  project.go
```

## Applied to Each Stage

### inline

Not applicable. Code is not separated, so axis has no effect.

### functions

```go
// entity layer
func newUser(id, name string) User { ... }
func newProject(id, name string) Project { ... }

// handler layer
func handleGetUser(w http.ResponseWriter, r *http.Request) { ... }
func handleGetProject(w http.ResponseWriter, r *http.Request) { ... }

// repository layer
func saveUser(db *sql.DB, u User) error { ... }
func saveProject(db *sql.DB, p Project) error { ... }
```

### files

```
entity.go      # User, Project structs
handler.go     # handleGetUser, handleGetProject
repository.go  # saveUser, saveProject
```

### packages

```
entity/
  user.go
  project.go
handler/
  user.go
  project.go
repository/
  user.go
  project.go
```

### services

**Not recommended.** Leads to "distributed monolith". Prefer by-feature at services stage.

## Characteristics

- Easy to find all code in a Layer
- Layer changes contained in one location
- Feature changes touch multiple directories
