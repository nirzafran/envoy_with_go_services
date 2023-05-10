#!/bin/bash
sh ./scripts/pbgen.sh
sh ./scripts/buildclient.sh
docker build -t nir-envoy-lab ./
