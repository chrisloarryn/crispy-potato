# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/api/main.go

# Final stage
FROM alpine:3.18

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/main .

# Create uploads directory
RUN mkdir -p uploads/avatars uploads/banners

# Expose port
EXPOSE 8080

# Set environment variables
ENV PORT=8080
ENV MONGODB_URI=mongodb://localhost:27017/twittor
ENV JWT_SECRET=MastersOfDevelopment_facebookGroup
ENV STORAGE_PATH=.

# Run the binary
CMD ["./main"]
