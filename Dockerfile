# ----------------For GoLang---------------------------
# Start from golang base image
FROM golang:alpine as builder

ADD . /app

# Set the current working directory inside the container 
WORKDIR /app/backend

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 

# Build the Go app
RUN  GOOS=linux go build -ldflags="-s -w" -o main .

# ----------------React---------------------------
# Start from node base image
FROM node:alpine AS node_builder

# 
COPY --from=builder /app/frontend ./

RUN npm install
RUN npm run build

# ----------container which will deploy--------------------
# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /main ./
COPY --from=node_builder /build ./web
# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/main .     

# Expose port 50051 8090 to the outside world
EXPOSE 8090

#Command to run the executable
CMD ["./main"]