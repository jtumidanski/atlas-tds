# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.20-alpine3.17 AS build-env

# Copy the local package files to the container's workspace.

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN apk add --no-cache git
RUN apk add make

ADD ./atlas.com/tds /atlas.com/tds
WORKDIR /atlas.com/tds

RUN go build -o /server

FROM alpine:3.17

# Port 8080 belongs to our application
EXPOSE 8080

RUN apk add --no-cache libc6-compat

WORKDIR /

COPY --from=build-env /server /
COPY /atlas.com/tds/config.yaml /

CMD ["/server"]
