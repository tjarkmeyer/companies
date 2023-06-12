FROM golang:1.19 AS build

WORKDIR /go/src

COPY go.mod .
COPY go.sum .

ENV CGO_ENABLED=0
ENV GOPRIVATE=https://github.com/tjarkmeyer/*

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y git-core

RUN go mod download

COPY cmd ./cmd
COPY configs ./configs
COPY internal ./internal
COPY pkg ./pkg

# build
RUN go build -a -installsuffix cgo -o api ./cmd

# runtime
FROM alpine:3.16 AS runtime
COPY --from=build /go/src/api ./

EXPOSE 8080
ENTRYPOINT ["./api"]
