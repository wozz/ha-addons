FROM ghcr.io/snicker/juicepassproxy:0.2.2

ARG BASHIO_VERSION="v0.16.2"
RUN curl -J -L -o /tmp/bashio.tar.gz \
        "https://github.com/hassio-addons/bashio/archive/${BASHIO_VERSION}.tar.gz" \
    && mkdir /tmp/bashio \
    && tar zxvf \
        /tmp/bashio.tar.gz \
        --strip 1 -C /tmp/bashio \
    \
    && mv /tmp/bashio/lib /usr/lib/bashio \
    && ln -s /usr/lib/bashio/bashio /usr/bin/bashio \
  && apt-get update && apt-get install -y --no-install-recommends \
      jq \
  && rm -rf /var/lib/apt/lists/* \
  && rm -rf /usr/src/*

COPY run.sh /run.sh

ENTRYPOINT /run.sh
