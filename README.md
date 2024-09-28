# intl

CLDR v45.0 based date and number formatting in Golang.

## DateTimeFormat

| Option                 | Supported |
| ---------------------- | :-------: |
| year                   |    ✅︎    |
| month                  |    ❌     |
| day                    |    ❌     |
| hour                   |    ❌     |
| minute                 |    ❌     |
| second                 |    ❌     |
| month                  |    ❌     |
| era                    |    ❌     |
| weekday                |    ❌     |
| fractionalSecondDigits |    ❌     |
| hourCycle              |    ❌     |
| timeZoneName           |    ❌     |

# Development

The project uses [Earthly](https://earthly.dev) to automate all development tasks that can be run locally and in CI/CD environments.

```shell
✗ earthly doc
TARGETS:
  +init
      init prepares the project for local development
  +cldr [--cldr_version=45.0]
      cldr saves CLDR files to .cldr
  +testdata
      testdata generates test cases and saves to tests.json
  +generate
      generate generates cldr.go
  +test
      test runs unit tests
  +lint [--golangci_lint_version=1.61.0]
      lint runs all linters for golang
  +check
      check verifies code quality by running linters and tests.
```
