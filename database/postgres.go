package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"todo-list/entity"
)

type PostgresDatabase struct {
	dbPool *pgxpool.Pool
	db     *pgx.Conn
}

func NewPostgresDatabase(dbPool *pgxpool.Pool, db *pgx.Conn) *PostgresDatabase {
	return &PostgresDatabase{dbPool: dbPool, db: db}
}

func InitAndConnectDB(databaseURL string) (*PostgresDatabase, error) {
	var err error
	dbPool, err := pgxpool.New(context.Background(), databaseURL)
	if err != nil {

		return nil, fmt.Errorf("не удалось подключиться к базе данных: %v", err)
	}

	// Тестовое подключение
	err = dbPool.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("не удалось подключиться к базе данных: %v", err)
	}

	fmt.Println("Подключение к базе данных установлено")

	db, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		log.Fatalf("Ошибка инициализации базы данных: %v\n", err)
	}

	return NewPostgresDatabase(dbPool, db), err
}

func (db *PostgresDatabase) CloseDB() {
	if db.dbPool != nil {
		db.dbPool.Close()
	}

	if db.db != nil {
		err := db.db.Close(context.Background())
		if err != nil {
			return
		}
	}
}

func (db *PostgresDatabase) GetAllNotes(ctx context.Context) ([]entity.Note, error) {
	// SQL-запрос для выборки всех заметок
	query := `SELECT id, title, description, created_at, time, done FROM notes`

	rows, err := db.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Массив для хранения заметок
	var notes []entity.Note

	// Обход строк результатов запроса
	for rows.Next() {
		var note entity.Note
		if err := rows.Scan(&note.ID, &note.Title, &note.Description, &note.CreatedAt, &note.Time, &note.Done); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	// Проверка на ошибки после завершения обработки строк
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return notes, nil
}
