package main

import (
	"fmt"
	"notes-mod/models"
)

func main() {
	var noteModel models.NoteModel
	//noteModel.Add("titulo2", "descripcion 3", 123)
	//noteModel.Remove(1)
	noteModel.Update(2, "3", "4", 5)
	notes, err := noteModel.FindAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, note := range notes {
		fmt.Println(note.Title)
		fmt.Println(note.Body)
		fmt.Println(note.Color)
		fmt.Println("---------------")
	}
}
