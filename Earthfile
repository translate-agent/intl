VERSION 0.8
ARG go_version=1.24.5
FROM golang:$go_version-alpine3.22
WORKDIR /intl

# init prepares the project for local development
init:
  BUILD +cldr
  BUILD +generate
  BUILD +testdata

# cldr saves CLDR files to .cldr
cldr:
  WORKDIR /cldr
  ARG cldr_version=46.0
  ARG out=.cldr
  RUN wget https://unicode.org/Public/cldr/$( printf "%.0f" $cldr_version )/cldr-common-$cldr_version.zip
  RUN unzip cldr-common-$cldr_version.zip
  RUN rm cldr-common-$cldr_version.zip
  SAVE ARTIFACT /cldr AS LOCAL $out

# testdata generates test cases and saves to tests.json
testdata:
  ARG node_version=23.11.0
  FROM node:$node_version-alpine
  WORKDIR /intl
  COPY testdata.js .
  COPY --dir +cldr/cldr/common/main .cldr/common/main
  RUN node testdata.js
  SAVE ARTIFACT tests.json AS LOCAL tests.json

# Generate unformatted data.go - gofumpt does not support older go versions.
data:
  COPY --dir +cldr/cldr .
  COPY --dir internal .
  COPY go.mod go.sum .
  RUN \
    --mount=type=cache,id=go-mod,target=/go/pkg/mod \
    --mount=type=cache,id=go-build,target=/root/.cache/go-build \
    go run -C internal/gen . -cldr-dir /intl/cldr -out=/intl
  SAVE ARTIFACT internal/cldr/data.go data.go

# generate generates formatted internal/cldr/data.go from CLDR xml
generate:
  RUN go install mvdan.cc/gofumpt@latest
  COPY go.mod go.sum .
  COPY +data/data.go internal/cldr/
  RUN \
    --mount=type=cache,id=go-mod,target=/go/pkg/mod \
    --mount=type=cache,id=go-build,target=/root/.cache/go-build \
    gofumpt -w internal/cldr/data.go
  SAVE ARTIFACT internal/cldr/data.go data.go AS LOCAL internal/cldr/data.go

# test runs unit tests
test:
  COPY go.mod go.sum *.go .
  COPY --dir internal .
  COPY +testdata/tests.json .
  COPY +data/data.go internal/cldr/
  RUN \
    --mount=type=cache,id=go-mod,target=/go/pkg/mod \
    --mount=type=cache,id=go-build,target=/root/.cache/go-build \
    go test ./...

# lint runs all linters for golang
lint:
  ARG golangci_lint_version=2.1.6
  FROM golangci/golangci-lint:v$golangci_lint_version-alpine
  WORKDIR /intl
  COPY go.mod go.sum *.go .golangci.yml .
  COPY --dir internal .
  COPY +data/data.go internal/cldr/
  COPY +testdata/tests.json .
  RUN \
    --mount=type=cache,id=go-mod,target=/go/pkg/mod \
    --mount=type=cache,id=go-build,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/.cache/golangci_lint \
    golangci-lint run

# check verifies code quality by running linters and tests
check:
  BUILD +test
  BUILD +lint