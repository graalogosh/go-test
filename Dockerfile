FROM golang:1.13.4 as builder
WORKDIR /go/src/github.com/graalogosh/go-test/
COPY go-test.go .
ENV GOOS=linux
RUN go build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder go-test .
CMD ["./go-test"]