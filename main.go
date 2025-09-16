package main

import (
	"log"
	"net/http"

	"todoapp/api"
	"todoapp/database"
	"todoapp/database/generated"
	"todoapp/settings"

	"github.com/gorilla/mux"
)

func main() {
	dsn := database.DSN(settings.PG_HOST, settings.PG_PORT, settings.PG_USER, settings.PG_PASS, settings.PG_DB, settings.PG_SSLMODE)
	if err := database.Init(dsn); err != nil {
		log.Fatalf("db init: %v", err)
	}

	q := generated.New(database.DB)

	r := mux.NewRouter()

	// API маршруты
	api.RegisterRoutes(r, q)

	// Статические файлы для всех остальных маршрутов
	fs := http.FileServer(http.Dir("./static"))
	r.PathPrefix("/").Handler(fs)

	log.Printf("listening on :%s", settings.APP_PORT)
	log.Fatal(http.ListenAndServe(":"+settings.APP_PORT, r))

}
