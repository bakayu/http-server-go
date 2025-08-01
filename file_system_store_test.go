package game

import (
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Jon", "Wins": 20},
			{"Name": "Doe", "Wins": 21}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)

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

		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		got := store.GetPlayerScore("Doe")
		want := 21
		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
		{"Name": "Jon", "Wins": 20},
		{"Name": "Doe", "Wins": 21}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		store.RecordWin("Jon")

		got := store.GetPlayerScore("Jon")
		want := 21
		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Jon", "Wins": 20},
			{"Name": "Doe", "Wins": 21}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		store.RecordWin("Snow")

		got := store.GetPlayerScore("Snow")
		want := 1
		assertScoreEquals(t, got, want)
	})

	t.Run("return sorted league", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Jon", "Wins": 20},
			{"Name": "Doe", "Wins": 21}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)

		assertNoError(t, err)

		got := store.GetLeague()

		want := League{
			{"Doe", 21},
			{"Jon", 20},
		}

		assertLeague(t, got, want)

		// read again
		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, "")
		defer cleanDatabase()

		_, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)
	})
}

func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
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

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}
