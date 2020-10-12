package lb

import "google.golang.org/grpc/resolver"

type MyResolverBuilder struct {
	SchemeName  string
	ServiceName string
	AddressList []string
}

func (m *MyResolverBuilder) Build(target resolver.Target, clientConn resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &MyResolver{
		target:     target,
		clientConn: clientConn,
		addressStore: map[string][]string{
			m.ServiceName: m.AddressList,
		},
	}
	r.start()
	return r, nil
}

func (m *MyResolverBuilder) Scheme() string {
	return m.SchemeName
}
