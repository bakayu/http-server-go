package game_test

import (
	"strings"
	"testing"

	game "github.com/bakayu/http-server-go"
)

func TestCLI(t *testing.T) {
	t.Run("Record Jon's win from user input", func(t *testing.T) {
		in := strings.NewReader("Jon wins\n")
		playerStore := &game.StubPlayerStore{}

		cli := game.NewCLI(playerStore, in)
		cli.PlayGame()

		game.AssertPlayerWin(t, playerStore, "Jon")
	})

	t.Run("Record Snow's win from user input", func(t *testing.T) {
		in := strings.NewReader("Snow wins\n")
		playerStore := &game.StubPlayerStore{}

		cli := game.NewCLI(playerStore, in)
		cli.PlayGame()

		game.AssertPlayerWin(t, playerStore, "Snow")
	})
}
