# build stage 1
FROM golang:1.19 AS builder

WORKDIR /var/server/bin

WORKDIR /var/server/src
COPY apps/server .

# Fetch project depenencies
RUN go mod download
# Build the go binaries
RUN go build -o /var/server/bin .

WORKDIR /var/server/bin
# Run
# CMD "./goserver" <-- works

# build stage 2
# create a docker file for running envoy proxy
# FROM envoyproxy/envoy:v1.19.0
FROM  envoyproxy/envoy:dev-73420f09ba059d29815b4e528171292683693a97

# install go on container
RUN apt-get update && apt-get install -y golang-go

WORKDIR /var/server/bin
COPY --from=builder /var/server/bin/goserver .

# Set up configuration files
COPY envoy.yaml /etc/envoy/envoy.yaml

WORKDIR /usr/local/bin
COPY docker/initcontainer.sh /usr/local/bin/initcontainer.sh

# Expose ports
EXPOSE 10000/tcp
# EXPOSE 8088/tcp

# install python3 and http server module
# RUN apt-get update && apt-get install -y python3 && apt-get install -y python3-pip && pip3 install httpserver
# RUN apt-get install -y curl

# sleep for 1 second
RUN sleep 1

# init the container
CMD ["sh", "/usr/local/bin/initcontainer.sh"]

