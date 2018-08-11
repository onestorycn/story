package main

import (
	"context"
	"fmt"
	grpc2 "google.golang.org/grpc"
	"github.com/processout/grpc-go-pool"
	"time"
	story "story/grpc/proto/posts"
	"flag"
)

func main() {

	p, errPoll := grpcpool.New(func() (*grpc2.ClientConn, error) {
		return grpc2.Dial("127.0.0.1:9994", grpc2.WithInsecure())
	}, 1, 100, time.Second)

	if errPoll != nil {
		fmt.Println(errPoll)
	}

	client, err := p.Get(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	cl := story.NewPostsServiceClient(client.ClientConn)
	//Info := &story.FullPost{StoryId:0, PassId:"15315342405009", Header:"tom", Ext:"", Rel:"1111", Content:"1111111111111111111"}
	//rsp, err := cl.InsertAccountInfo(context.TODO(), &account.Request{Info:Info})
	//rsp, err := cl.InsertPostsInfo(context.TODO(), &story.Request{Info: Info})
	rsp, err := cl.GetPostsList(context.TODO(), &story.RequestQuery{PassId:"15315342405009", EndTime:1531539456, Limit:1, IsDesc:false, Page:2})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("code : %v \n " , rsp.Code)
	fmt.Printf("message : %v \n " , rsp.Message)
	fmt.Printf("data : %v \n " , rsp.List)
}

