# Example: CLI Data Processor — Subsequent (No Split)

## Input (CLAUDE.md)

```markdown
## Application

CLI tool that reads CSV files, validates and transforms records, and outputs results.

## External Interfaces

- stdin/file: CSV input
- stdout/file: Processed output
- stderr: Error reporting

## External Dependencies

- None

## Layer Structure

### Processor (feature-bound)

**Called by:** IO

### IO (cross-feature)

**Called by:** External
```

## Analysis

1. **Innermost layer**: Processor
2. **One level outside**: IO
3. **Judgments in Processor**: Validation rules (required fields, format checks), transformation logic (field mapping, aggregation)
4. **Judgments independent of IO?**: The rules do not depend on IO specifics, but they share a single modeling target — "data processing." No distinct inner modeling target emerges within Processor.

**Result:** No split.

## Output

No change to Layer Structure.
