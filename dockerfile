# syntax=docker/dockerfile:1
FROM golang:1.23 AS builder

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o backend .

FROM alpine:latest

# Set the timezone to Asia/Bangkok
RUN apk update && apk add --no-cache tzdata && cp /usr/share/zoneinfo/Asia/Bangkok /etc/localtime && echo "Asia/Bangkok" > /etc/timezone

RUN mkdir /config

# Copy the built binary from the previous stage
COPY --from=builder /app/backend /usr/local/bin/backend

# Expose the port your application listens on
# EXPOSE 8000

# Set the command to start your application
CMD ["/usr/local/bin/backend"]