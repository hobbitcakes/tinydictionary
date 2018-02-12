FROM golang:1.8
WORKDIR /go/src/app
COPY tinydictionary.go ./

RUN go get -v github.com/gin-gonic/gin
RUN go build tinydictionary.go

EXPOSE 9999
ENTRYPOINT ./tinydictionary
