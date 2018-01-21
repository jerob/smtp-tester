# https://hub.docker.com/_/golang/
FROM golang

MAINTAINER Jeremie Robert <appydo@gmail.com>

#  Go compilation environment
ENV GOOS=linux
ENV CGO_ENABLED=0

COPY main.go main.go
RUN go build -ldflags '-w -s' -a -installsuffix cgo -o smtp

# https://hub.docker.com/_/scratch/
# This image is most useful in the context of building
# minimal images that contain only a single binary
FROM scratch

MAINTAINER Jeremie Robert <appydo@gmail.com>

# copy static linked smtp executable
COPY --from=0 /go/smtp smtp

# tell how to run this container
CMD ["./smtp"]
