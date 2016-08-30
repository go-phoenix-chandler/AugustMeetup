package endpoints

import (
	"github.com/go-phoenix-chandler/AugustMeetup/talk-vote/database"
	"github.com/go-phoenix-chandler/AugustMeetup/talk-vote/models"
	"golang.org/x/net/context"
)

// TalkServicer
type TalkServicer interface {
	List(context.Context, interface{}) (interface{}, error)
}

// TalkService
type TalkService struct {
	db *database.Database
}

func NewTalkService(db *database.Database) TalkServicer {
	return &TalkService{db}
}

// TalkServiceRequest
type TalkServiceRequest struct{}

// TalkServiceResponse
type TalkServiceResponse struct {
	Talks []models.Talk `json:"talks"`
}

// List will list the talks to be voted on
func (t *TalkService) List(ctx context.Context, r interface{}) (interface{}, error) {
	talks, err := t.db.Talks()
	if err != nil {
		return nil, err
	}
	return &TalkServiceResponse{Talks: talks}, nil
}
