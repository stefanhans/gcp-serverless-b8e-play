
## Imagine

Create a helloworld gRPC/protocolbuffer client/server implementation guided by 
[github.com/grpc/grpc-go/examples/helloworld](https://github.com/grpc/grpc-go/tree/master/examples/helloworld)

## Create
Create the `helloworld.pb.go` file
```bash
cd helloworld
protoc --go_out=plugins=grpc:. *.proto
```

Prepare the packages needed
```bash
cd ..

go mod init
go mod vendor
```

Run the server in one terminal
```bash
cd greeter_server

go run main.go
```

Run the client in another terminal
```bash
cd greeter_client

go run main.go
```

## Play

Add new service definitions to `helloworld.proto`
```proto

// The greeting service definition.
service Greeter {
    // Sends a greeting
    rpc SayHello (HelloRequest) returns (HelloReply) {}
    rpc SayBye (ByeRequest) returns (ByeReply) {}
}

// The request message containing the user's name.
message ByeRequest {
    string name = 1;
}

// The response message containing the goodbyes
message ByeReply {
    string message = 1;
}
```

Recreate the `helloworld.pb.go` file
```bash
cd helloworld
protoc --go_out=plugins=grpc:. *.proto
```

Add new function to `greeter_server/main.go` file
```go

// SayBye implements helloworld.GreeterServer
func (s *server) SayBye(ctx context.Context, in *pb.ByeRequest) (*pb.ByeReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.ByeReply{Message: "Bye " + in.GetName()}, nil
}
```

Replace code in `greeter_client/main.go`
```go

	//r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//log.Printf("Greeting: %s", r.GetMessage())

	r, err := c.SayBye(ctx, &pb.ByeRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Goodbye: %s", r.GetMessage())
```



Run the server in one terminal
```bash
cd greeter_server

go run main.go
```

Run the client in another terminal
```bash
cd greeter_client

go run main.go
```

## Share

Not yet

## Reflect

Cloud Functions are not applicable with gRPC easily.


