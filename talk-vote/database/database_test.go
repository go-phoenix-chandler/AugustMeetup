package database

import (
	"log"
	"testing"
)

func setup() *Database {
	db, err := NewDatabase()
	if err != nil {
		log.Fatal(err)
	}
	db.Build()
	return db
}

func TestNewDatabase(t *testing.T) {
	db, err := NewDatabase()
	if err != nil {
		t.Error(err)
	}
	db.Close()
}

func TestBuild(t *testing.T) {
	db, err := NewDatabase()
	if err != nil {
		t.Error(err)
	}
	if err := db.Build(); err != nil {
		t.Error(err)
	}
	db.Close()
}

func TestTalk(t *testing.T) {
	db := setup()
	_, err := db.Talk(3)
	if err != nil {
		t.Error(err)
	}
	db.Close()
}

func TestTalks(t *testing.T) {
	db := setup()
	_, err := db.Talks()
	if err != nil {
		t.Error(err)
	}
	db.Close()
}

func TestVote(t *testing.T) {
	db := setup()
	if err := db.Vote(4); err != nil {
		t.Error(err)
	}
	db.Close()
}

func TestItob(t *testing.T) {
	_ = itob(9)
}
