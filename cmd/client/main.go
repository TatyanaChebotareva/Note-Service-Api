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

	createNote(client)
	getNote(client)
	getListNote(client)
	updateNote(client)
	deleteNote(client)
}

func createNote(client desc.NoteV1Client) {
	createRes, err := client.CreateNote(context.Background(), &desc.CreateNoteRequest{
		Title:  "Wow!",
		Text:   "It's working",
		Author: "Tanya",
	})

	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println("Id:", createRes.Id)
}

func getNote(client desc.NoteV1Client) {
	getRes, err := client.GetNote(context.Background(), &desc.GetNoteRequest{
		Id: 2,
	})

	if err != nil {
		log.Println(err.Error())
	}

	fmt.Printf("Title: %s\nText: %s\nAuthor:%s\n", getRes.Title, getRes.Text, getRes.Author)
}

func getListNote(client desc.NoteV1Client) {
	getListRes, err := client.GetListNote(context.Background(), &desc.GetListNoteRequest{})

	if err != nil {
		log.Println(err.Error())
	}

	for _, note := range getListRes.NoteList {
		fmt.Printf("Title: %s\nText: %s\nAuthor:%s\n\n", note.Title, note.Text, note.Author)
	}
}

func updateNote(client desc.NoteV1Client) {
	updateRes, err := client.UpdateNote(context.Background(), &desc.UpdateNoteRequest{
		Id:   2,
		Text: "25.05.2023"})

	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println(updateRes.Result)
}

func deleteNote(client desc.NoteV1Client) {
	deleteRes, err := client.DeleteNote(context.Background(), &desc.DeleteNoteRequest{
		Id: 3,
	})

	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println(deleteRes.Result)
}
