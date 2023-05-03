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
	createRes, err := client.CreateNote(context.Background(), &desc.CreateNoteRequest{
		Title:  "Wow!",
		Text:   "I'm stucked",
		Author: "Tanya",
	})

	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println("Id:", createRes.Id)

	getRes, err := client.GetNote(context.Background(), &desc.GetNoteRequest{
		Id: 2,
	})

	if err != nil {
		log.Println(err.Error())
	}

	fmt.Printf("Title: %s\nText: %s\nAuthor:%s\n", getRes.Title, getRes.Text, getRes.Author)

}
