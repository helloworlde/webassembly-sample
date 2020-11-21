FROM golang:alpine AS builder

WORKDIR /app
COPY . /app

RUN GOOS=js GOARCH=wasm go build -o wasm/index.wasm wasm/index.go
RUN cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" static/wasm_exec.js
RUN go build -o bin/main main.go

FROM alpine

WORKDIR /app
COPY --from=builder /app/static ./static
COPY --from=builder /app/wasm ./wasm
COPY --from=builder /app/bin ./bin
EXPOSE 8080

CMD ["/app/bin/main"]
