FROM golang:1.9.2-alpine AS build

WORKDIR /go/src/github.com/avegao/iot-temp-service

RUN apk add --no-cache --update git glide

COPY glide.yaml glide.yaml
COPY glide.lock glide.lock

RUN glide install

COPY ./ ./

ARG VCS_REF="unknown"
ARG BUILD_DATE="unknown"

RUN go install -ldflags "-X main.buildDate=$BUILD_DATE -X main.commmitHash=$VCS_REF"

########################################################################################################################

FROM alpine:3.6

MAINTAINER "Álvaro de la Vega Olmedilla <alvarodlvo@gmail.com>"

ENV GRPC_VERBOSITY ERROR

RUN addgroup iot-temp-service && \
    adduser -D -G iot-temp-service iot-temp-service

USER iot-temp-service

WORKDIR /app

COPY --from=build /go/bin/iot-temp-service /app/iot-temp-service

EXPOSE 50000/tcp

LABEL com.avegao.iot.temp.vcs_ref=$VCS_REF \
      com.avegao.iot.temp.build_date=$BUILD_DATE \
      maintainer="Álvaro de la Vega Olmedilla <alvarodlvo@gmail.com>"

ENTRYPOINT ["./iot-temp-service"]
