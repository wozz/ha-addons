FROM golang:1.22 as builder

WORKDIR /app
COPY app/ .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# unfortunately, the default envoy images won't run on raspberry pi
# see discussion here: https://github.com/envoyproxy/envoy/issues/23339
# Dockerfile for this image is included in custom_envoy directory
# currently: v1.29.0
FROM us-docker.pkg.dev/wozz-hass/public/envoyproxy@sha256:8a0c04a7c8cc75a07878c62bd1db8e8063ff1eb39cfa054ea3e1f7984cab0fdd

COPY --from=builder /app/main /app/main

RUN chmod a+x /usr/local/bin/envoy

COPY run.sh /
RUN chmod a+x /run.sh

WORKDIR /data

CMD [ "/run.sh" ]
