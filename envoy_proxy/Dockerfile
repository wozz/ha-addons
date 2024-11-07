FROM golang:1.23 as builder

WORKDIR /app
COPY app/ .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# unfortunately, the default envoy images won't run on raspberry pi
# see discussion here: https://github.com/envoyproxy/envoy/issues/23339
# Dockerfile for this image is included in custom_envoy directory
# currently: v1.32.1
FROM us-docker.pkg.dev/wozz-hass/public/envoyproxy:v1.32.1

COPY --from=builder /app/main /app/main

RUN chmod a+x /usr/local/bin/envoy

COPY run.sh /
RUN chmod a+x /run.sh

WORKDIR /data

CMD [ "/run.sh" ]

# Build arguments
ARG BUILD_ARCH
ARG BUILD_DATE
ARG BUILD_DESCRIPTION
ARG BUILD_NAME
ARG BUILD_REF
ARG BUILD_REPOSITORY
ARG BUILD_VERSION

# Labels
LABEL \
    io.hass.name="${BUILD_NAME}" \
    io.hass.description="${BUILD_DESCRIPTION}" \
    io.hass.arch="${BUILD_ARCH}" \
    io.hass.type="addon" \
    io.hass.version=${BUILD_VERSION} \
    maintainer="Michael Wozniak <wozz@koh.ms>" \
    io.hass.name="${BUILD_NAME}" \
    io.hass.description="${BUILD_DESCRIPTION}" \
    org.opencontainers.image.vendor="wozz/ha-addons" \
    org.opencontainers.image.authors="Michael Wozniak <wozz@koh.ms>" \
    org.opencontainers.image.licenses="MIT" \
    org.opencontainers.image.url="https://github.com/wozz/ha-addons" \
    org.opencontainers.image.source="https://github.com/${BUILD_REPOSITORY}" \
    org.opencontainers.image.documentation="https://github.com/${BUILD_REPOSITORY}/blob/main/README.md" \
    org.opencontainers.image.created=${BUILD_DATE} \
    org.opencontainers.image.revision=${BUILD_REF} \
    org.opencontainers.image.version=${BUILD_VERSION}
