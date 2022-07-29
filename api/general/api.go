package general

import (
	"github.com/NpoolPlatform/message/npool/inspire/mgr/v1/archivement/general"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	general.UnimplementedGeneralMgrServer
}

func Register(server grpc.ServiceRegistrar) {
	general.RegisterGeneralMgrServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
