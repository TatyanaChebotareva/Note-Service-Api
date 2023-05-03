package main

import (
	"context"
	"fmt"
	desc "github.com/TatyanaChebotareva/Note-Service-Api/pkg/note_v1"
	"google.golang.org/grpc"
	"log"
)

const address = "localhost:50051"

func main() {
	con, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("didn't connect: %s", err.Error())
	}
	defer con.Close()

	client := desc.NewNoteV1Client(con)
	res, err := client.CreateNote(context.Background(), &desc.CreateNoteRequest{
		Title:  "Wow!",
		Text:   "I'm stucked",
		Author: "Tanya",
	})

	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println("Id:", res.Id)
}
