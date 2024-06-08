ARG BUILD_ARCH
FROM ghcr.io/home-assistant/${BUILD_ARCH}-base-python:3.12-alpine3.19
# WARNING: Do not update python past 3.12.* if telnetlib is still being used.

ENV DEBUG=false

ARG JPP_VERSION="v0.3.1"
RUN curl -J -L -o /tmp/jpp.tar.gz \
      "https://github.com/snicker/juicepassproxy/archive/refs/tags/${JPP_VERSION}.tar.gz" \
    && mkdir /tmp/jpp \
    && tar zxvf \
      /tmp/jpp.tar.gz \
      --strip 1 -C /tmp/jpp \
    && mv /tmp/jpp /juicepassproxy \
  && pip install --root-user-action=ignore --no-cache-dir -r /juicepassproxy/requirements.txt && \
  chmod -f +x /juicepassproxy/*.sh && \
  mkdir -p /log

COPY run.sh /run.sh
ENTRYPOINT /run.sh
