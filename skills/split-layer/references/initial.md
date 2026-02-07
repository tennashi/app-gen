# Initial: Establish Layer Boundary

Used when no Layer Structure exists yet. Establishes the boundary between externals and application code.

## Procedure

1. **List externals** — what does this application interface with? (I/O, libraries, frameworks)
2. **List judgments** — what judgments does the application make?
   - For existing codebases: gather from code
   - For greenfield projects: infer from the application description
3. **Result:**
   - If no judgments → single Application layer
   - If judgments exist → Application layer (inner) + external-facing layer (outer)

This step only separates externals from application logic. It does not further split the Application layer.

## Examples

- `examples/simple-crud/initial.md` — No split
- `examples/simple-internal-tool/initial.md` — Split (Application + Adapter)
- `examples/cli-data-processor/initial.md` — Split (Processor + IO)
- `examples/ecommerce-order-system/initial.md` — Split (Application + Adapter)
