package service

import (
	"context"
	"todo-list/database"
	"todo-list/entity"
)

type NoteService struct {
	database *database.PostgresDatabase
}

func NewNoteService(database *database.PostgresDatabase) *NoteService {
	return &NoteService{database: database}
}

func (noteService *NoteService) GetAllNotes(ctx context.Context) ([]entity.Note, error) {
	return noteService.database.GetAllNotes(ctx)
}
