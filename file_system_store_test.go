package main

import (
	"io"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Jon", "Wins": 20},
			{"Name": "Doe", "Wins": 21}]`)
		defer cleanDatabase()

		store := FileSystemPlayerStore{database}

		got := store.GetLeague()

		want := []Player{
			{"Jon", 20},
			{"Doe", 21},
		}

		assertLeague(t, got, want)

		// Read Again
		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Jon", "Wins": 20},
			{"Name": "Doe", "Wins": 21}]`)
		defer cleanDatabase()

		store := FileSystemPlayerStore{database}

		got := store.GetPlayerScore("Doe")
		want := 21
		assertScoreEquals(t, got, want)
	})
}

func createTempFile(t testing.TB, initialData string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	tmpfile, err := os.CreateTemp("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}

func assertScoreEquals(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
