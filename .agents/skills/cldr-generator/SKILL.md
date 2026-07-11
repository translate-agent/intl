---
name: cldr-generator
description: Understand, run, or update the CLDR data generation pipeline that produces internal/cldr/data.go.
---

# CLDR Generator Skill

This skill provides guidance on how to run, update, and debug the CLDR (Common Locale Data Repository) Go code generator in this project.

## Overview

The `internal/gen` tool processes raw XML files from the `.cldr` directory and compiles them into a single optimized Go source file [data.go](../../../internal/cldr/data.go). It deduplicates all localized string tokens into a single read-only constant string (`data`), which is sliced at runtime without memory allocation.

## Generator Behaviour

The generation pipeline works in three distinct phases:

### 1. Parallel XML Decoding & Filtering
* **Parallel Loading**: Parses all locale-specific XML files in `common/main` concurrently using goroutines.
* **Draft Filtering**: Discards any entries that are not fully contributed or approved (draft states other than `""` or `"contributed"`).

### 2. Merging and Inheritance Resolution
CLDR structures use inheritance to avoid redundancy. The generator resolves and merges this data:
* **Aliases**: Replaces alias targets with actual calendar data.
* **Fallback Hierarchies**: Merges parent locale elements (e.g., `root` -> base languages like `en` -> region/territory sub-locales like `en-US`).
* **Local Calendars**: Merges `generic` calendar date/time formatting options into `persian` and `buddhist` calendars.

### 3. String Coalescing & Slicing Optimization
To avoid runtime allocations and keep the binary footprint small:
* All unique string constants (eras, fields, months) are collected and deduplicated (short strings that are substrings of longer strings are removed).
* These strings are compiled into a single global read-only string constant `data` in [data.go](../../../internal/cldr/data.go).
* All lookups in [data.go](../../../internal/cldr/data.go) reference slices of this constant (e.g., `data[start:end]`). Since string slicing does not allocate in Go, this yields zero allocations at runtime.

## Development Workflow

When modifying locale formatting or adding a new locale fallback, use the following sequence:

1. **Modify XML or Generator Code**:
   * To add custom overrides or update the parser, edit [gen.go](../../../internal/gen/gen.go) or [cldr.go](../../../internal/gen/cldr/cldr.go).
   * CLDR templates are in [cldr_data.go.tmpl](../../../internal/gen/cldr_data.go.tmpl).

2. **Run Code Generation**:
   * Generate Go code by running the Earthly recipe:
     ```bash
     earth +generate
     ```
   * Or execute the generator binary directly:
     ```bash
     go run ./internal/gen -cldr-dir .cldr -out .
     ```

3. **Verify Changes**:
   * Run the test suite:
     ```bash
     earth +test
     ```
   * Or for a specific locale (e.g. `lv`):
     ```bash
     go test -run=TestDateTime_Format/lv .
     ```
