package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/pkg/errors"

	pb "grpc4patterns/clientstreaming/pb/upload"

	"google.golang.org/grpc"
)

func request(client pb.UploadClient) error {
	stream, err := client.Upload(context.Background())
	if err != nil {
		return err
	}
	values := []int32{1, 2, 3, 4, 5}
	for _, value := range values {
		fmt.Println("送る値:", value)
		if err := stream.Send(&pb.UploadRequest{
			Value: value,
		}); err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		time.Sleep(time.Second * 1)
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}
	log.Printf("結果: %v", reply)
	return nil
}

func exec() error {
	address := "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return errors.Wrap(err, "コネクションエラー")
	}
	defer conn.Close()
	client := pb.NewUploadClient(conn)
	return request(client)
}

func main() {
	if err := exec(); err != nil {
		log.Println(err)
	}
}
