all: hello

hello:
	vyper -f bytecode hello.vy > hello.bin
	vyper -f json hello.vy > hello.abi
	abigen --abi hello.abi --bin hello.bin --type Hello -pkg main --out hello.go

run:
	go run main.go hello.go
