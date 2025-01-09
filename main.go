package main

import (
	"context"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"todo-list/database"
	"todo-list/handlers"
)

// Config Структура конфигурации
type Config struct {
	DatabaseURL string `env:"DATABASE_URL" env-required:"true"`
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

	connectDB, err := database.InitAndConnectDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Ошибка инициализации базы данных: %v", err)
		return
	}

	handlers.SetupRoutes(connectDB, context.Background())

	/*ns := service.NewNoteService(connectDB)
	notes, err := ns.GetAllNotes(context.Background())
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}

	fmt.Println(notes)*/

	// Запуск сервера
	fmt.Println("Сервер запущен на :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
