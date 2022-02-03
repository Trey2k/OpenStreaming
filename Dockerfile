# Start from golang base image
FROM golang:alpine as builder

# ENV GO111MODULE=on

# Add Maintainer info
LABEL maintainer="Trey M <treym2k.dev@gmail.com>"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

WORKDIR /app
COPY ./resources ./resources
COPY ./.env ./.env

# Set the current working directory inside the webApp directory.
WORKDIR /app/webApp

# Copy go mod and sum files 
COPY webApp/go.mod webApp/go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 

# Copy the source from the current directory to the working Directory inside the container 
COPY ./webApp .

# Build the WebApp app
RUN GOOS=linux go build  -o openStreaming ./

# Set the current working directory to the overlayWASM directory
WORKDIR /app/overlayWASM

COPY overlayWASM/go.mod overlayWASM/go.sum ./

RUN go mod download

COPY ./overlayWASM .

# Build the WebAssembly
RUN GOOS=js GOARCH=wasm go build -o overlay.wasm ./


# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/webApp/openStreaming . 
COPY --from=builder /app/overlayWASM/overlay.wasm .
COPY --from=builder /app/resources/ ./resources
COPY --from=builder /app/.env .
RUN mkdir ./logs


# Expose port 8080 to the outside world
EXPOSE 80 443

#Command to run the executable
CMD ["./openStreaming"]