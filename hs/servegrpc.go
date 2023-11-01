package hs

// import (
// 	"fmt"
// 	"habitat/pb/hs"

// 	"net"

// 	"google.golang.org/grpc"
// )

// func (h *HabitatServer) ServeGrpc() {
// 	lis, err := net.Listen("tcp", h.grpcAddr)
// 	if err != nil {
// 		panic(err)
// 	}
// 	s := grpc.NewServer()
// 	hs.RegisterHabitatServiceServer(s, h)

// 	if err := s.Serve(lis); err != nil {
// 		panic(err)
// 	}
// }
