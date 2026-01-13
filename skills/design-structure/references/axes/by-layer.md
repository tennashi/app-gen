# By Layer

Partition by technical responsibility.

## Axis

Group code by what it does technically.

## Example Names

| Layer | Possible Names |
|-------|---------------|
| Domain | `domain/`, `model/`, `entity/` |
| Handler | `handler/`, `controller/`, `api/` |
| Repository | `repository/`, `store/`, `persistence/` |
| UseCase | `usecase/`, `service/`, `application/` |

## Applied to Each Stage

### inline

Not applicable. Code is not separated, so axis has no effect.

### functions

```go
// domain layer
func newUser(id, name string) User { return User{ID: id, Name: name} }
func newProject(id, name string) Project { return Project{ID: id, Name: name} }

// repository layer
func saveUser(db *sql.DB, u User) error { ... }
func saveProject(db *sql.DB, p Project) error { ... }

// handler layer
func handleGetUser(w http.ResponseWriter, r *http.Request) { ... }
func handleGetProject(w http.ResponseWriter, r *http.Request) { ... }
```

### files

```
domain.go      # User, Project structs
handler.go     # handleGetUser, handleGetProject
repository.go  # saveUser, saveProject
```

### packages

```
domain/
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

**Not recommended.** Leads to "distributed monolith" - feature changes require coordinated deployment across multiple services. Prefer by-feature at services stage.

```
api-gateway/         # all handlers
  user_handler.go
  project_handler.go
domain-service/      # all domain + usecase
persistence-service/ # all repositories
```

## Characteristics

- Easy to find all handlers, all repositories
- Change to one layer affects one location
- Feature changes touch multiple directories
