FROM golang:1.23 as gobuild

WORKDIR /app
COPY wsproxy/ .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o wsproxy .

FROM node:22-alpine3.18 as build
WORKDIR /app

ENV MATTER_SERVER_UI_VERSION="v0.2.4"

RUN \
    apk add --no-cache git \
    && git clone https://github.com/wozz/matter-server-ui.git \
    && cd matter-server-ui \
    && git checkout ${MATTER_SERVER_UI_VERSION} \
    && npm install \
        --no-audit \
        --no-fund \
        --no-update-notifier \
    && npm run build

FROM alpine:3.21

COPY --from=build /app/matter-server-ui/build /static
COPY --from=gobuild /app/wsproxy /app/main
EXPOSE 8099

CMD ["/app/main"]
