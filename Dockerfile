FROM golang:1.18

WORKDIR /app
COPY go.* ./

RUN go env && go mod download

COPY . .
RUN go build  -tags 'osusergo,netgo' \
    -v -o bin/demo *.go

FROM alpine as running

WORKDIR /app

COPY --from=builder /app/bin /app/

CMD ["/app/demo"]
