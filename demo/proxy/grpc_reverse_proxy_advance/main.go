package main

import (
	"fmt"
	"github.com/iralance/go-gateway/proxy/grpc_interceptor"
	"github.com/iralance/go-gateway/proxy/load_balance"
	proxy2 "github.com/iralance/go-gateway/proxy/proxy"
	"github.com/iralance/go-gateway/proxy/public"
	"github.com/mwitkow/grpc-proxy/proxy"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const port = ":50051"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	rb := load_balance.LoadBalanceFactory(load_balance.LbWeightRoundRobin)
	rb.Add("127.0.0.1:50055", "40")

	counter, _ := public.NewFlowCountService("local_app", time.Second)
	grpcHandler := proxy2.NewGrpcLoadBalanceHandler(rb)
	s := grpc.NewServer(
		grpc.ChainStreamInterceptor(
			grpc_interceptor.GrpcAuthStreamInterceptor,
			grpc_interceptor.GrpcFlowCountStreamInterceptor(counter)),
		grpc.CustomCodec(proxy.Codec()),
		grpc.UnknownServiceHandler(grpcHandler))

	fmt.Printf("server listening at %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
