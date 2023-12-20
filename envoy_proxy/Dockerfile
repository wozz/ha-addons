FROM golang:1.21 as builder

WORKDIR /app
COPY app/ .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# unfortunately, the default envoy images won't run on raspberry pi
# see discussion here: https://github.com/envoyproxy/envoy/issues/23339
# Dockerfile for this image is included in custom_envoy directory
FROM us-docker.pkg.dev/wozz-hass/public/envoyproxy@sha256:a605d96ceb262e396ae8c6d321b54082c7094f54daf1627a5ae41b6120eecb1b

COPY --from=builder /app/main /app/main

RUN chmod a+x /usr/local/bin/envoy

COPY run.sh /
RUN chmod a+x /run.sh

WORKDIR /data

CMD [ "/run.sh" ]
