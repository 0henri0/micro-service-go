FROM golang:1.22-alpine AS builder

WORKDIR /app
# Download dependencies
COPY go.mod go.sum ./
RUN go mod download
# Add source code
COPY . .

RUN ls

#CMD ["go", "run", "cmd/main.go"]
RUN go build -o main cmd/main.go

# Multi-Stage production build
FROM alpine AS production
RUN apk --no-cache add ca-certificates
WORKDIR /app
# Retrieve the binary from the previous stage
COPY --from=builder /app/main .
# Expose port
EXPOSE 5000
# Set the binary as the entrypoint of the container
CMD ["./main"]
