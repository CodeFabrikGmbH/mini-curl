FROM golang:1.11.10-alpine3.9 as build

RUN apk add --no-cache curl build-base git ca-certificates
RUN go get -u github.com/alexflint/go-arg

WORKDIR /go/src
COPY . code-fabrik.com/mini-curl

RUN echo $(go version)
RUN go build -o /app code-fabrik.com/mini-curl

FROM alpine:3.9
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app /usr/local/bin/app
ENTRYPOINT ["/usr/local/bin/app"]
