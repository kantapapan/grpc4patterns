package main

import (
	"context"
	"io"
	"log"

	"github.com/pkg/errors"

	pb "grpc4patterns/srvstreaming/pb/notification"

	"google.golang.org/grpc"
)

func request(client pb.NotificationClient, num int32) error {
	req := &pb.NotificationRequest{
		Num: num,
	}
	stream, err := client.Notification(context.Background(), req)
	if err != nil {
		return errors.Wrap(err, "streamエラー")
	}
	for {
		reply, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Println("これ：", reply.GetMessage())
	}
	return nil
}

func exec(num int32) error {
	address := "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return errors.Wrap(err, "コネクションエラー")
	}
	defer conn.Close()
	client := pb.NewNotificationClient(conn)
	return request(client, num)
}

func main() {
	num := int32(5)
	if err := exec(num); err != nil {
		log.Println(err)
	}
}
