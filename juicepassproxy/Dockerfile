FROM ghcr.io/home-assistant/aarch64-base-python:3.12-alpine3.19
# WARNING: Do not update python past 3.12.* if telnetlib is still being used.

ENV DEBUG=false
ENV UPDATE_UDPC=false

ARG JPP_VERSION="v0.2.2"
RUN curl -J -L -o /tmp/jpp.tar.gz \
      "https://github.com/snicker/juicepassproxy/archive/refs/tags/${JPP_VERSION}.tar.gz" \
    && mkdir /tmp/jpp \
    && tar zxvf \
      /tmp/jpp.tar.gz \
      --strip 1 -C /tmp/jpp \
    && mv /tmp/jpp /juicepassproxy \
  && pip install --no-cache-dir -r /juicepassproxy/requirements.txt && \
  chmod -f +x /juicepassproxy/*.sh

COPY run.sh /run.sh
ENTRYPOINT /run.sh
