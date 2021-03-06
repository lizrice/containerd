package healthcheck

import (
	"github.com/docker/containerd/plugin"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type Service struct {
	serve *health.Server
}

func init() {
	plugin.Register("healthcheck-grpc", &plugin.Registration{
		Type: plugin.GRPCPlugin,
		Init: NewService,
	})
}

func NewService(ic *plugin.InitContext) (interface{}, error) {
	return &Service{
		health.NewServer(),
	}, nil
}

func (s *Service) Register(server *grpc.Server) error {
	grpc_health_v1.RegisterHealthServer(server, s.serve)
	return nil
}
