ARG BUILD_FROM=ghcr.io/home-assistant/amd64-base:3.18

FROM golang:1.21 as build
#LABEL maintainer = "bertusbaron@gmail.com"
ENV LOGSPOUT_HASSIO_VERSION='master'

RUN date

WORKDIR /go/src/github.com/wozz/

RUN git clone --depth 1 https://github.com/wozz/logspout.git -b ${LOGSPOUT_HASSIO_VERSION}
WORKDIR /go/src/github.com/wozz/logspout

# for local development, comment out before commit!
# RUN rm -rf *
# COPY local/logspout/go.* ./
# RUN go mod download
# COPY local/logspout/ .
# end for local development

RUN CGO_ENABLED=0 go build -v -ldflags "-X main.Version=$(cat VERSION)-${LOGSPOUT_HASSIO_VERSION}" -o /go/bin/logspout

FROM $BUILD_FROM
COPY --from=build /go/bin/logspout /logspout

COPY rootfs /
CMD ["/start"]
