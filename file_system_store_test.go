package game

import (
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, `[
			{"Name": "Jon", "Wins": 20},
			{"Name": "Doe", "Wins": 21}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		AssertNoError(t, err)

		got := store.GetLeague()

		want := []Player{
			{"Doe", 21},
			{"Jon", 20},
		}

		AssertLeague(t, got, want)

		// Read Again
		got = store.GetLeague()
		AssertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, `[
			{"Name": "Jon", "Wins": 20},
			{"Name": "Doe", "Wins": 21}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		AssertNoError(t, err)

		got := store.GetPlayerScore("Doe")
		want := 21
		AssertScoreEquals(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, `[
		{"Name": "Jon", "Wins": 20},
		{"Name": "Doe", "Wins": 21}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		AssertNoError(t, err)

		store.RecordWin("Jon")

		got := store.GetPlayerScore("Jon")
		want := 21
		AssertScoreEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, `[
			{"Name": "Jon", "Wins": 20},
			{"Name": "Doe", "Wins": 21}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		AssertNoError(t, err)

		store.RecordWin("Snow")

		got := store.GetPlayerScore("Snow")
		want := 1
		AssertScoreEquals(t, got, want)
	})

	t.Run("return sorted league", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, `[
			{"Name": "Jon", "Wins": 20},
			{"Name": "Doe", "Wins": 21}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)

		AssertNoError(t, err)

		got := store.GetLeague()

		want := League{
			{"Doe", 21},
			{"Jon", 20},
		}

		AssertLeague(t, got, want)

		// read again
		got = store.GetLeague()
		AssertLeague(t, got, want)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, "")
		defer cleanDatabase()

		_, err := NewFileSystemPlayerStore(database)
		AssertNoError(t, err)
	})
}
