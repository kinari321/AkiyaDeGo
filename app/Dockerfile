FROM golang:1.16 as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/work

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o cmd/main cmd/main.go

 # -----------------------------------
FROM alpine

RUN apk add --no-cache ca-certificates && \
    apk --update add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk del tzdata

WORKDIR /work

COPY --from=builder go/src/work/cmd/main /work/

CMD ["./main"]

EXPOSE 8080