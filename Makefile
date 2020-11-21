clean:
	rm -rf static/wasm_exec.js
	rm -rf wasm/index.wasm
	rm -rf main

buildIndex:
	GOOS=js GOARCH=wasm go build -o wasm/index.wasm wasm/index.go

buildMain:
	go build -o main main.go

copyExec:
	cp "$(GOROOT)/misc/wasm/wasm_exec.js" static/wasm_exec.js

run:
	./main

image:
	docker build -t hellowoodes/webassembly-sample .

all:
	make clean
	make buildIndex
	make copyExec
	make buildMain
	make run