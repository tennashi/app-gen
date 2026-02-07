# Example: CLI Data Processor — Initial

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
```

## Analysis

1. **Externals**: stdin/file, stdout/file, stderr
2. **Judgments**: Validation rules (required fields, format checks), transformation logic (field mapping, aggregation)
3. **Result**: Judgments exist → Processor (inner) + IO (outer)

## Output

```markdown
## Layer Structure

### Processor (feature-bound)

**Called by:** IO

### IO (cross-feature)

**Called by:** CLI entry point
```
