# Subsequent: Split Innermost Layer

Used when Layer Structure already exists. Determines whether the current innermost layer should split further.

If only one layer exists, the subsequent procedure does not apply — there is no layer outside to compare against.

## Procedure

1. **Identify the current innermost layer**
2. **Identify the layer one level outside** — the layer that calls it
3. **List judgments in the innermost layer**
4. **Ask: are there judgments that do not depend on the outer layer?**
   - If no → no split needed
   - If yes → those judgments form a new inner layer

## Examples

- `examples/simple-crud/subsequent.md` — No split (single layer)
- `examples/simple-internal-tool/subsequent.md` — No split
- `examples/cli-data-processor/subsequent.md` — No split
- `examples/ecommerce-order-system/subsequent.md` — Split (Domain + UseCase + Adapter)
