package main

import (
	"fmt"
	"notes-mod/entities"
	"notes-mod/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type TapPostit struct {
	*widget.Entry //composition

	//function pointers to set to get events
	OnTapped func(id int32)
	Id       int32
}

func NewTapPostit(title string, body string, tappedLeft func(id int32), id int32) *TapPostit {
	txt := (fmt.Sprintf("%s\n\n%s", title, body))
	lbl1 := widget.NewMultiLineEntry()
	lbl1.Text = txt
	lbl1.Enable()
	return &TapPostit{
		lbl1,
		tappedLeft,
		id,
	}
}

func (mc *TapPostit) Tapped(pe *fyne.PointEvent) {
	if mc.OnTapped != nil {
		mc.OnTapped(mc.Id)

	}
}

func main() {

	myApp := app.New()
	w := myApp.NewWindow("Notes")
	w.Resize(fyne.NewSize(800, 600))
	w.SetMaster()
	w.CenterOnScreen()
	w.Show()

	c2 := container.NewGridWrap(fyne.NewSize(150, 150))
	c2.Resize(fyne.NewSize(700, 500))
	c2.Move(fyne.NewPos(50, 50))

	seeAll(c2, w)
	b := widget.NewButton("+", func() {
		openDialog(-1, "ADD", c2, w)
	})

	b.Resize(fyne.NewSize(20, 40))
	b.Move(fyne.NewPos(20, 20))
	c1 := container.NewWithoutLayout(b, c2)
	w.SetContent(c1)
	myApp.Run()

} // end main

// seeAll show all notes
func seeAll(c2 *fyne.Container, w fyne.Window) {

	c2.Objects = nil
	var noteModel models.NoteModel
	notes, err := noteModel.FindAll()
	if err != nil {
		fmt.Println(err)
	}

	for _, note := range notes {
		lbl11 := NewTapPostit(note.Title, note.Body, func(id int32) {
			openDialog(id, "UPDATE", c2, w)
		}, note.Id)
		//fmt.Println(note.Id)
		lbl11.Wrapping = fyne.TextWrapWord
		c2.Add(lbl11)

	}
	c2.Refresh()
}

// openDialog show a dialog inside the window's app
func openDialog(id int32, mode string, c2 *fyne.Container, w fyne.Window) {
	box := container.NewVBox()
	lbl := widget.NewEntry()
	txt := widget.NewMultiLineEntry()
	box.Resize(fyne.NewSize(400, 300))
	box.Refresh()
	txt.MultiLine = true
	box.Add(lbl)
	box.Add(txt)
	if mode == "UPDATE" {
		var m models.NoteModel
		var selected *entities.Note
		selected, _ = m.Find(id)
		lbl.Text = selected.Title
		txt.Text = selected.Body

	}
	fi := widget.NewFormItem("", box)
	var cont []*widget.FormItem
	cont = append(cont, fi)

	if mode == "ADD" {
		dialog.ShowForm("Note", "Add", "Cancel", cont, func(b bool) {
			if b {

				var m models.NoteModel
				m.Add(lbl.Text, txt.Text, 1)
				seeAll(c2, w)
			}
		}, w)
	}

	if mode == "UPDATE" {
		dialog.ShowForm("Note", "Update", "Delete Note", cont, func(b bool) {

			if b {
				var m models.NoteModel
				m.Update(id, lbl.Text, txt.Text, 1)
				seeAll(c2, w)
			} else {
				var m models.NoteModel
				m.Remove(id)
				seeAll(c2, w)
			}
		}, w)
	}
}
