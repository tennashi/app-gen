# Task Management API

## Tech Stack
- Language: Go 1.21+
- HTTP Router: chi
- Database: SQLite with sqlx

## Structure
LayerDefinition: Clean Architecture
TopLevel: packages, by-layer
  Entity: files, by-feature
  UseCase: inline
  InterfaceAdapter: packages, by-layer, expand
    Handler: files, by-feature
    Repository: files, by-feature
