package server

import (
	"fmt"
	"net"
	"time"

	"github.com/joaoptgaino/server-streaming-rpc/pb"
	"google.golang.org/grpc"
)

type StatusServer struct {
	pb.UnimplementedStatusServiceServer
}

func (s *StatusServer) StreamStatus(req *pb.StreamRequest, stream grpc.ServerStreamingServer[pb.StreamResponse]) error {
	taskId := req.TaskId
	fmt.Printf("Streaming status for task id: %s \n", taskId)

	for i := 0; i <= 100; i += 10 {
		status := pb.StreamResponse{
			Message:  getMessage(i),
			Progress: int64(i),
		}
		err := stream.Send(&status)

		if err != nil {
			return err
		}

		time.Sleep(1 * time.Second)

	}
	return nil
}

func getMessage(val int) string {
	if val == 100 {
		return "Done"
	}
	return "Progressing"
}

func Run() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterStatusServiceServer(grpcServer, &StatusServer{})

	if err := grpcServer.Serve(listen); err != nil {
		panic(err)
	}
}
