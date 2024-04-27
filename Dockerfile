FROM golang AS build

WORKDIR /src/sanji
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o sanji main.go

FROM cgr.dev/chainguard/wolfi-base
WORKDIR /app

COPY --from=build /src/sanji/sanji /app/sanji

RUN apk update
RUN apk add --no-cache ffmpeg

VOLUME /config /data

ENTRYPOINT ["/app" "-c" "/config/config.yaml"]