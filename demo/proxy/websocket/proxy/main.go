package main

import (
	"github.com/iralance/go-gateway/proxy/load_balance"
	"github.com/iralance/go-gateway/proxy/middleware"
	proxy2 "github.com/iralance/go-gateway/proxy/proxy"

	"log"
	"net/http"
)

var (
	addr = "127.0.0.1:2002"
)

func main() {
	rb := load_balance.LoadBalanceFactory(load_balance.LbWeightRoundRobin)
	rb.Add("http://127.0.0.1:2003", "50")
	proxy := proxy2.NewLoadBalanceReverseProxy(&middleware.SliceRouterContext{}, rb)
	log.Println("Starting httpserver at " + addr)
	log.Fatal(http.ListenAndServe(addr, proxy))
}
