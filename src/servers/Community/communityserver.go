package main

import (
	"context"
	"fmt"
	"log"
	"net"

	communitypb "github.com/dhivakarj/Community/src/proto/Community"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("Server is running....")

	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatalf("Server failed to connect %v", err)
	}

	s := grpc.NewServer()
	communitypb.RegisterCommunityServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Server is not serving %v", err)
	}
}

func (s *server) AddMembers(ctx context.Context, req *communitypb.AddMembersRequest) (*communitypb.AddMembersResponse, error) {
	res := &communitypb.AddMembersResponse{
		MemRes: &communitypb.Member{
			Id:             req.MemReq.GetId(),
			FirstName:      req.MemReq.GetFirstName(),
			LastName:       req.MemReq.GetLastName(),
			CognizantTitle: req.MemReq.GetCognizantTitle(),
			CdeTitle:       req.MemReq.GetCdeTitle(),
		},
		Result: "Successful",
	}
	return res, nil
}

func (s *server) ShowMember(ctx context.Context, req *communitypb.ShowMemberRequest) (*communitypb.ShowMemberResponse, error) {
	res := &communitypb.ShowMemberResponse{
		MemRes: &communitypb.Member{
			Id:             req.MemReq.GetId(),
			FirstName:      req.MemReq.GetFirstName(),
			LastName:       req.MemReq.GetLastName(),
			CognizantTitle: req.MemReq.GetCognizantTitle(),
			CdeTitle:       req.MemReq.GetCdeTitle(),
		},
		Result: "Successful",
	}
	return res, nil
}

func (s *server) GetMembers(ctx context.Context, req *communitypb.GetMembersRequest) (*communitypb.GetMembersResponse, error) {
	res := &communitypb.GetMembersResponse{
		MemRes: &communitypb.Member{
			Id:             req.MemReq.GetId(),
			FirstName:      req.MemReq.GetFirstName(),
			LastName:       req.MemReq.GetLastName(),
			CognizantTitle: req.MemReq.GetCognizantTitle(),
			CdeTitle:       req.MemReq.GetCdeTitle(),
		},
		Result: "Successful",
	}
	return res, nil
}

func (s *server) DeleteMember(ctx context.Context, req *communitypb.DeleteMemberRequest) (*communitypb.DeleteMemberResponse, error) {
	res := &communitypb.DeleteMemberResponse{
		Result: "Successful",
	}
	return res, nil
}
