FROM gollumwiki/gollum:v5.3.2

USER root

WORKDIR /data
RUN apk add --update \
            --no-cache \
            curl \
            jq \
    && rm -rf /var/cache/apk/* \
    && adduser -u 1000 -S www-data -G www-data \
    && git config --file /home/www-data/.gitconfig --add safe.directory /data \
    && chown www-data:www-data /home/www-data/.gitconfig

COPY docker-run.sh /docker-run.sh
RUN chmod +x /docker-run.sh

USER www-data

ENV GOLLUM_AUTHOR_USERNAME=homeassistant
ENV GOLLUM_AUTHOR_EMAIL=homeassistant@ha.local

ENTRYPOINT ["/docker-run.sh"]
