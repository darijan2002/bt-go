# Build the protobuf structures and then build the go executable
run:
	@make proto
	@go run .

# Build the protobuf structures and then build the go executable
compile:
	@make proto
	@go build -o ./build/

# Build the protobuf structures and then test the module
test:
	@make proto
	@go test

# Generates the protobuf structs defined in the protos folder
proto:
	@protoc --go_out=. --proto_path=./protos/ ./protos/*.proto	
