# tinydictionary

Small, fake RESTful server that can scale and handle large load. Primarily for load testing.

## Getting Started

Install GO
```
wget https://raw.githubusercontent.com/hobbitcakes/tinydictionary/master/install_go.sh
bash install_go.sh
```


### Prerequisites

This project uses the Gin framework

```
go get github.com/gin-gonic/gin
```

## Build the server binary
```
go build tinydictionary.go
```

## Run the server
```
./tinydictionary& 
```

## Execute a GET
```
curl -XGET http://server:9999/version
```

## Execute a POST

Create a 1 gb file
```
dd if=/dev/zero of=/tmp/file.txt count=1024000 bs=1024
```
POST the file to the running server 
```
curl -X POST -F "file=@/tmp/file.txt" -H "Content-Type: multipart/form-data" http://server:9999/dinosaurs
```
