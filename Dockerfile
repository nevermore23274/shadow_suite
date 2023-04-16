# syntax=docker/dockerfile:1

FROM golang:1.20 AS build-stage

# Build needed requirements
COPY /build/requirements.sh .
RUN ./requirements.sh

# Set destination for COPY
WORKDIR /app

# Copy all project files into directory
ADD . /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Download fyne modules
RUN go get fyne.io/fyne/v2@latest
RUN go install fyne.io/fyne/v2/cmd/fyne@latest

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY *.go ./

# Build go binary
RUN go build -o /go_network_tool

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./..

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-realease-stage

WORKDIR /

COPY --from=build-stage /go_network_tool /go_network_tool

USER nonroot:nonroot

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
# EXPOSE 8080

# Run
# CMD ["/go_network_tool"]
ENTRYPOINT ["/go_network_tool"]
