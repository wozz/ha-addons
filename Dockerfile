FROM golang:1.20 as builder

WORKDIR /app
COPY app/ .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM us-docker.pkg.dev/wozz-hass/public/envoyproxy@sha256:3b34ee917af1cbcebadd097fd3937a393d254fa49b47a023f2a10ecfc43ac2a9

COPY --from=builder /app/main /app/main

RUN chmod a+x /usr/local/bin/envoy

COPY run.sh /
RUN chmod a+x /run.sh

WORKDIR /data

CMD [ "/run.sh" ]
