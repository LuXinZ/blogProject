package commentApp

import (
	"context"
	comment "github.com/LuXinZ/blogProject/internal/commentApp/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

type Application struct {
}

func (app *Application) Run() {
	r := gin.Default()
	CommentRouter(r)
}
func InitRpcToComment() (comment.CommentServiceClient, context.Context) {

	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := comment.NewCommentServiceClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return c, ctx
}
func SendRpcToComment(newComment *comment.Comment) {

	_, err = c.CreateComment(ctx, &comment.CreateCommentRequest{Comment: newComment})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
}
