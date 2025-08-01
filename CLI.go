package game

import (
	"bufio"
	"io"
	"strings"
)

// CLI helps players through the game.
type CLI struct {
	playerStore PlayerStore
	in          *bufio.Scanner
}

// NewCLI generates a cli for playing the game.
func NewCLI(store PlayerStore, in io.Reader) *CLI {
	return &CLI{
		playerStore: store,
		in:          bufio.NewScanner(in),
	}
}

// PlayGame starts the game.
func (cli *CLI) PlayGame() {
	userInput := cli.readLine()
	cli.playerStore.RecordWin(extractWinner(userInput))
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}
