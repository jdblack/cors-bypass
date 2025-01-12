FROM golang:1.23.4 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

COPY . .
RUN go build . 

# Start a new stage from scratch
FROM gcr.io/distroless/base

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/corsless /corsless



# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
ENTRYPOINT ["/corsless"]

