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

	//dogetQuestions(c)
	//doCreateQuestions(c)
	doupdateQuestion(c)
}

//create a questions
func doCreateQuestions(c questionspb.QuestionServiceClient) {
	//CreateQuestions(ctx context.Context, in *CreateQuestionsRequest) (*CreateQuestionsResponse, error)
	ques := questionspb.Question{
		QuestionDesc:  "What is your Name?",
		QuestionType:  "Dropdown",
		QuestionValid: "1",
	}

	req := &questionspb.CreateQuestionsRequest{
		CQuestion: &ques,
	}

	resp, err := c.CreateQuestions(context.Background(), req)
	if err != nil {
		fmt.Printf("Error occurred during creating a question %v ", err)
		return
	}
	fmt.Printf("Details of the questions added %+v \n", resp)

}

func dogetQuestions(c questionspb.QuestionServiceClient) {
	res, err := c.GetQuestions(context.Background(), &questionspb.GetQuestionsRequest{})
	if err != nil {
		fmt.Printf("Error occurred during getMember %v", err)
		return
	}
	for i, key := range res.GetQuestionSlice() {
		fmt.Printf("Result of getMember %v %+v \n", i, key)
	}
}

func doupdateQuestion(c questionspb.QuestionServiceClient) {
	//UpdateQuestions(ctx context.Context, in *UpdateQuestionsRequest, opts ...grpc.CallOption) (*UpdateQuestionsResponse, error)
	ques := questionspb.Question{
		QuestionId:    "5f4d714008316b59a91a7149",
		QuestionDesc:  "What is my Name?",
		QuestionType:  "freeflow",
		QuestionValid: "0",
	}

	resp, err := c.UpdateQuestions(context.Background(), &questionspb.UpdateQuestionsRequest{
		CQuestion: &ques,
	})

	if err != nil {
		fmt.Printf("Error occurred during update of question %v ", err)
		return
	}
	fmt.Printf("Successfully updated %+v", resp)
}
