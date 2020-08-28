package main

import (
	"context"
	"fmt"
	"log"
	"net"

	questionspb "github.com/dhivakarj/Community/src/proto/Questions"
	"google.golang.org/grpc"
)

type server struct{}

func main() {

	fmt.Println("Server is running....")

	lis, err := net.Listen("tcp", "0.0.0.0:50053")
	if err != nil {
		log.Fatalf("Server failed to connect %v", err)
	}

	s := grpc.NewServer()
	questionspb.RegisterQuestionServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Server is not serving %v", err)
	}
}

func (s *server) GetQuestions(ctx context.Context, req *questionspb.GetQuestionsRequest) (*questionspb.GetQuestionsResponse, error) {
	var qs []*questionspb.Question
	for i := 0; i < 20; i++ {
		q := questionspb.Question{
			QuestionId:    "1",
			QuestionDesc:  "Question 1",
			QuestionType:  "FreeForm",
			QuestionValid: "1",
		}
		qs = append(qs, &q)
	}
	resp := &questionspb.GetQuestionsResponse{
		QuestionSlice: qs,
	}
	return resp, nil
}
