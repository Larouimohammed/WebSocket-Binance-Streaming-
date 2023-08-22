package server

import (
	"context"
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

func (p *Server) Sendmsg(ctx context.Context, Msg *Message.Msg) (*Message.Resp, error) {
	    x := <-binance.Oput
	    out, _ := json.Marshal(&x)
        return &Message.Resp{Resp: string(out)}, nil
}

func (p *Server) Run() error {

	lis, err := net.Listen("tcp", ":3333")
	if err != nil {
		p.logger.Errorf("failed to listen: %v", err)
		return err
	}
	s := grpc.NewServer()
	Message.RegisterSendServer(s, p)
	p.logger.Infof("server listening at %v", lis.Addr())
	p.logger.Infof("Starting striming to khabchi ")
	
	if err := s.Serve(lis); err != nil {
		p.logger.Errorf("failed to serve: %v", err)
		return err
	}
	
	return nil
}
