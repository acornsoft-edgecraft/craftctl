FROM golang:1.19-alpine AS builder

LABEL maintainer="acornsoft"

# Move to working directory (/build).
WORKDIR /build

# Build output 
APP_NAME = edge-summarize
BUILD_DIR = $(PWD)/bin

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .

# Set necessary environment variables needed for our image and build the API server.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o $(BUILD_DIR)/$(APP_NAME) /edge-benchmarks/main.go

# Command to run when starting the container.
ENTRYPOINT [$(BUILD_DIR)/$(APP_NAME)]
