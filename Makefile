# Build the protobuf structures and then build the go executable
run:
	@protoc --go_out=. --proto_path=./protos/ ./protos/*.proto
	@go run .

# Build the protobuf structures and then build the go executable
compile:
	@protoc --go_out=. --proto_path=./protos/ ./protos/*.proto
	@go build -o ./build/

# Generates the protobuf structs defined in the protos folder
proto:
	@protoc --go_out=. --proto_path=.\protos/ ./protos/*.proto	
