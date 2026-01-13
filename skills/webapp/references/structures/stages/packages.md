# Packages

Separation by directories/packages with explicit imports.

## Characteristics

| Aspect | Value |
|--------|-------|
| Boundary | Package/Module |
| Enforcement | Strong (import required, visibility rules) |
| Suitable for | Medium to large projects |

## Example

```
project/
├── main.go
├── domain/
│   └── user.go
├── handler/
│   └── user.go
└── repository/
    └── user.go
```

## When to use

- Clear architectural boundaries needed
- Enforce dependency direction via imports
- Multiple teams working on different areas
- Reusable packages
