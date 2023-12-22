FROM gollumwiki/gollum:v5.3.2

WORKDIR /data
RUN apk add --upgrade \
            --no-cache \
            curl \
            libcurl \
            jq \
    && rm -rf /var/cache/apk/*

COPY docker-run.sh /docker-run.sh
RUN chmod +x /docker-run.sh

ENV GOLLUM_AUTHOR_USERNAME=homeassistant
ENV GOLLUM_AUTHOR_EMAIL=homeassistant@ha.local

ENTRYPOINT ["/docker-run.sh"]
