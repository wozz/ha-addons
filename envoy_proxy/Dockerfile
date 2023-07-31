FROM golang:1.20 as builder

WORKDIR /app
COPY app/ .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# unfortunately, the default envoy images won't run on raspberry pi
# see discussion here: https://github.com/envoyproxy/envoy/issues/23339
# Dockerfile for this image is included in custom_envoy directory
FROM us-docker.pkg.dev/wozz-hass/public/envoyproxy@sha256:c1f3688d13c99efe7643aaaf6033ed77201aabd6364e637099bae62015823598

COPY --from=builder /app/main /app/main

RUN chmod a+x /usr/local/bin/envoy

COPY run.sh /
RUN chmod a+x /run.sh

WORKDIR /data

CMD [ "/run.sh" ]
