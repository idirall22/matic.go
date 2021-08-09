solc:
	solc --abi --bin --pretty-json --overwrite contracts/Todo.sol -o build

test:
	go test -v ./...