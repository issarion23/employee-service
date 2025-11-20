FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY docker .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o employee-service ./cmd/employee-service

FROM gcr.io/distroless/static-debian11
COPY --from=builder /app/employee-service /employee-service
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/employee-service"]