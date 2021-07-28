docker run --rm \
  -p 7880:7880 \
  -p 7881:7881 \
  -p 7882:7882/udp \
  -e LIVEKIT_KEYS="APIvt6eXMbdp2kj: qJefyDYS631CqfJiCeuo8u5KY8LtTuen5iMrHsDOSxSG" \
  livekit/livekit-server \
  --dev \
  --node-ip=127.0.0.1
