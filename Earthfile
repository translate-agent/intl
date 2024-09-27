VERSION 0.8
FROM alpine
WORKDIR /cldr

cldr:
  ARG cldr_version=45.0
  RUN wget https://unicode.org/Public/cldr/45/cldr-common-$cldr_version.zip
  RUN unzip cldr-common-$cldr_version.zip
  RUN rm cldr-common-$cldr_version.zip
  SAVE ARTIFACT /cldr AS LOCAL .cldr