VERSION 0.8
FROM golang:alpine
WORKDIR /intl

x:
  WORKDIR /cldr
  ARG cldr_version=45.0
  RUN wget https://unicode.org/Public/cldr/45/cldr-common-$cldr_version.zip
  RUN unzip cldr-common-$cldr_version.zip
  RUN rm cldr-common-$cldr_version.zip
  SAVE ARTIFACT /cldr AS LOCAL .cldr

cldr:
  COPY --dir +x/cldr .
  COPY go.mod go.sum .
  COPY --dir private/gen private/
  RUN ls -la private
  RUN \
    --mount=type=cache,id=go-mod,target=/go/pkg/mod \
    --mount=type=cache,id=go-build,target=/root/.cache/go-build \
    go run -C private/gen . -dir /intl/cldr > cldr.go
  SAVE ARTIFACT cldr.go AS LOCAL cldr.go