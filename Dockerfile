FROM golang:1.20-bullseye

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

ENV APP_HOME /go/src/savemyrpg
RUN mkdir -p "$APP_HOME"

WORKDIR "$APP_HOME"
EXPOSE 443
CMD ["./savemyrpg"]