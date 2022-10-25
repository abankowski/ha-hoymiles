ARG BUILD_FROM=homeassistant/amd64-base:latest
FROM golang:latest as builder

WORKDIR /app
COPY src src
WORKDIR /app/src
RUN echo "run compiles"
RUN CGO_ENABLED=0 GOOS=linux go build -o ../ha-hoymiles

WORKDIR /app

FROM $BUILD_FROM

WORKDIR /ha-hoymiles
COPY --from=builder /app/ha-hoymiles ./
COPY ha-hoymiles.sh .
ENTRYPOINT ["/ha-hoymiles/ha-hoymiles.sh"]