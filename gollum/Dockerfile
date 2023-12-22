FROM gollumwiki/gollum:v5.3.2

RUN git config --file /home/www-data/.gitconfig --add safe.directory /data

USER root

RUN apk add --update \
            --no-cache \
            curl \
            jq \
    && rm -rf /var/cache/apk/*

COPY docker-run.sh /docker-run.sh
RUN chmod +x /docker-run.sh

USER www-data
WORKDIR /data

ENV GOLLUM_AUTHOR_USERNAME=homeassistant
ENV GOLLUM_AUTHOR_EMAIL=homeassistant@ha.local

ENTRYPOINT ["/docker-run.sh"]
