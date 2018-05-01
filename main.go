package main

import (
	"context"
	"log"
	"net/http"

	"app/cors"
	pb "app/grpc"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/urfave/negroni"
	"google.golang.org/grpc"
)

var (
	TonatiuhEndpoint = "localhost"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterArticleServiceHandlerFromEndpoint(ctx, mux, TonatiuhEndpoint, opts)
	if err != nil {
		log.Printf("error: %v", err.Error())
	}

	n := negroni.Classic()
	n.Use(cors.Setup())
	//n.Use(negroni.HandlerFunc(firebase.VerifyToken))
	n.UseHandler(mux)
	if err := http.ListenAndServe(":80", n); err != nil {
		log.Printf("error: %v", err.Error())
	}
}
