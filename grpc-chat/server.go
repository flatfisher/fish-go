package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"

	pb "./protos"
)

type Server struct {
	host string
}

func NewServer(host string) *Server {
	return &Server{
		host: host,
	}
}

func (s *Server) Run(ctx context.Context) error {
	log.Printf("server: starting on %s", s.host)

	srv := grpc.NewServer()
	pb.RegisterChatServer(srv, s)

	lis, err := net.Listen("tcp", s.host)
	if err != nil {
		return err
	}

	errCh := make(chan error, 1)
	go func() {
		errCh <- srv.Serve(lis)
	}()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		return errors.New("server: canceled")
	}
}

func (s *Server) SendMessage(stream pb.Chat_SendMessageServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		now := time.Now().In(defaultLocation)
		ts, err := ptypes.TimestampProto(now)
		if err != nil {
			return err
		}
		msg := fmt.Sprintf("%8s | %s", req.Username, req.Message)
		err = stream.Send(&pb.SendMessageResponse{
			Message:   msg,
			Timestamp: ts,
		})
		if err != nil {
			return err
		}
		fmt.Fprintf(os.Stderr, "%s %s\n", now.Format(dateFormat), msg)
	}
}
