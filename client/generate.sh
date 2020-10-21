#!/bin/sh

# Path to this plugin
PROTOC_GEN_TS_PATH="./node_modules/.bin/protoc-gen-ts"

# Directory to write generated code to (.js and .d.ts files)
OUT_DIR="./src/proto"

# Directory where all proto files are stored
SCHEMA_DIR="../schema"

rm -rf $OUT_DIR
mkdir $OUT_DIR

protoc \
    -I ${SCHEMA_DIR} \
    --plugin="protoc-gen-ts=${PROTOC_GEN_TS_PATH}" \
    --js_out="import_style=commonjs,binary:${OUT_DIR}" \
    --ts_out="service=grpc-web:${OUT_DIR}" \
    ${SCHEMA_DIR}/*.proto
