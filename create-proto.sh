if [ "$1" != "" ]; then
  protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative "$1"
  else
    echo "Please provide a target proto file"
fi