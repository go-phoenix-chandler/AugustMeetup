package database

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/go-phoenix-chandler/AugustMeetup/talk-vote/models"
)

const bucket = "talks"

var talks = []models.Talk{
	{
		ID:        1,
		Presenter: "Brian Downs",
		Topic:     "Creating a simple service using Go-Kit",
		Votes:     0,
	},
	{
		ID:        2,
		Presenter: "Brian Downs",
		Topic:     "Testing with GoConvey",
		Votes:     0,
	},
	{
		ID:        3,
		Presenter: "Paul Crofts",
		Topic:     "Containerizing a Go-Kit service",
		Votes:     0,
	},
	{
		ID:        4,
		Presenter: "Josh Baker",
		Topic:     "gJSON",
		Votes:     0,
	},
}

// Database holds the database connection
type Database struct {
	*bolt.DB
}

// NewDatabase provides a new value of type Database pointer which
// contains a connection to the database
func NewDatabase() (*Database, error) {
	db, err := bolt.Open("talks.db", 0600, nil)
	if err != nil {
		return nil, err
	}
	return &Database{db}, nil
}

// Build talks all entries in the talks slice and adds them to the database
func (d *Database) Build() error {
	for _, talk := range talks {
		b, err := json.Marshal(talk)
		if err != nil {
			return err
		}
		err = d.DB.Update(func(tx *bolt.Tx) error {
			bucket, err := tx.CreateBucketIfNotExists([]byte(bucket))
			if err != nil {
				return err
			}
			err = bucket.Put(itob(talk.ID), b)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// Talk will retrieve a talk from the database for the given talk ID
func (d *Database) Talk(talkID int) (*models.Talk, error) {
	var talk models.Talk
	err := d.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucket))
		if bucket == nil {
			return errors.New("bucket not found")
		}
		val := bucket.Get(itob(talkID))
		decoder := json.NewDecoder(strings.NewReader(string(val)))
		if err := decoder.Decode(&talk); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &talk, nil
}

// Talks retrieves all talks in the database
func (d *Database) Talks() ([]models.Talk, error) {
	var talks []models.Talk
	err := d.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucket))
		var talk models.Talk
		if err := bucket.ForEach(func(k, v []byte) error {
			decoder := json.NewDecoder(strings.NewReader(string(v)))
			if err := decoder.Decode(&talk); err != nil {
				return err
			}
			talks = append(talks, talk)
			return nil
		}); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return talks, nil
}

// Vote increments a vote count for the given talk
func (d *Database) Vote(talkID int) error {
	return d.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucket))
		talk, err := d.Talk(talkID)
		if err != nil {
			return err
		}
		talk.Votes++
		buf, err := json.Marshal(talk)
		if err != nil {
			return err
		}
		return bucket.Put(itob(talkID), buf)
	})
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
