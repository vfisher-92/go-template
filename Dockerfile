FROM golang:1.18-alpine

ARG VERSION=prod

# Git for go mod download
RUN apk update

WORKDIR /build

# Install dependecies
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy source files
COPY . .

RUN export PATH="$PATH:$(go env GOPATH)/bin"
RUN go build -ldflags "-X main.version=$VERSION" -o bin/go_template ./cmd/server

FROM alpine:3.16
RUN apk --no-cache add ca-certificates

WORKDIR /data/backend

COPY --from=0 /build/bin/go_template /data/backend/

EXPOSE 9000

CMD ["./go_template"]
