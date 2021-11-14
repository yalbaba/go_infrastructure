package balancer

import (
	"google.golang.org/grpc/resolver"
)

type Resolver struct {
	address []string
	schema  string
	service string
	cc      resolver.ClientConn
}

func NewResolver(schema string, address []string) resolver.Builder {
	return &Resolver{
		address: address,
		schema:  schema,
	}
}

func (r *Resolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {

	r.cc = cc

	addrList := make([]resolver.Address, len(r.address))
	for k, v := range r.address {
		addrList[k] = resolver.Address{Addr: v}
	}
	r.cc.NewAddress(addrList)

	return r, nil
}

func (r *Resolver) Scheme() string {
	return r.schema
}

func (r *Resolver) ResolveNow(resolver.ResolveNowOptions) {

}

func (r *Resolver) Close() {

}
