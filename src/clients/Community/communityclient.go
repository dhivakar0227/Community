package main

import (
	"context"
	"fmt"
	"log"

	communitypb "github.com/dhivakarj/Community/src/proto/Community"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Client started ...")
	cc, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	c := communitypb.NewCommunityServiceClient(cc)

	doAddMembers(c)
	doShowMember(c)
	doGetMembers(c)
	doDeleteMembers(c)
}

func doAddMembers(c communitypb.CommunityServiceClient) {
	res, err := c.AddMembers(context.Background(), &communitypb.AddMembersRequest{
		MemReq: &communitypb.Member{
			Id:             "386398",
			FirstName:      "Dhivakar",
			LastName:       "Jeganathan",
			CognizantTitle: "Director",
			CdeTitle:       "Engineering Director",
		},
	})
	if err != nil {
		fmt.Printf("Error occurred during addMember %v - ", err)
		return
	}
	fmt.Printf("Result of addMember %v \n", res.Result)
}

func doShowMember(c communitypb.CommunityServiceClient) {
	res, err := c.ShowMember(context.Background(), &communitypb.ShowMemberRequest{
		MemReq: &communitypb.Member{
			Id:             "386398",
			FirstName:      "Dhivakar",
			LastName:       "Jeganathan",
			CognizantTitle: "Director",
			CdeTitle:       "Engineering Director",
		},
	})
	if err != nil {
		fmt.Printf("Error occurred during showMember %v - ", err)
		return
	}
	fmt.Printf("Result of showMember %v \n", res.Result)
}

func doGetMembers(c communitypb.CommunityServiceClient) {
	res, err := c.GetMembers(context.Background(), &communitypb.GetMembersRequest{
		MemReq: &communitypb.Member{
			Id:             "386398",
			FirstName:      "Dhivakar",
			LastName:       "Jeganathan",
			CognizantTitle: "Director",
			CdeTitle:       "Engineering Director",
		},
	})
	if err != nil {
		fmt.Printf("Error occurred during getMember %v", err)
		return
	}
	fmt.Printf("Result of getMember %v\n", res.Result)
}

func doDeleteMembers(c communitypb.CommunityServiceClient) {
	res, err := c.DeleteMember(context.Background(), &communitypb.DeleteMemberRequest{
		MemReq: &communitypb.Member{
			Id:             "386398",
			FirstName:      "Dhivakar",
			LastName:       "Jeganathan",
			CognizantTitle: "Director",
			CdeTitle:       "Engineering Director",
		},
	})
	if err != nil {
		fmt.Printf("Error occurred during deleteMember %v", err)
		return
	}
	fmt.Printf("Result of deleteMember %v\n", res.Result)
}
