FROM golang:1.18-alpine as builder

WORKDIR /go/src/app
ADD . /go/src/app

RUN go get -d -v ./...

RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/base
COPY --from=builder /go/bin/app /
CMD ["/app"]
