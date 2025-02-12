VERSION 0.8
ARG golang_version=1.23.4
FROM golang:$golang_version-alpine
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
  FROM node:23.4.0-alpine
  WORKDIR /intl
  COPY testdata.js .
  COPY --dir +cldr/cldr/common/main .cldr/common/main
  RUN node testdata.js
  SAVE ARTIFACT tests.json AS LOCAL tests.json

# generate generates cldr_*.go from CLDR xml
generate:
  RUN go install mvdan.cc/gofumpt@latest
  COPY --dir +cldr/cldr .
  COPY go.mod go.sum .
  COPY --dir internal .
  RUN \
    --mount=type=cache,id=go-mod,target=/go/pkg/mod \
    --mount=type=cache,id=go-build,target=/root/.cache/go-build \
    go run -C internal/gen . -cldr-dir /intl/cldr -out=/intl && \
    gofumpt -w .
  SAVE ARTIFACT cldr_data.go AS LOCAL cldr_data.go

# test runs unit tests
test:
  COPY go.mod go.sum *.go .
  COPY --dir internal .
  COPY +testdata/tests.json .
  COPY +generate/cldr_data.go .
  RUN \
    --mount=type=cache,id=go-mod,target=/go/pkg/mod \
    --mount=type=cache,id=go-build,target=/root/.cache/go-build \
    go test ./...

# lint runs all linters for golang
lint:
  ARG golangci_lint_version=1.62.0
  FROM golangci/golangci-lint:v$golangci_lint_version-alpine
  WORKDIR /intl
  COPY go.mod go.sum *.go .golangci.yml .
  COPY --dir internal .
  COPY +testdata/tests.json .
  COPY +generate/cldr_data.go .
  RUN \
    --mount=type=cache,id=go-mod,target=/go/pkg/mod \
    --mount=type=cache,id=go-build,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/.cache/golangci_lint \
    golangci-lint run ./...

# check verifies code quality by running linters and tests
check:
  BUILD +test
  BUILD +lint