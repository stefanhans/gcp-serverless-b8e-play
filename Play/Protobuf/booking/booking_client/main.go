package main

import (
	//"context"
	//"log"
	//"os"
	//"time"

	pb "github.com/stefanhans/gcp-serverless-b8e-play/Play/Protobuf/booking"
	"google.golang.org/protobuf/encoding/protojson"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

var (
	user  = &pb.User{Name: "Alice"}
	share = pb.Share{Name: "Tesla Cool",
		Id:   123,
		Type: pb.Share_CAR}
)

func main() {

	// Type does not implement 'proto.Message' as some methods are missing: ProtoReflect() Message
	_, _ := protojson.Marshal(user)
}
