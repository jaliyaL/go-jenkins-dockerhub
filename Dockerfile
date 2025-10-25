FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN go build -o app ./cmd/api

FROM gcr.io/distroless/base
COPY --from=builder /app/app /app/app
EXPOSE 8080
ENTRYPOINT ["/app/app"]
