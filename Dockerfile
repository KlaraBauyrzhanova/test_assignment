FROM golang:latest AS builder
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /build
COPY . .
RUN go build -o test_assignment -mod vendor .

FROM alpine:latest
COPY --from=builder /build/test_assignment /usr/local/bin
RUN chmod a+x /usr/local/bin/test_assignment
COPY /migrates/ /migrates/

CMD ["test_assignment"]
