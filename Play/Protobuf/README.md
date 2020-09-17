## Imagine

- Define Protocol Buffer messages for bookings 

- Implement client

    1) create protobuf request message
    2) transform protobuf request message to JSON request
    3) send JSON request to server
    4) transform JSON reply to protobuf reply message
    
- Implement server

    1) receive JSON request
    2) transform JSON request to protobuf request message
    3) create protobuf reply message
    4) transform protobuf reply message to JSON reply
    5) send JSON reply to client

 ## Create
Create the `booking.pb.go` file
```bash
cd booking
protoc --go_out=plugins=grpc:. *.proto
```

Prepare the packages needed
```bash
cd ..

go get google.golang.org/protobuf/encoding/protojson

go mod init
go mod vendor
```


## Quit Implementing

Editing `booking_client/main.go`

No solution for `Type does not implement 'proto.Message' as some methods are missing: ProtoReflect() Message` found

