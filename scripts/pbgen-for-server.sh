# RUN FROM ROOT DIR!
docker run \
    -v `pwd`/apps/server/interface:/api \
    -v `pwd`/apps/server/generated:/goclient \
    jfbrandhorst/grpc-web-generators \
    protoc -I /api \
        --go_out=plugins=grpc,paths=source_relative:/goclient \
        time_service.proto
