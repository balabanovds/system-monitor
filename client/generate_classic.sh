#!/bin/sh

# Path to
# Directory to write generated code to (.js and .d.ts files)
OUT_DIR="./src/proto_classic"

# Directory where all proto files are stored
SCHEMA_DIR="../schema"

rm -rf $OUT_DIR
mkdir $OUT_DIR

PATH=$PATH:./node_modules/.bin protoc \
    -I=${SCHEMA_DIR} \
    --js_out=import_style=commonjs:$OUT_DIR \
    --grpc-web_out=import_style=commonjs+dts,mode=grpcwebtext:$OUT_DIR \
    ${SCHEMA_DIR}/*.proto
    