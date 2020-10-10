module server

go 1.15

require (
	github.com/golang/protobuf v1.4.2 // indirect
	google.golang.org/grpc v1.32.0
	google.golang.org/protobuf v1.25.0 // indirect
	mathproto v0.0.0
)

replace mathproto => ../mathproto
