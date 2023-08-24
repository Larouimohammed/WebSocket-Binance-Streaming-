package server

import (
	"encoding/json"
	"hh/binance"
	Message "hh/proto"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const GRPC_SERVER = ("grpc server")

type Server struct {
	name   string
	logger *logrus.Entry
	Message.UnimplementedSendServer
}

func New() *Server {
	return &Server{name: GRPC_SERVER,
		logger: logrus.WithField("logger", GRPC_SERVER)}
}

func (s *Server) Sendmsg(Msg *Message.Msg, stream Message.Send_SendmsgServer) error {
	for {
		x := <-binance.Oput
		out, _ := json.Marshal(&x)
		if err := stream.Send(&Message.Resp{Resp: string(out)}); err != nil {
			return err
		}

		
	}

}

func (p *Server) Run() error {

	lis, err := net.Listen("tcp", ":3333")
	if err != nil {
		p.logger.Errorf("failed to listen: %v", err)
		return err
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	//s := grpc.NewServer()
	Message.RegisterSendServer(grpcServer, p)
	p.logger.Infof("server listening at %v", lis.Addr())
	p.logger.Infof("Starting striming to khabchi ")

	if err := grpcServer.Serve(lis); err != nil {
		p.logger.Errorf("failed to serve: %v", err)
		return err
	}

	return nil
}
