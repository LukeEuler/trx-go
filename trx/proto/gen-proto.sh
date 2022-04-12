
#!/usr/bin/env bash

set -e -u

cd "$(dirname "$0")"

protoc --proto_path=googleapis --proto_path=tron --go_out=./../ --go-grpc_out=./../ --go_opt=paths=source_relative ./tron/core/contract/*.proto ./tron/core/*.proto ./tron/api/*.proto
