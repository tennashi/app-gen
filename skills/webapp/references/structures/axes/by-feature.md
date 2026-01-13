# By Feature

Partition by business capability.

## Axis

Group code by what business function it serves.

## Example Names

Based on domain entities or business capabilities:
- `user/`
- `project/`
- `billing/`
- `notification/`

## Applied to Each Stage

### inline

Not applicable. Code is not separated, so axis has no effect.

### functions

```go
// user feature
func newUser(id, name string) User { return User{ID: id, Name: name} }
func saveUser(db *sql.DB, u User) error { ... }
func handleGetUser(w http.ResponseWriter, r *http.Request) { ... }

// project feature
func newProject(id, name string) Project { return Project{ID: id, Name: name} }
func saveProject(db *sql.DB, p Project) error { ... }
func handleGetProject(w http.ResponseWriter, r *http.Request) { ... }
```

### files

```
user.go      # User struct, saveUser, handleGetUser
project.go   # Project struct, saveProject, handleGetProject
```

### packages

```
user/
  domain.go
  handler.go
  repository.go
project/
  domain.go
  handler.go
  repository.go
```

### services

```
user-service/
  domain/
  handler/
  repository/
project-service/
  domain/
  handler/
  repository/
```

## Characteristics

- Easy to find everything about a feature
- Feature changes contained in one location
- Harder to see all handlers, all repositories
- Better for team ownership per feature
