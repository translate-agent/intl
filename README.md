![check workflow](https://github.com/translate-agent/intl/actions/workflows/ci.yaml/badge.svg?event=push)

# intl

[CLDR v46.0](https://cldr.unicode.org/downloads/cldr-46) based date formatting in Golang. The formatting output is identical to the [Intl.DateTimeFormat](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Intl/DateTimeFormat) by the [Node.js v23.4.0](https://nodejs.org/docs/v23.4.0/api/intl.html).

## Requirements

Go version 1.22+.

## DateTimeFormat

| Option                 | Supported |
| ---------------------- | :-------: |
| era                    |    ✅︎    |
| year                   |    ✅︎    |
| month                  |    ✅︎    |
| day                    |    ✅︎    |
| hour                   |    ❌     |
| minute                 |    ❌     |
| second                 |    ❌     |
| fractionalSecondDigits |    ❌     |
| weekday                |    ❌     |
| hourCycle              |    ❌     |
| timeZoneName           |    ❌     |

# Development

The project uses [Earthly](https://earthly.dev) to automate all development tasks that can be run locally and in CI/CD environments.

```shell
✗ earthly doc
TARGETS:
  +init
      init prepares the project for local development
  +cldr [--cldr_version=46.0] [--out=.cldr]
      cldr saves CLDR files to .cldr
  +testdata
      testdata generates test cases and saves to tests.json
  +generate
      generate generates cldr_*.go from CLDR xml
  +test
      test runs unit tests
  +lint [--golangci_lint_version=1.62.0]
      lint runs all linters for golang
  +check
      check verifies code quality by running linters and tests.
```
