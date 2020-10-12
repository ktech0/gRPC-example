package lb

import (
	"fmt"

	"google.golang.org/grpc/resolver"
)

type MyResolver struct {
	target       resolver.Target
	clientConn   resolver.ClientConn
	addressStore map[string][]string
}

func (r *MyResolver) start() {
	addressStrs := r.addressStore[r.target.Endpoint]
	addresses := make([]resolver.Address, len(addressStrs))
	for i, s := range addressStrs {
		addresses[i] = resolver.Address{Addr: s}
	}
	r.clientConn.UpdateState(resolver.State{Addresses: addresses})
}

func (*MyResolver) ResolveNow(o resolver.ResolveNowOptions) {
	fmt.Println("ResolveNow")
}

func (*MyResolver) Close() {}
