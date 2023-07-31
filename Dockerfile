
# Stage 1: Build the Go server
FROM golang:latest AS builder

WORKDIR /app

# Copy only the necessary files for building the Go server (main.go)
COPY . .

RUN go build -o main ./cmd

FROM mongo:latest

EXPOSE 27017
EXPOSE 8000  

COPY --from=builder /app/main /app/main

CMD mongod & /app/main

