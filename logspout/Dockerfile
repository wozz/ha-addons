ARG BUILD_ARCH

FROM golang:1.23 as build
ENV LOGSPOUT_HASSIO_VERSION="v0.0.1"

WORKDIR /go/src/github.com/wozz/

RUN git clone --depth 1 https://github.com/wozz/logspout.git -b ${LOGSPOUT_HASSIO_VERSION}
WORKDIR /go/src/github.com/wozz/logspout

RUN CGO_ENABLED=0 go build -v -ldflags "-X main.Version=ha-addon-${LOGSPOUT_HASSIO_VERSION}" -o /go/bin/logspout

FROM ghcr.io/home-assistant/${BUILD_ARCH}-base:3.20
COPY --from=build /go/bin/logspout /logspout

COPY rootfs /
CMD ["/start"]
