![check workflow](https://github.com/translate-agent/intl/actions/workflows/ci.yaml/badge.svg?event=push)

# intl

[CLDR v46.0](https://cldr.unicode.org/downloads/cldr-46) based date formatting in Golang. The formatting output is identical to the [Intl.DateTimeFormat](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Intl/DateTimeFormat) by the [Node.js v23.4.0](https://nodejs.org/docs/v23.4.0/api/intl.html).

## Requirements

Go version 1.23+.

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
