FROM golang:1.19 as builder

WORKDIR /go/src/github.com/dmytroserhiienko02/furniture/api-svc

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /usr/local/bin/api-svc github.com/dmytroserhiienko02/furniture/api-svc

FROM alpine:3.9

COPY --from=builder /usr/local/bin/api-svc /usr/local/bin/api-svc
RUN apk add --no-cache ca-certificates

EXPOSE 8080

ENTRYPOINT ["api-svc"]
