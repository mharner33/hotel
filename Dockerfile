FROM golang:latest as base

# FROM base as dev

# RUN apk update && apk upgrade && \
#    apk add --no-cach bash git openssh

LABEL maintainer="Mike Harner"

# Set the current working directory
WORKDIR /app

# copy go mod and sum files
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

#copy the source from current to working dir
COPY . .

# Build the go app
RUN go build -o bin/api cmd/*.go

# Expose port 3000 to the outside
EXPOSE 3000

# Run the app
CMD ["./bin/api/main"]
