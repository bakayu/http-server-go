package main

import (
	"fmt"
	"log"
	"os"

	game "github.com/bakayu/http-server-go"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := game.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")
	game.NewCLI(store, os.Stdin).PlayGame()
}
