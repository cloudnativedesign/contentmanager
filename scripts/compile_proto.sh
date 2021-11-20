#!/bin/bash
# Compile domain types
protoc -I. --go_out=. --go-grpc_out=.  \
internal/proto-files/domain/writtenContent.proto \
internal/proto-files/domain/videoContent.proto \
internal/proto-files/domain/spokenContent.proto

# Compile services
protoc -I. --go_out=. --go-grpc_out=.  \
internal/proto-files/service/writtencontent-service.proto \
internal/proto-files/service/videocontent-service.proto \
internal/proto-files/service/spokencontent-service.proto

