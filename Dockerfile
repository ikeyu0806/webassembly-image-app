FROM golang:1.21-bullseye

COPY ./ /app

WORKDIR /app

RUN cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./webassembly/js

RUN cd /app/webassembly && GOARCH=wasm GOOS=js go build -o main.wasm

CMD ["go","run","/app/server/main.go"]
