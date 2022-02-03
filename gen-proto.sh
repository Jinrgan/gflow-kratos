function genProto {
  DOMAIN=$1
  PROTO_PATH=./handler/$DOMAIN/api
  GO_OUT_PATH=./handler/$DOMAIN/api/gen/v1
  mkdir -p "$GO_OUT_PATH"

  protoc --proto_path="$PROTO_PATH" \
    --proto_path=./third-part \
    --go_out=paths=source_relative:"$GO_OUT_PATH" \
    --go-http_out=paths=source_relative:"$GO_OUT_PATH" \
    --swagger_out=logtostderr=true:"$GO_OUT_PATH" \
    "$DOMAIN".proto
}