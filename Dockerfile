FROM golang:1.11.10-alpine3.9 as build

WORKDIR /go/src

RUN apk add --no-cache curl build-base git
RUN go get -u github.com/alexflint/go-arg

COPY . code-fabrik.com/mini-curl

RUN echo $(go version)
RUN go build -o /app code-fabrik.com/mini-curl

FROM alpine:3.9
COPY --from=build /app /usr/local/bin/app
ENTRYPOINT ["/usr/local/bin/app"]
