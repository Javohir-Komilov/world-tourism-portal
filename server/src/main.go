package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"world-tourism-portal/server/queries"

	_ "github.com/mattn/go-sqlite3"
)

const (
	DRIVER = "sqlite3"
	DBNAME = "world-tourism-portal.db"
)

func main() {
	// Подключение к базе данных
	database, err := sql.Open(DRIVER, DBNAME)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	// Инициализация базы данных через глобальную переменную Q
	queries.Q = queries.NewDB(database)
	if queries.Q == nil {
		log.Fatalf("failed to connect to the database")
	}

	log.Println("Starting World Tourism Portal server...")

	// Инициализация роутера и обработчиков
	r := mux.NewRouter()
	r.HandleFunc("/api/users", getUsersHandler()).Methods("GET")
	r.HandleFunc("/api/hotels", getHotelsHandler()).Methods("GET")
	// Добавьте другие маршруты по мере необходимости

	// Настройка и запуск сервера
	log.Println("Server started on :8080")
	err = startServer(":8080", r)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func startServer(address string, handler *mux.Router) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	defer listener.Close()

	log.Printf("Listening on %s...", address)
	return http.Serve(listener, handler)
}

// Обработчик для получения пользователей
func getUsersHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Получаем пользователей из базы данных
		users, err := queries.Q.GetUsers(r.Context()) // Используем глобальную переменную Q
		if err != nil {
			http.Error(w, "Error fetching users", http.StatusInternalServerError)
			return
		}

		// Отправляем список пользователей в формате JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}

// Обработчик для получения отелей
func getHotelsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Получаем отели из базы данных
		hotels, err := queries.Q.GetHotels(r.Context()) // Используем глобальную переменную Q
		if err != nil {
			http.Error(w, "Error fetching hotels", http.StatusInternalServerError)
			return
		}

		// Отправляем список отелей в формате JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(hotels)
	}
}
