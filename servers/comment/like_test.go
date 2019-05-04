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

func TestCommentServer_Like(t *testing.T) {

	// go initServer(message)
	// <-message
	fmt.Println("客户端开始运行.....")
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

	res, err := client.Like(ctx, &comment.LikeRequest{
		CommentId: 343191254370103296,
	})

	log.Printf("%s: %s", name, res)

	if err != nil {
		t.Fatalf("Like Client err: %+v", err)
	}
	t.Logf("Like Client: %+v", res)
}
