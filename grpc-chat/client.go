package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"

	pb "./protos"
)

type client struct {
	cli      pb.ChatClient
	conn     *grpc.ClientConn
	username string
}

func NewClient(host, username string) (*client, error) {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := &client{
		cli:      pb.NewChatClient(conn),
		conn:     conn,
		username: username,
	}
	return c, nil
}

func (c *client) Run(ctx context.Context) error {
	log.Printf("client: starting")

	stream, err := c.cli.SendMessage(ctx)
	if err != nil {
		return err
	}
	defer func() {
		stream.CloseSend()
		log.Printf("client: disconnected from stream")
	}()
	log.Printf("client: connected to stream")

	errCh := make(chan error)
	go func() {
		errCh <- c.send(ctx, stream, c.username)
	}()
	go func() {
		errCh <- c.receive(ctx, stream)
	}()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		return errors.New("client: canceled")
	}
}

func (c *client) send(ctx context.Context, stream pb.Chat_SendMessageClient, username string) error {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		select {
		case <-ctx.Done():
			return errors.New("client: canceled")
		default:
			if scanner.Scan() {
				err := stream.Send(&pb.SendMessageRequest{
					Username: username,
					Message:  scanner.Text(),
				})
				if err != nil {
					return err
				}
				continue
			}
			return errors.New("failed to scan stdin")
		}
	}
}

func (c *client) receive(ctx context.Context, stream pb.Chat_SendMessageClient) error {
	for {
		fmt.Fprintf(os.Stderr, ">>> ")
		resp, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		t, err := ptypes.Timestamp(resp.Timestamp)
		if err != nil {
			return err
		}
		fmt.Fprintf(os.Stderr, "%s %s\n", t.In(defaultLocation).Format(dateFormat), resp.Message)
	}
}

func (c *client) Close() error {
	return c.conn.Close()
}
