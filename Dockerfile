FROM golang:1.8
WORKDIR /go/src/app
RUN go get -v github.com/gin-gonic/gin

COPY tinydictionary.go ./
RUN go build tinydictionary.go

EXPOSE 9999
ENTRYPOINT ./tinydictionary
