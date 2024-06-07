FROM golang:1.22.0 AS builder
WORKDIR $GOPATH/github.com/jackthomsonn/home-plan
ARG SERVICE_NAME
ADD . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app cmd/${SERVICE_NAME}/main.go

FROM alpine:latest
WORKDIR /root/
RUN mkdir -p ./cmd/app
COPY --from=builder /go/github.com/jackthomsonn/home-plan/app ./app

CMD ["./app"]