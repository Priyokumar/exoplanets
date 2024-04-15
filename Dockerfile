FROM golang:1.22.2-bookworm as builder
WORKDIR /exoplanets

ADD configs /exoplanets/configs
ADD fuel /exoplanets/fuel
ADD gravity /exoplanets/gravity
ADD handlers /exoplanets/handlers
ADD middlewares /exoplanets/middlewares
ADD models /exoplanets/models
ADD repository /exoplanets/repository
ADD routes /exoplanets/routes
ADD web /exoplanets/web
ADD go.mod /exoplanets/go.mod
ADD go.sum /exoplanets/go.sum
ADD server.go /exoplanets/server.go

ARG G0111MODULE=on
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build -o exoplanets

FROM alpine:3.18.3
WORKDIR /exoplanets

COPY --from=builder /exoplanets/exoplanets /exoplanets/
COPY --from=builder /exoplanets/go.mod /exoplanets/
COPY --from=builder /exoplanets/go.sum /exoplanets/
COPY --from=builder /exoplanets/web /exoplanets/web/

EXPOSE 1323
RUN ls
ENTRYPOINT /exoplanets/exoplanets


