package rpc

type RpcServer interface {
	Start() error
	Stop()
}

type Server struct {

}

func (s *Server) Start() error {
	
}

func (s *Server) Stop()  {
	
}