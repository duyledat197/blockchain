#!/bin/sh

#* variables
PROTO_PATH=/app/api/protos
PROTO_OUT=/proto_out
IDL_PATH=./
DOC_OUT=/app/docs

#* clean
# rm -rf ${PROTO_OUT}
protoc \
  ${PROTO_PATH}/**/*.proto \
  -I=/usr/local/include \
  --proto_path=${PROTO_PATH} \
  --go_out=:${PROTO_OUT} \
  --validate_out=lang=go:${PROTO_OUT} \
  --go-grpc_out=:${PROTO_OUT} \
  --grpc-gateway_out=:${PROTO_OUT} \
  \
  --openapiv2_out=:${DOC_OUT}/swagger # --event_out=:${IDL_PATH} \
# --enum_out=:${IDL_PATH} \
# --http_out=:${IDL_PATH} \
