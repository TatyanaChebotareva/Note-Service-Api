package main

import (
	"context"
	"fmt"
	"log"

	desc "github.com/TatyanaChebotareva/Note-Service-Api/pkg/note_v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

const address = "localhost:50051"

func main() {
	con, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("didn't connect: %s", err.Error())
	}
	defer con.Close()

	client := desc.NewNoteV1Client(con)

	// createNote(client)
	// getNote(client)
	getListNote(client)
	// updateNote(client)
	// deleteNote(client)
}

func createNote(client desc.NoteV1Client) {
	note := desc.Note{
		Title:  "Final",
		Text:   "",
		Author: "Tanya",
	}

	createRes, err := client.Create(context.Background(), &desc.CreateRequest{
		Note: &note,
	})

	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println("Id:", createRes.GetId())
}

func getNote(client desc.NoteV1Client) {
	getRes, err := client.Get(context.Background(), &desc.GetRequest{
		Id: 3,
	})

	if err != nil {
		log.Println(err.Error())
	}

	fmt.Printf("Title: %s\nText: %s\nAuthor: %s\nCreated at: %s\n",
		getRes.Note.GetTitle(), getRes.Note.GetText(), getRes.Note.GetAuthor(),
		getRes.Timestamp.GetCreatedAt().AsTime())

	if getRes.Timestamp.GetUpdatedAt().IsValid() {
		fmt.Printf("Updated at: %s\n", getRes.Timestamp.GetUpdatedAt().AsTime())
	}
}

func getListNote(client desc.NoteV1Client) {
	getListRes, err := client.GetList(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Println(err.Error())
	}

	for i := 0; i < len(getListRes.NoteList); i++ {
		fmt.Printf("Title: %s\nText: %s\nAuthor: %s\nCreated at: %s\n",
			getListRes.NoteList[i].GetTitle(), getListRes.NoteList[i].GetText(), getListRes.NoteList[i].GetAuthor(),
			getListRes.TimestampList[i].GetCreatedAt().AsTime())
		if getListRes.TimestampList[i].GetUpdatedAt().IsValid() {
			fmt.Printf("Updated at: %s\n", getListRes.TimestampList[i].GetUpdatedAt().AsTime())
		}
		fmt.Println()
	}
}

func updateNote(client desc.NoteV1Client) {
	note := desc.Note{
		Title:  "Second note",
		Text:   "Testing update",
		Author: "Tatyana",
	}

	_, err := client.Update(context.Background(), &desc.UpdateRequest{
		Id:   2,
		Note: &note,
	})

	if err != nil {
		log.Println(err.Error())
	}
}

func deleteNote(client desc.NoteV1Client) {

	_, err := client.Delete(context.Background(), &desc.DeleteRequest{
		Id: 2,
	})

	if err != nil {
		log.Println(err.Error())
	}
}
