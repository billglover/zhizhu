package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/billglover/zhizhu/database"
	"github.com/billglover/zhizhu/feed"
	"github.com/billglover/zhizhu/storage"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	db, err := database.NewDB("host=localhost port=5432 user=zhizhu password=zhizhu dbname=zhizhu sslmode=disable")
	if err != nil {
		return fmt.Errorf("database set-up: %w", err)
	}
	defer db.Close()

	store := storage.NewStorage(db)
	fm := feed.NewManager(store)
	srv := newServer()
	srv.fm = fm

	err = http.ListenAndServe(":8082", srv)
	return err
}
