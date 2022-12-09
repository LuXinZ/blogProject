package commentServer

import (
	"context"
	comment "github.com/LuXinZ/blogProject/internal/commentServer/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Application struct {
	Address    string
	RpcAddress string
}
type server struct {
	comment.UnimplementedCommentServiceServer
}

func (s *server) CreateComment(ctx context.Context, request *comment.CreateCommentRequest) (*comment.CreateCommentResponse, error) {
	// sql
	log.Printf("recevied %v", request.Comment.Content)
	newComment := comment.Comment{
		Id:       0,
		AuthorId: 0,
		BlogId:   0,
		Content:  "success",
	}
	return &comment.CreateCommentResponse{Comment: &newComment}, nil
}
func (app *Application) Run() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal("failed to listen")
	}
	// create gRPC
	s := grpc.NewServer()
	// register
	comment.RegisterCommentServiceServer(s, &server{})
	log.Printf("server listen on %v ", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
