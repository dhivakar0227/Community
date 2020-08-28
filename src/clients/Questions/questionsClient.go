package main

import (
	"context"
	"fmt"
	"log"

	questionspb "github.com/dhivakarj/Community/src/proto/Questions"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Client started ...")
	cc, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	c := questionspb.NewQuestionServiceClient(cc)

	dogetQuestions(c)

}

func dogetQuestions(c questionspb.QuestionServiceClient) {
	res, err := c.GetQuestions(context.Background(), &questionspb.GetQuestionsRequest{})
	if err != nil {
		fmt.Printf("Error occurred during addMember %v - ", err)
		return
	}
	for i, key := range res.GetQuestionSlice() {
		fmt.Printf("Result of addMember %v %+v \n", i, key)
	}

}
