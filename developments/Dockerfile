FROM golang:1.22.0-alpine AS build-stage

# RUN apk add build-base

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o app-exe

ENTRYPOINT [ "/app/app-exe" ]

