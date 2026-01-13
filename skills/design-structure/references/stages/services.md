# Services

Separation by deploying as independent services.

## Characteristics

| Aspect | Value |
|--------|-------|
| Boundary | Process / Network |
| Enforcement | Very strong (API contract) |
| Suitable for | Large projects, independent scaling/deployment |

## Examples

### Feature first (services)

Each service owns all Layers for one Feature:

```
user-service/
  main.go
  entity/
  handler/
  repository/
project-service/
  main.go
  entity/
  handler/
  repository/
api-gateway/
  main.go
```

Internal separation: packages, files, functions, or inline

### Layer first (services) - Less common

Shared services by technical responsibility:

```
api-service/        // All Handlers
  main.go
  user_handler.go
  project_handler.go
domain-service/     // All Entities + UseCases
  main.go
data-service/       // All Repositories
  main.go
```

Internal separation: packages, files, functions, or inline

## Communication

- REST / HTTP
- gRPC
- Message Queue (async)

## When to use

- Independent deployment required
- Different scaling needs
- Team autonomy
- Technology diversity needed
