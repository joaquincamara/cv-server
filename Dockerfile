# Alpine is chosen for its small footprint
# compared to Ubuntu

FROM golang:1.16-alpine

# Set destination for COPY

RUN mkdir cv-server
WORKDIR "/cv-server"
# Download Go modules
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the source code
COPY . .

# Build
RUN go build -o main .

# This is for documentation purposes only.
# To actually open the port, runtime parameters
# must be supplied to the docker command.
EXPOSE 8080

CMD ["./main"]