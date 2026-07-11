---
name: cldr-format
description: Guide to understanding how date/time formatting options and locale sequences are structured in the fmt_*.go files.
---

# CLDR Date/Time Formatting Skill

This skill explains how date and time formatting is structured, routed, and resolved across the `fmt_*.go` files.

## Naming Conventions for Options (`fmt_*.go`)

Each formatting file handles a specific combination of options. The files are named using a combination of the following letters representing the option fields:

| Character | Option Field | Description |
|-----------|--------------|-------------|
| `y`       | `Year`       | Numeric or 2-digit year. |
| `m`       | `Month`      | Month width (numeric, 2-digit, long, short, narrow). |
| `d`       | `Day`        | Numeric or 2-digit day. |
| `e`       | `Weekday`    | Weekday width (long, short, narrow) - corresponds to CLDR E symbol. |
| `g`       | `Era`        | Era format (narrow, short, long). |

For example:
* [fmt_d.go](../../../fmt_d.go) handles only the `Day` option.
* [fmt_gm.go](../../../fmt_gm.go) handles `Era` and `Month`.
* [fmt_ymd.go](../../../fmt_ymd.go) handles `Year`, `Month`, and `Day`.

## Design Principles

### 1. Calendar Dispatching
The entrypoint is [NewDateTimeFormat](../../../intl.go#L386). It checks the default calendar for the locale (`cldr.DefaultCalendar(locale)`) and returns the appropriate constructor:
* `gregorianDateTimeFormat(locale, options)`
* `persianDateTimeFormat(locale, options)`
* `buddhistDateTimeFormat(locale, options)`

### 2. Sequence Function Names
Inside each `fmt_*.go` file, functions are defined to construct formatting sequences:
* **Gregorian (Default):** `seq[OptionFields]` (e.g., `seqYearMonthDay` in [fmt_ymd.go](../../../fmt_ymd.go#L10))
* **Persian Calendar:** `seq[OptionFields]Persian` (e.g., `seqYearMonthDayPersian` in [fmt_ymd.go](../../../fmt_ymd.go#L1261))
* **Buddhist Calendar:** `seq[OptionFields]Buddhist` (e.g., `seqYearMonthDayBuddhist` in [fmt_ymd.go](../../../fmt_ymd.go#L1298))

These functions return a `*symbols.Seq` which holds a slice of formatting symbols.

### 3. Locale Resolution
Each `seq...` function performs structural dispatching based on language, script, and region:
* A `switch lang` block handles language-specific ordering and default separators.
* Region-specific sub-dispatching (e.g., `switch region`) is used if countries have different separator or formatting preferences (e.g., `cldr.RegionCL` in Spanish).
* Script-specific checks (e.g., checking `script == cldr.Shaw`) are used where script changes the visual pattern representation.

### 4. Width-Dependent Overrides
Formatting details (like zero-padding or name formats) depend on the requested options width:
* Width check helpers are called (e.g., `opts.Month.numeric()`, `opts.Day.twoDigit()`).
* These checks influence either which symbol is added (e.g., choosing `symbols.Symbol_dd` over `opts.Day.symbol()`) or change the separators dynamically (e.g., using hyphens instead of slashes when both fields are numeric).

### 5. Symbol Sequencing (`symbols.Seq`)
Formatting is built using a sequence of tokens from [internal/symbols/symbols.go](../../../internal/symbols/symbols.go):
* **CLDR Symbols:** Placeholders for formatting functions (e.g., `Symbol_y`, `Symbol_MM`, `Symbol_dd`, `MonthUnit`, `DayUnit`, `Symbol_G`).
* **Text Separators & Constants:** Separators (`'/'`, `'.'`, `'-'`) and localized texts (e.g., `Txt日`, `Txt일`, `Txt年`).

### 6. Compilation Optimization
Calling `seq.Func()` compiles the sequence:
1. Contiguous text symbols/runes are merged into single `cldr.Text` blocks.
2. Variable symbols are bound to functions (e.g., `cldr.YearNumeric(digits)`, `cldr.MonthTwoDigit(digits)`) using the locale's numbering system.
3. It returns an optimized pipeline using `strings.Builder` for zero allocation formatting at runtime.

## Adding Support for a New Locale Override

To add support for a new locale:
1. Identify the correct `fmt_*.go` file based on which option fields are involved.
2. Add a `case cldr.XX:` clause to the `switch lang` block in the appropriate `seq*` function (and calendar variant if necessary).
3. Use `seq.Add(...)` to describe the token sequence. Use the symbol constants for date fields and literal runes or `Txt*` constants for separators/units.
4. If script or region overrides are required, add nested switch blocks.
5. Add width-dependent conditions (e.g., `opts.Month.numeric()`) to output the correct layout structure.
