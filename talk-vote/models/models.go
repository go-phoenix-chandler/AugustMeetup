package models

import (
	"sync"
)

// Talks holds all registered talks
type Talks struct {
	Lock  sync.Locker `json:"-"`
	Talks []Talk      `json:"talks"`
}

// Talk holds the relevant data about a single talk
type Talk struct {
	ID        int    `json:"id"`
	Presenter string `json:"presenter"`
	Topic     string `json:"topic"`
	Votes     int    `json:"votes"`
}
