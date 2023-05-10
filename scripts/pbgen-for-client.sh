#!/bin/bash
TARGET_DIR=apps/client/src/generated/grpc/jsclient
mkdir -p $TARGET_DIR
npx grpc_tools_node_protoc \
  --proto_path=apps/server/interface \
  --js_out=import_style=commonjs:$TARGET_DIR \
  --grpc-web_out=import_style=commonjs,mode=grpcwebtext:$TARGET_DIR \
  --plugin=protoc-gen-grpc-web=./node_modules/.bin/protoc-gen-grpc-web \
  apps/server/interface/time_service.proto