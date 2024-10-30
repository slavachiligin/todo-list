package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"log"
	"net/http"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jackc/pgx/v5/pgxpool"
	"todo-list/handlers"
)

// Config Структура конфигурации
type Config struct {
	DatabaseURL string `env:"DATABASE_URL" env-required:"true"`
}

type Note struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var db *pgx.Conn

// Создание переменной для хранения пула соединений
var dbPool *pgxpool.Pool

// Инициализация подключения к базе данных
func initDB(cfg Config) error {
	// Подключение к базе данных
	var err error
	dbPool, err = pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		return fmt.Errorf("не удалось подключиться к базе данных: %v", err)
	}

	// Тестовое подключение
	err = dbPool.Ping(context.Background())
	if err != nil {
		return fmt.Errorf("не удалось подключиться к базе данных: %v", err)
	}

	fmt.Println("Подключение к базе данных установлено")
	return nil
}

func main() {
	// Явная загрузка .env файла перед инициализацией cleanenv
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %v", err)
	}

	// Загружаем конфигурацию из .env
	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	// Инициализация базы данных
	if err := initDB(cfg); err != nil {
		log.Fatalf("Ошибка инициализации базы данных: %v", err)
	}
	defer dbPool.Close()

	var err error
	db, err = pgx.Connect(context.Background(), "postgres://slavachiligin@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatalf("Ошибка инициализации базы данных: %v\n", err)
	}
	defer db.Close(context.Background())

	// Роутинг
	http.HandleFunc("/note", handlers.NoteHandler)
	http.HandleFunc("/api/notes", handlers.CreateNoteHandler)

	// Запуск сервера
	fmt.Println("Сервер запущен на :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
