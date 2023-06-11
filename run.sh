#!/bin/bash
set -ex

# generate envoy config from options
/app/main

ls -lh /usr/local/bin
# Run Envoy
exec /usr/local/bin/envoy -c /data/envoy.json
