FROM golang:1.21 as gobuilder

WORKDIR /app
COPY proxy/ .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o prxy .

FROM alpine:3.19.0 as pandocbuilder

RUN apk add --upgrade --no-cache \
      curl \
    && rm -rf /var/cache/apk/*

ENV PANDOC_VERSION 3.1.11
ENV PANDOC_DOWNLOAD_URL https://github.com/jgm/pandoc/releases/download/${PANDOC_VERSION}/pandoc-${PANDOC_VERSION}-linux-arm64.tar.gz

RUN mkdir -p /pandoc && \
    cd /pandoc && \
    curl -fsSL "$PANDOC_DOWNLOAD_URL" | tar -xzf - && \
    cd pandoc-$PANDOC_VERSION && \
    mv bin/pandoc /pandoc-bin

FROM gollumwiki/gollum:v5.3.2

WORKDIR /data
RUN apk add --upgrade --no-cache \
            curl \
            libcurl \
            jq \
            supervisor \
    && rm -rf /var/cache/apk/*

ADD Gemfile.lock /Gemfile.lock
RUN bundle install --gemfile=/Gemfile.lock

COPY --from=gobuilder /app/prxy /prxy
COPY --from=pandocbuilder /pandoc-bin /usr/local/bin/pandoc
COPY docker-run.sh /docker-run.sh
RUN chmod +x /docker-run.sh

ADD supervisor /supervisor
ADD templates /templates

ENV GOLLUM_AUTHOR_USERNAME=homeassistant
ENV GOLLUM_AUTHOR_EMAIL=homeassistant@ha.local

ENTRYPOINT ["supervisord","-c","/supervisor/service.conf"]
