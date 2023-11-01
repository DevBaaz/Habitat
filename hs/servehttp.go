package hs

// import (
// 	"context"
// 	"habitat/pb/hs"

// 	"net/http"

// 	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
// 	"google.golang.org/grpc"
// )

// func (h *HabitatServer) ServeHttp() {
// 	mux := runtime.NewServerMux()
// 	opts := []grpc.DialOption{grpc.WithInSecure()}
// 	endpoint := h.httpAddr

// 	err := hs.RegisterHabitatServiceHandleFromEndpoint(context.Background(), mux, endpoint, opts)
// 	if err != nil {
// 		panic(err)
// 	}

// 	httpServer := &http.Server{
// 		Addr:    h.httpAddr,
// 		Handler: mux,
// 	}

// 	if err = httpServer.ListenAndServe(); err != nil {
// 		panic(err)
// 	}
// }
