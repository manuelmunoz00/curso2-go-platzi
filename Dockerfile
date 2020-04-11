# FROM golang:alpine

# WORKDIR $GOPATH/src/github.com/manuelmunoz00/calculadora/
# COPY goservidores.go .
# RUN go build -o /go/app

# CMD ["/go/app"]

FROM golang:alpine AS builder

WORKDIR $GOPATH/src/github.com/manuelmunoz00/calculadora/
COPY goservidores.go .
RUN go build -o /go/app

FROM scratch

COPY --from=builder /go/app /go/app

CMD ["/go/app"]