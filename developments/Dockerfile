FROM golang:1.22.0-alpine AS build-stage

# RUN apk add build-base

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o app-exe

FROM golang:1.22.0-alpine AS main-stage
COPY --from=build-stage /app/app-exe /app/app-exe

ENTRYPOINT [ "/app/app-exe" ]

