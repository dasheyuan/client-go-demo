FROM golang:1.17 as builder

WORKDIR /app

COPY . .

ENV GOPROXY="https://goproxy.io,direct"
RUN CGO_ENABLED=0 go build -o ingress-manager main.go

FROM alpine

WORKDIR /app

COPY --from=builder /app/ingress-manager .

CMD ["./ingress-manager"]