package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Config Структура конфигурации
type Config struct {
	DatabaseURL string `env:"DATABASE_URL" env-required:"true"`
}

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

// Функция-обработчик для маршрута /note
func noteHandler(w http.ResponseWriter, r *http.Request) {
	// Выполнение запроса к базе данных
	var noteContent string
	err := dbPool.QueryRow(context.Background(), "SELECT content FROM notes WHERE id=$1", 1).Scan(&noteContent)
	if err != nil {
		http.Error(w, "Ошибка получения заметки", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Note:", noteContent)
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

	// Роутинг
	http.HandleFunc("/note", noteHandler)

	// Запуск сервера
	fmt.Println("Сервер запущен на :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
