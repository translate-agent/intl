# Description

This project implements locale-aware date/time formatting in Go, mirroring the behaviour of JavaScript's Intl.DateTimeFormat. The formatting pipeline transforms (locale, Options, time.Time) into a locale-correct string. It supports Gregorian, Persian, and Buddhist calendars.

# Structure

The project is divided into:

* `.cldr` - CLDR data
* `internal/gen` - code generator
* `internal/intl` - core intl implementation
* `internal/symbols` - formatting symbols
* `*.go` - package level files
* `intl_test.go` - main test file that uses the generated `tests.json` to test the implementation in Golang.

## File Naming Convention for fmt_*.go

Each file is named with a combination of letters representing which option fields it handles:

| Letter | Field |
|--------|-------|
| y      | Year  |
| m      | Month |
| d      | Day   |
| e      | Weekday (Era "E" as in CLDR E symbol) |
| g      | Era (from German "Geschichte" / "G" CLDR symbol) |

# Guardrails

All temporary script files are stored in `.tmp` directory. No need to delete.

# Development flow

The failing test output is very large. When solving issues for specific locale, ensure you run tests for specific locale. E.g. for lv locale - `earth +test --run=TestDateTime_Format/lv` or `go test -run=TestDateTime_Format/lv .`.

Time unit names MUST be generated from the CLDR files.

Unlike ICU which parses the given formatting configuration and finds the closest CLDR pattern to use, this code hardcodes the formatting for the given formatting configuration.

The unit tests log a detailed output of the formatting configuration (same options as Intl.DateTimeFormat in JS) and expected output. E.g.

```
month=numeric,weekday=narrow,out=1 O
month=numeric,day=2-digit,weekday=narrow,out=O, 02.01.
month=2-digit,weekday=long,out=01 Otrdiena
```

The workflow when adding a feature or fixing a bug is as follows:

1) generate test data by running `earth +testdata`
2) generate go code by running `earth +generate`
3) run `earth +lint` to verify the code style
4) run `earth +test` to verify the implementation

##  Rules for Adding or Modifying Formatting

Adding support for a new locale override

1. Identify the correct fmt_*.go file based on which options are involved.
2. Add a case cldr.XX: clause to the switch lang in the appropriate seq* function.
3. Use seq.Add(...) to describe the token sequence — use Symbol constants for date fields and literal runes/Txt* constants for separators.
4. If the locale requires script or region sub-dispatch, add nested switch blocks.
5. For width-dependent behaviour (e.g. numeric vs 2-digit producing different patterns), add conditionals on opts.Month.numeric(), opts.Day.twoDigit(), etc.

# Tools

Use `gofumpt` to format files.

The primary dev tool is "earth":

* `earth +lint` - run lint.
* `earth +test` - run tests.
* `earth +testdata` - generate test cases and save them to `tests.json`.
* `earth +generate` - generate go code from the CLDR data.
* `earth doc` - for more commands.

# Definition of Done

* `earth +check` passes