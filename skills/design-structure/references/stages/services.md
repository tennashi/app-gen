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
  main.{ext}
  entity/
  handler/
  repository/
project-service/
  main.{ext}
  entity/
  handler/
  repository/
api-gateway/
  main.{ext}
```

Internal separation: packages, files, functions, or inline

### Layer first (services) - Less common

Shared services by technical responsibility:

```
api-service/        // All Handlers
  main.{ext}
  user_handler.{ext}
  project_handler.{ext}
domain-service/     // All Entities + UseCases
  main.{ext}
data-service/       // All Repositories
  main.{ext}
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
