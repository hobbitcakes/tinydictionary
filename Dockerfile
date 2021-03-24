FROM golang:1.10.3 as builder
WORKDIR /go/src/app
RUN go get -v github.com/gin-gonic/gin

COPY tinydictionary.go ./
RUN go build tinydictionary.go

FROM golang:alpine3.7
WORKDIR /go/src/app
COPY --from=builder /go/src/app/tinydictionary .

EXPOSE 9999
ENTRYPOINT ./tinydictionary
