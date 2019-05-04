package comment_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"google.golang.org/grpc"

	"github.com/sundogrd/comment-grpc/grpc_gen/comment"
	commentGen "github.com/sundogrd/comment-grpc/grpc_gen/comment"
)

// var message = make(chan bool)

func TestCommentServer_Create(t *testing.T) {

	// go initServer(message)
	// <-message
	fmt.Println("客户端开始运行.....")
	fmt.Println(address)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	defer conn.Close()
	client := commentGen.NewCommentServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.CreateComment(ctx, &comment.CreateCommentRequest{
		AppId: "service test",
		Comment: &comment.CreateCommentRequest_CommentCreateParams{
			TargetId:    23333,
			CreatorId:   32222,
			ParentId:    0,
			ReCommentId: 0,
			Content:     "Test Client Content",
			Extra:       "Test Client Extra",
		},
	})

	log.Printf("%s: %s", name, res)

	if err != nil {
		t.Fatalf("CreateComment Client err: %+v", err)
	}
	t.Logf("CreateComment Client: %+v", res)
}
