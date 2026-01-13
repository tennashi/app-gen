# Services

Separation by deploying as independent services.

## Characteristics

| Aspect | Value |
|--------|-------|
| Boundary | Process / Network |
| Enforcement | Very strong (API contract) |
| Suitable for | Large projects, independent scaling/deployment |

## Example

```
system/
├── user-service/
│   ├── main.go
│   └── ...
├── project-service/
│   ├── main.go
│   └── ...
└── api-gateway/
    ├── main.go
    └── ...
```

## Communication

- REST / HTTP
- gRPC
- Message Queue (async)

## When to use

- Independent deployment required
- Different scaling needs
- Team autonomy
- Technology diversity needed
