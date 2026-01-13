# By Feature

Vertical separation within a Layer. Partition by Feature.

## Axis

Group code by Feature. Each Layer defines its own Features independently.

Examples:
- Entity Layer → User, Project, Order
- UseCase Layer → CreateUser, CreateProject, ...
- InterfaceAdapter Layer → Handler, Repository, Gateway

## Example

```
user/
  entity.go
  usecase.go
  handler.go
  repository.go
project/
  entity.go
  usecase.go
  handler.go
  repository.go
```

## Applied to Each Stage

### inline

Not applicable. Code is not separated, so axis has no effect.

### functions

```go
// user feature
func newUser(id, name string) User { ... }
func saveUser(db *sql.DB, u User) error { ... }
func handleGetUser(w http.ResponseWriter, r *http.Request) { ... }

// project feature
func newProject(id, name string) Project { ... }
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
  entity.go
  handler.go
  repository.go
project/
  entity.go
  handler.go
  repository.go
```

### services

```
user-service/
  entity/
  handler/
  repository/
project-service/
  entity/
  handler/
  repository/
```

## Characteristics

- Easy to find everything about a Feature
- Feature changes contained in one location
- Harder to see all code in a Layer
- Better for team ownership per Feature
