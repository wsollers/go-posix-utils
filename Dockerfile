FROM golang:1.21.3-alpine3.17

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY *.go .
COPY cmd/ ./cmd/
RUN ls -la ./cmd/*

# Build the Go app (implicitly in /app dir)
RUN go build -o /go-posix-utils

# Install sudo
RUN apk --no-cache add ca-certificates tzdata libcap
#RUN set -ex && apk --no-cache add sudo

#RUN sudo setcap 'cap_net_admin,cap_net_raw=eip' /go-posix-utils
RUN setcap cap_net_admin,cap_net_raw=eip /go-posix-utils

#RUN sudo getcap /go-posix-utils
RUN getcap /go-posix-utils

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/go-posix-utils"]
