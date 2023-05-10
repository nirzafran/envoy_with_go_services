cd /var/www
# python3 -m http.server 8088 &
./var/server/bin/goserver
cd -

# Start Envoy with the configuration file
envoy -c /etc/envoy/envoy.yaml
