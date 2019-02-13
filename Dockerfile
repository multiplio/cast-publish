# Builder
FROM golang:1.11.4 as builder

WORKDIR /build

# Force the go compiler to use modules
ENV GO111MODULE=on
# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the rest of the source code
COPY . .
# Compile the project
RUN GOPATH=/go GOOS=linux CGO_ENABLED=0 go install -a -installsuffix cgo .

# Deploy
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /app/
COPY --from=builder /go/bin/publish .

EXPOSE 3000
ENTRYPOINT ["/app/publish"]
