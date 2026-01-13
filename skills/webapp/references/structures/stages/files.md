# Files

Separation by splitting into multiple files.

## Characteristics

| Aspect | Value |
|--------|-------|
| Boundary | File |
| Enforcement | Medium (separate compilation units) |
| Suitable for | Small to medium projects |

## Example

```
project/
├── main.go
├── user.go
├── user_handler.go
└── user_repository.go
```

## When to use

- Project growing beyond single file
- Multiple developers
- Want logical grouping without package overhead
