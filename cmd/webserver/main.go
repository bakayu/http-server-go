package main

import (
	"log"
	"net/http"

	game "github.com/bakayu/http-server-go"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := game.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	server := game.NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen of port 5000 %v", err)
	}
}
