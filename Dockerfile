FROM golang:latest AS builder
RUN go version

WORKDIR /go/src/github.com/aplJake/react-nginx-docker-test-1/

COPY main.go .
COPY db db

#RUN set -x && \
#    go get github.com/golang/dep/cmd/dep && \
#    dep ensure -v

# Force the go compiler to use modules
ENV GO111MODULE=on
# RUN export GO111MODULE=on

COPY go.mod .
COPY go.sum .

#This is the ‘magic’ step that will download all the dependencies that are specified in
# the go.mod and go.sum file.
# Because of how the layer caching system works in Docker, the  go mod download
# command will _ only_ be re-run when the go.mod or go.sum file change
# (or when we add another docker instruction this line)
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o main .

# Stage 2 (to create a downsized "container executable", ~7MB)

# If you need SSL certificates for HTTPS, replace `FROM SCRATCH` with:
#
#   FROM alpine:3.7
#   RUN apk --no-cache add ca-certificates
#
FROM scratch
#WORKDIR /root/
COPY --from=builder /go/src/github.com/aplJake/react-nginx-docker-test-1/main /go/bin/
WORKDIR /go/bin/

EXPOSE 5000
ENTRYPOINT ["./main"]