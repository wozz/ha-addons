FROM arm64v8/debian:bullseye-slim as build-env
RUN apt-get update && apt-get install -y \
   curl \
   build-essential \
   libc++-dev \
   libc++abi-dev \
   ninja-build \
   pkg-config \
   unzip \
   automake \
   cmake \
   git \
   g++ \
   golang \
   python3 \
   python3-pip \
   python3-setuptools

RUN curl -L -s -o /usr/local/bin/bazel https://github.com/bazelbuild/bazelisk/releases/latest/download/bazelisk-linux-arm64 && chmod +x /usr/local/bin/bazel

ARG ENVOY_VERSION=1.26.2
RUN git clone --branch v${ENVOY_VERSION} https://github.com/envoyproxy/envoy.git
RUN cd envoy && \
   bazel build \
      --define tcmalloc=gperftools \
      --define hot_restart=disabled \
      --define google_grpc=disabled \
      --define signal_trace=disabled \
      --define deprecated_features=disabled \
      --define admin_html=disabled \
      //source/exe:envoy-static

FROM arm64v8/debian:bullseye-slim
COPY --from=build-env /envoy/bazel-bin/source/exe/envoy-static /usr/local/bin/envoy
CMD ["/usr/local/bin/envoy"]

