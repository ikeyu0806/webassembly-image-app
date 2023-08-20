FROM golang:1.21-bullseye

COPY ./ /app

WORKDIR /app

RUN go get -u github.com/gin-gonic/gin

RUN cd /app/webassembly && GOARCH=wasm GOOS=js go build -o main.wasm

CMD ["go","run","/app/server/main.go"]
