ARG BUILD_FROM=homeassistant/amd64-base:latest
ARG BUILD_ARCH=amd64
FROM golang:latest as builder

WORKDIR /app
COPY src src
WORKDIR /app/src
RUN echo "run compiles"
RUN CGO_ENABLED=0 GOOS=linux GOARCH=$BUILD_ARCH go build -o ../ha-hoymiles

WORKDIR /app

FROM $BUILD_FROM

WORKDIR /ha-hoymiles
COPY --from=builder /app/ha-hoymiles ./
COPY run.sh .
ENTRYPOINT ["/ha-hoymiles/run.sh"]