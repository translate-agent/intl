VERSION 0.8
# renovate: datasource=docker packageName=golang
ARG go_version=1.25.1-alpine3.22
FROM golang:$go_version
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
  # renovate: datasource=docker packageName=node
  ARG node_version=23.11.1
  FROM node:$node_version-alpine
  # renovate: datasource=npm packageName=npm
  ARG npm_version=11.6.1
  RUN npm i -g npm@$npm_version
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
    --mount type=cache,sharing=shared,id=go-mod,target=/go/pkg/mod \
    --mount type=cache,sharing=shared,id=go-build,target=/root/.cache/go-build \
    go run -C internal/gen . -cldr-dir /intl/cldr -out=/intl
  SAVE ARTIFACT internal/cldr/data.go data.go

# generate generates formatted internal/cldr/data.go from CLDR xml
generate:
  RUN go install mvdan.cc/gofumpt@latest
  COPY go.mod go.sum .
  COPY +data/data.go internal/cldr/
  RUN \
    --mount type=cache,sharing=shared,id=go-mod,target=/go/pkg/mod \
    --mount type=cache,sharing=shared,id=go-build,target=/root/.cache/go-build \
    gofumpt -w internal/cldr/data.go
  SAVE ARTIFACT internal/cldr/data.go data.go AS LOCAL internal/cldr/data.go

# test runs unit tests
test:
  COPY go.mod go.sum *.go .
  COPY --dir internal .
  COPY +testdata/tests.json .
  COPY +data/data.go internal/cldr/
  RUN \
    --mount type=cache,sharing=shared,id=go-mod,target=/go/pkg/mod \
    --mount type=cache,sharing=shared,id=go-build,target=/root/.cache/go-build \
    go test ./...

# lint runs all linters for golang
lint:
  # renovate: datasource=docker packageName=golangci/golangci-lint
  ARG golangci_lint_version=2.5.0
  FROM golangci/golangci-lint:v$golangci_lint_version-alpine
  WORKDIR /intl
  COPY go.mod go.sum *.go .golangci.yml .
  COPY --dir internal .
  COPY +data/data.go internal/cldr/
  COPY +testdata/tests.json .
  RUN \
    --mount type=cache,sharing=shared,id=go-mod,target=/go/pkg/mod \
    --mount type=cache,sharing=shared,id=go-build,target=/root/.cache/go-build \
    --mount type=cache,sharing=shared,target=/root/.cache/golangci_lint \
    golangci-lint run

# check verifies code quality by running linters and tests
check:
  BUILD +test
  BUILD +lint