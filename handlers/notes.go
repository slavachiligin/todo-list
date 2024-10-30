package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Функция-обработчик для маршрута /note
func NoteHandler(w http.ResponseWriter, r *http.Request) {
	// Выполнение запроса к базе данных
	var noteContent string
	err := dbPool.QueryRow(context.Background(), "SELECT content FROM notes WHERE id=$1", 1).Scan(&noteContent)
	if err != nil {
		http.Error(w, "Ошибка получения заметки", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Note:", noteContent)
}

func CreateNoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var note Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		http.Error(w, "Ошибка в данных запроса", http.StatusBadRequest)
		return
	}

	err = createNote(note.Title, note.Content)
	if err != nil {
		http.Error(w, "Ошибка при добавлении заметки в базу данных", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Заметка успешно создана"})
}

func CreateNote(title, content string) error {
	_, err := db.Exec(context.Background(), "INSERT INTO notes (title, content) VALUES ($1, $2)", title, content)
	return err
}
