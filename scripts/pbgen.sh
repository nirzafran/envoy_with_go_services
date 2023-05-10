#!/bin/bash
echo "Generating protobuf files for server and client..."
sh ./scripts/pbgen-for-server.sh
sh ./scripts/pbgen-for-client.sh
echo "Done."
