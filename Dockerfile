FROM golang:1.9.2-alpine AS build

WORKDIR /go/src/github.com/avegao/iot-temp-service

RUN apk add --no-cache git glide

COPY glide.yaml glide.yaml
COPY glide.lock glide.lock

RUN glide install

COPY ./ ./

RUN go install

########################################################################################################################

FROM alpine:3.6

MAINTAINER "Álvaro de la Vega Olmedilla <alvarodlvo@gmail.com>"

RUN addgroup iot-temp-service && \
    adduser -D -G iot-temp-service iot-temp-service

USER iot-temp-service

WORKDIR /app

COPY --from=build /go/bin/iot-temp-service /app/iot-temp-service

EXPOSE 50000/tcp

LABEL maintainer="Álvaro de la Vega Olmedilla <alvarodlvo@gmail.com>"

ENTRYPOINT ["./iot-temp-service"]
