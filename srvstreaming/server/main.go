package main

import (
	"fmt"
	"log"
	"net"
	"time"

	pb "grpc4patterns/srvstreaming/pb/notification"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

const port = ":50051"

// ServerServerSide is server
type ServerServerSide struct {
	pb.UnimplementedNotificationServer
}

// Notification is
func (s *ServerServerSide) Notification(req *pb.NotificationRequest, stream pb.Notification_NotificationServer) error {
	fmt.Println("リクエスト受け取った")
	for i := int32(0); i < req.GetNum(); i++ {
		message := fmt.Sprintf("%d", i)
		if err := stream.Send(&pb.NotificationReply{
			Message: message,
		}); err != nil {
			return err
		}
		time.Sleep(time.Second * 1)
	}
	return nil
}

func set() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return errors.Wrap(err, "ポート失敗")
	}
	s := grpc.NewServer()
	var server ServerServerSide
	pb.RegisterNotificationServer(s, &server)
	if err := s.Serve(lis); err != nil {
		return errors.Wrap(err, "サーバ起動失敗")
	}
	return nil
}

func main() {
	fmt.Println("起動")
	if err := set(); err != nil {
		log.Fatalf("%v", err)
	}
}
