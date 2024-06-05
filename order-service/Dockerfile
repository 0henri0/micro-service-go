FROM golang:1.22-alpine AS builder

WORKDIR /app
# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

RUN ls -la
# Add source code
COPY . .

#CMD ["go", "run", "cmd/main.go"]
RUN go build -o main cmd/main.go

# Multi-Stage production build
FROM alpine AS production
RUN apk --no-cache add ca-certificates

WORKDIR /app
# Retrieve the binary from the previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/internal/config/config.yaml .
# Expose port
EXPOSE 5003
# Set the binary as the entrypoint of the container
CMD ["./main"]
