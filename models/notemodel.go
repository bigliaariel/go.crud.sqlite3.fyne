package models

import (
	"fmt"
	"notes-mod/config"
	"notes-mod/entities"
)

type NoteModel struct {
}

func (*NoteModel) FindAll() ([]entities.Note, error) {
	db, err := config.GetDB()

	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("select title,body,color from note;")

		if err2 != nil {
			return nil, err2
		} else {
			var notes []entities.Note
			for rows.Next() {
				var note entities.Note
				rows.Scan(&note.Title, &note.Body, &note.Color)
				notes = append(notes, note)

			}
			rows.Close()
			db.Close()
			return notes, nil
		}

	}

}

func (*NoteModel) Add(title string, note string, color int) (err error) {
	db, err := config.GetDB()

	if err != nil {
		return err
	} else {
		s := `INSERT INTO note  VALUES (null,'%s','%s',%d)`
		q := fmt.Sprintf(s, title, note, color)
		_, err2 := db.Exec(q)
		defer db.Close()
		if err2 != nil {
			return err
		}
	}
	return nil
}

func (*NoteModel) Remove(id int) (err error) {
	db, err := config.GetDB()

	if err != nil {
		return err
	} else {
		s := `DELETE FROM note WHERE id=%d`
		q := fmt.Sprintf(s, id)
		_, err2 := db.Exec(q)
		defer db.Close()
		if err2 != nil {
			return err
		}
	}
	return nil
}

func (*NoteModel) Update(id int, title string, note string, color int) (err error) {
	db, err := config.GetDB()

	if err != nil {
		return err
	} else {
		s := `UPDATE note  SET title='%s',body='%s',color=%d WHERE id=%d`
		q := fmt.Sprintf(s, title, note, color, id)
		_, err2 := db.Exec(q)
		defer db.Close()
		if err2 != nil {
			return err
		}
	}
	return nil
}
