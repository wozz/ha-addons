FROM arm64v8/debian:bullseye-slim as build-env
RUN --mount=type=cache,target=/var/cache/apt apt-get update && apt-get install -y \
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
   python3 \
   python3-pip \
   python3-setuptools

RUN curl -L -s -o /usr/local/bin/bazel https://github.com/bazelbuild/bazelisk/releases/latest/download/bazelisk-linux-arm64 && chmod +x /usr/local/bin/bazel

COPY --from=arm64v8/golang:1.21.6-bullseye /usr/local/go/ /usr/local/go/

ENV PATH="/usr/local/go/bin:${PATH}"

ARG ENVOY_VERSION=1.29.0
RUN git clone --branch v${ENVOY_VERSION} https://github.com/envoyproxy/envoy.git

# https://github.com/envoyproxy/envoy/issues/28670
RUN cd envoy && sed -i -e 's/envoy_dependencies_extra()/envoy_dependencies_extra(ignore_root_user_error=True)/' WORKSPACE

# note: this disables a bunch of stuff that probably isn't needed to reduce binary size/compile time
RUN --mount=type=cache,target=/root/.cache cd envoy && bazel build \
      --define tcmalloc=gperftools \
      --define hot_restart=disabled \
      --define google_grpc=disabled \
      --define signal_trace=disabled \
      --define admin_html=disabled \
      --//source/extensions/grpc_credentials/file_based_metadata:enabled=false \
      --//source/extensions/grpc_credentials/aws_iam:enabled=false \
      --//source/extensions/health_checkers/redis:enabled=false \
      --//source/extensions/health_checkers/thrift:enabled=false \
      --//source/extensions/filters/http/aws_lambda:enabled=false \
      --//source/extensions/filters/http/aws_request_signing:enabled=false \
      --//source/extensions/filters/http/bandwidth_limit:enabled=false \
      --//source/extensions/filters/http/basic_auth:enabled=false \
      --//source/extensions/filters/http/cdn_loop:enabled=false \
      --//source/extensions/filters/http/compressor:enabled=false \
      --//source/extensions/filters/http/connect_grpc_bridge:enabled=false \
      --//source/extensions/filters/http/csrf:enabled=false \
      --//source/extensions/filters/http/decompressor:enabled=false \
      --//source/extensions/filters/http/fault:enabled=false \
      --//source/extensions/filters/http/file_system_buffer:enabled=false \
      --//source/extensions/filters/http/gcp_authn:enabled=false \
      --//source/extensions/filters/http/geoip:enabled=false \
      --//source/extensions/filters/http/grpc_field_extraction:enabled=false \
      --//source/extensions/filters/http/grpc_http1_bridge:enabled=false \
      --//source/extensions/filters/http/grpc_http1_reverse_bridge:enabled=false \
      --//source/extensions/filters/http/grpc_json_transcoder:enabled=false \
      --//source/extensions/filters/http/grpc_stats:enabled=false \
      --//source/extensions/filters/http/grpc_web:enabled=false \
      --//source/extensions/filters/http/ip_tagging:enabled=false \
      --//source/extensions/filters/http/jwt_authn:enabled=false \
      --//source/extensions/filters/http/rate_limit_quota:enabled=false \
      --//source/extensions/filters/http/lua:enabled=false \
      --//source/extensions/filters/http/oauth2:enabled=false \
      --//source/extensions/filters/http/on_demand:enabled=false \
      --//source/extensions/filters/http/original_src:enabled=false \
      --//source/extensions/filters/http/tap:enabled=false \
      --//source/extensions/filters/http/set_filter_state:enabled=false \
      --//source/extensions/filters/http/stateful_session:enabled=false \
      --//source/extensions/filters/http/wasm:enabled=true \
      --//source/extensions/filters/network/dubbo_proxy:enabled=false \
      --//source/extensions/filters/network/mongo_proxy:enabled=false \
      --//source/extensions/filters/network/ratelimit:enabled=false \
      --//source/extensions/filters/network/redis_proxy:enabled=false \
      --//source/extensions/filters/network/thrift_proxy:enabled=false \
      --//source/extensions/filters/network/zookeeper_proxy:enabled=false \
      --//source/extensions/filters/udp/dns_filter:enabled=false \
      --//source/extensions/filters/udp/udp_proxy:enabled=false \
      --//source/extensions/filters/udp/udp_proxy/session_filters/dynamic_forward_proxy:enabled=false \
      --//source/extensions/filters/listener/original_dst:enabled=false \
      --//source/extensions/filters/listener/original_src:enabled=false \
      --//source/extensions/filters/listener/proxy_protocol:enabled=false \
      --//source/extensions/filters/listener/tls_inspector:enabled=false \
      --//source/extensions/geoip_providers/maxmind:enabled=false \
      --//source/extensions/stat_sinks/dog_statsd:enabled=false \
      --//source/extensions/stat_sinks/open_telemetry:enabled=false \
      --//source/extensions/stat_sinks/graphite_statsd:enabled=false \
      --//source/extensions/stat_sinks/hystrix:enabled=false \
      --//source/extensions/stat_sinks/statsd:enabled=false \
      --//source/extensions/filters/network/thrift_proxy/router:enabled=false \
      --//source/extensions/filters/network/thrift_proxy/filters/header_to_metadata:enabled=false \
      --//source/extensions/filters/network/thrift_proxy/filters/payload_to_metadata:enabled=false \
      --//source/extensions/filters/network/thrift_proxy/filters/ratelimit:enabled=false \
      --//source/extensions/filters/network/sni_dynamic_forward_proxy:enabled=false \
      --//source/extensions/tracers/dynamic_ot:enabled=false \
      --//source/extensions/tracers/datadog:enabled=false \
      --//source/extensions/tracers/zipkin:enabled=false \
      --//source/extensions/tracers/xray:enabled=false \
      --//source/extensions/tracers/skywalking:enabled=false \
      --//source/extensions/tracers/opentelemetry:enabled=false \
      --//source/extensions/tracers/opentelemetry/resource_detectors/dynatrace:enabled=false \
      --//source/extensions/tracers/opentelemetry/resource_detectors/environment:enabled=false \
      --//source/extensions/tracers/opentelemetry/samplers/always_on:enabled=false \
      --//source/extensions/access_loggers/filters/cel:enabled=false \
      --//source/extensions/access_loggers/grpc:enabled=false \
      --//source/extensions/access_loggers/open_telemetry:enabled=false \
      --//source/extensions/access_loggers/stream:enabled=false \
      --//source/extensions/clusters/redis:enabled=false \
      --//source/extensions/clusters/dynamic_forward_proxy:enabled=false \
      --//source/extensions/clusters/original_dst:enabled=false \
      --//source/extensions/compression/gzip/compressor:enabled=false \
      --//source/extensions/compression/gzip/decompressor:enabled=false \
      --//source/extensions/compression/brotli/compressor:enabled=false \
      --//source/extensions/compression/brotli/decompressor:enabled=false \
      --//source/extensions/compression/zstd/compressor:enabled=false \
      --//source/extensions/compression/zstd/decompressor:enabled=false \
      --//source/extensions/matching/http/cel_input:enabled=false \
      --//source/extensions/matching/input_matchers/cel_matcher:enabled=false \
      --//source/extensions/matching/input_matchers/consistent_hashing:enabled=false \
      --//source/extensions/matching/network/application_protocol:enabled=false \
      --//source/extensions/matching/network/common:enabled=false \
      --//source/extensions/matching/common_inputs/environment_variable:enabled=false \
      --//source/extensions/transport_sockets/tap:enabled=false \
      --//source/extensions/transport_sockets/starttls:enabled=false \
      --//source/extensions/transport_sockets/proxy_protocol:enabled=false \
      --//source/extensions/retry/host/omit_canary_hosts:enabled=false \
      --//source/extensions/transport_sockets/tls/cert_validator/spiffe:enabled=false \
      --//source/extensions/http/stateful_session/cookie:enabled=false \
      --//source/extensions/http/stateful_session/header:enabled=false \
      --//source/extensions/key_value/file_based:enabled=false \
      --//source/extensions/network/dns_resolver/apple:enabled=false \
      --//source/extensions/early_data:enabled=false \
      --//source/extensions/load_balancing_policies/maglev:enabled=false \
      --//source/extensions/load_balancing_policies/ring_hash:enabled=false \
      --//source/extensions/router/cluster_specifiers/lua:enabled=false \
      --//source/extensions/http/cache/file_system_http_cache:enabled=false \
      --//source/extensions/http/cache/simple_http_cache:enabled=false \
      //source/exe:envoy-static && cp /envoy/bazel-bin/source/exe/envoy-static /root/envoy-bin
RUN strip /root/envoy-bin

FROM arm64v8/debian:bullseye-slim
COPY --from=build-env /root/envoy-bin /usr/local/bin/envoy
CMD ["/usr/local/bin/envoy"]

