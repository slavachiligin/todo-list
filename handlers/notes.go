package handlers

import (
	"context"
	"fmt"
	"net/http"
	"todo-list/database"
	"todo-list/service"
)

func SetupRoutes(database *database.PostgresDatabase, context context.Context) {
	noteService := service.NewNoteService(database)

	http.HandleFunc("/api/all_notes", func(writer http.ResponseWriter, request *http.Request) {
		NoteHandler(writer, request, noteService, context)
	})

	/*http.HandleFunc("/note", func(writer http.ResponseWriter, request *http.Request) {
		NoteHandler(writer, request, noteService)
	})
	http.HandleFunc("/api/notes", CreateNoteHandler)*/

}

func NoteHandler(w http.ResponseWriter, r *http.Request, noteService *service.NoteService, ctx context.Context) {
	notes, err := noteService.GetAllNotes(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintln(w, "Note:", notes)
}

/*func CreateNoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var note entity.Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		http.Error(w, "Ошибка в данных запроса", http.StatusBadRequest)
		return
	}

	err = CreateNote(note.Title, note.Description)
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
}*/
