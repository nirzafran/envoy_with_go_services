# RUN FROM ROOT DIR!
docker run \
    -v `pwd`/apps/server/interface:/api \
    -v `pwd`/apps/server/generated:/goclient \
    -v `pwd`/apps/client/src/jsclient:/jsclient \
    jfbrandhorst/grpc-web-generators \
    protoc -I /api \
        --go_out=plugins=grpc,paths=source_relative:/goclient \
        --js_out=import_style=commonjs:/jsclient \
        --grpc-web_out=import_style=commonjs,mode=grpcwebtext:/jsclient \
        time_service.proto
