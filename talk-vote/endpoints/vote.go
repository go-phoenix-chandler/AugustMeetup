package endpoints

import (
	"github.com/go-phoenix-chandler/AugustMeetup/talk-vote/database"
	"golang.org/x/net/context"
)

// VoteServicer is the service used throughout the application
type VoteServicer interface {
	Vote(context.Context, interface{}) (interface{}, error)
}

// VoteService holds the state for the vote service
type VoteService struct {
	db *database.Database
}

// NewVoteService sets up a new Vote service
func NewVoteService(db *database.Database) VoteServicer {
	return &VoteService{db}
}

// VoteServiceRequest is the data in the POST request
type VoteServiceRequest struct {
	TalkID int `json:"talk_id"`
}

// VoteServiceResponse is the data returned to the caller of the Vote endpoint
type VoteServiceResponse struct {
	TalkID int    `json:"talk_id"`
	Status string `json:"status"`
	Err    string `json:"err"`
}

// Vote is the endpoint to submit a vote to
func (v *VoteService) Vote(ctx context.Context, request interface{}) (interface{}, error) {
	req, ok := request.(VoteServiceRequest)
	if !ok {
		return &VoteServiceResponse{Err: errUnableToAssertRequestType.Error()}, errUnableToAssertRequestType
	}
	if err := v.db.Vote(req.TalkID); err != nil {
		return nil, err
	}
	return &VoteServiceResponse{TalkID: req.TalkID, Status: "accepted"}, nil
}
