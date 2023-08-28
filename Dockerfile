FROM golang:1.21-bullseye

COPY ./ /app

WORKDIR /app

RUN mkdir -p /app/webassembly/js
RUN cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" /app/webassembly/js/wasm_exec.js

RUN cd /app/webassembly && GOARCH=wasm GOOS=js go build -o main.wasm

CMD ["go","run","/app/server/main.go"]
