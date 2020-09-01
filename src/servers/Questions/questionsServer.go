package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	questionspb "github.com/dhivakarj/Community/src/proto/Questions"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	objectid "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	options "go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

type server struct{}

var collection *mongo.Collection

func main() {

	//connect to Mongo db
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	// mongo client failure through error
	if err != nil {
		log.Fatal(err)
		fmt.Printf("MongoDB Client connect failure %v\n", err)
	}
	// connect to Mongo db
	err = client.Connect(context.TODO())
	// mongo connect failure
	if err != nil {
		log.Fatal(err)
		fmt.Printf("MongoDB connect failure %v\n", err)
	}
	// create a database or connect to database and its collection
	// collection here is a database
	collection = client.Database("community").Collection("questions")
	fmt.Printf("MongoDB Connected\n")

	// server is starting
	fmt.Printf("Server is starting....\n")

	lis, err := net.Listen("tcp", "0.0.0.0:50053")
	if err != nil {
		log.Fatalf("Server failed to connect %v", err)
		fmt.Printf("Error when Server start.... %v\n", err)
	}

	// Register the server to the struct
	s := grpc.NewServer()
	questionspb.RegisterQuestionServiceServer(s, &server{})
	// error when connecting
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Server is not serving %v", err)
		fmt.Printf("Server is not serving.... %v\n", err)
	}

	fmt.Printf("Server is ready to start serving clients\n")

	// wait for the control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// wait
	<-ch

	// Stop the service - and will not server clients
	s.Stop()
	// close the server
	lis.Close()
	// Disconnect from the database
	client.Disconnect(context.TODO())
}

// create a replica of questions
type questionItem struct {
	ID          objectid.ObjectID `bson:"_id,omitempty"`
	Description string            `bson:"qdesc"`
	Type        string            `bson:"qtype"`
	Valid       string            `bson:"valid"`
}

// implement GetQuestions
func (s *server) GetQuestions(ctx context.Context, req *questionspb.GetQuestionsRequest) (*questionspb.GetQuestionsResponse, error) {

	cur, err := collection.Find(context.Background(), bson.D{})
	fmt.Println("getting... ")
	if err != nil {
		log.Fatalf("error when trying to get all questions %v", err)
		fmt.Printf("error when trying to get all questions %v", err)
		return nil, err
	}
	fmt.Println("decoding... ")

	//defer cur.Close(context.Background())

	quesList := []*questionspb.Question{}
	for cur.Next(context.Background()) {
		data := &questionItem{}
		err := cur.Decode(data)
		if err != nil {
			log.Fatalf("error when decoding data %v", err)
			fmt.Printf("error when decoding data %v", err)
			return nil, err
		}
		quesList = append(quesList, &questionspb.Question{
			QuestionId:    data.ID.Hex(),
			QuestionDesc:  data.Description,
			QuestionType:  data.Type,
			QuestionValid: data.Valid,
		})

	}
	resp := &questionspb.GetQuestionsResponse{
		QuestionSlice: quesList}

	return resp, nil
}

// implement createQuestions
func (s *server) CreateQuestions(ctx context.Context, req *questionspb.CreateQuestionsRequest) (*questionspb.CreateQuestionsResponse, error) {
	ques := req.GetCQuestion()
	qi := questionItem{
		Description: ques.GetQuestionDesc(),
		Type:        ques.GetQuestionType(),
		Valid:       ques.GetQuestionValid(),
	}

	result, err := collection.InsertOne(context.Background(), qi)
	if err != nil {
		log.Fatalf("error when trying to insert into %v", err)
		fmt.Printf("error when trying to insert into %v", err)
		return nil, err
	}
	oid, ok := result.InsertedID.(objectid.ObjectID)
	if !ok {
		log.Fatalf("Object id not returned %v", err)
		fmt.Printf("Object id not returned %v", err)
		return nil, err
	}
	resp := &questionspb.CreateQuestionsResponse{
		CQuestion: &questionspb.Question{
			QuestionId:    oid.Hex(),
			QuestionDesc:  ques.GetQuestionDesc(),
			QuestionType:  ques.GetQuestionType(),
			QuestionValid: ques.GetQuestionValid(),
		},
	}
	return resp, nil
}

// implement Update questions
func (s *server) UpdateQuestions(ctx context.Context, res *questionspb.UpdateQuestionsRequest) (*questionspb.UpdateQuestionsResponse, error) {

	objID, err := primitive.ObjectIDFromHex(res.GetCQuestion().GetQuestionId())
	if err != nil {
		log.Fatalf("error converting objectid %v", err)
		fmt.Printf("error converting objectid %v", err)
		return nil, err
	}

	filter := bson.M{"_id": bson.M{"$eq": objID}}
	update := bson.M{"$set": bson.M{
		"qdesc":  res.GetCQuestion().GetQuestionDesc(),
		"qtype":  res.GetCQuestion().GetQuestionType(),
		"qvalid": res.GetCQuestion().GetQuestionValid(),
	},
	}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatalf("update error %v", err)
		fmt.Printf("update error %v", err)
		return nil, err
	}

	if result.ModifiedCount == 0 {
		log.Fatalf("update failed %v", err)
		fmt.Printf("update failed %v", err)
		return nil, err
	}

	resp := questionspb.UpdateQuestionsResponse{
		CQuestion: &questionspb.Question{
			QuestionId:    objID.Hex(),
			QuestionDesc:  res.GetCQuestion().GetQuestionDesc(),
			QuestionType:  res.GetCQuestion().GetQuestionType(),
			QuestionValid: res.GetCQuestion().GetQuestionValid(),
		},
	}

	return &resp, nil
}
