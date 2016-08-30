package bindings

import (
	"encoding/json"
	"net/http"

	gklog "github.com/go-kit/kit/log"
	levlog "github.com/go-kit/kit/log/levels"
	gkhttp "github.com/go-kit/kit/transport/http"
	"github.com/go-phoenix-chandler/AugustMeetup/talk-vote/database"
	"github.com/go-phoenix-chandler/AugustMeetup/talk-vote/endpoints"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
)

// StartApplicationHTTPListener creates a Go-routine that has an HTTP listener for the application endpoints
func StartApplicationHTTPListener(logger gklog.Logger, ctx context.Context, db *database.Database, errChan chan error) {
	go func() {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		l := levlog.New(logger)
		l.Info().Log("HTTPAddress", ":8888", "transport", "HTTP/JSON")

		ts := endpoints.NewTalkService(db)
		vs := endpoints.NewVoteService(db)

		router := createApplicationRouter(l, ctx, ts, vs)
		errChan <- http.ListenAndServe(":8888", handlers.RecoveryHandler()(handlers.CombinedLoggingHandler(gklog.NewStdlibAdapter(logger), router)))
	}()
}

// createApplicationRouter sets up the router that will handle all of the application routes
func createApplicationRouter(l levlog.Levels, ctx context.Context, ts endpoints.TalkServicer, vs endpoints.VoteServicer) *mux.Router {
	router := mux.NewRouter()
	router.Handle(
		"/api/v1/talks",
		gkhttp.NewServer(
			ctx,
			ts.List,
			decodeTalkListHTTPRequest,
			encodeTalkListHTTPResponse,
		)).Methods(http.MethodGet)

	router.Handle(
		"/api/v1/vote",
		gkhttp.NewServer(
			ctx,
			vs.Vote,
			decodeVoteHTTPRequest,
			encodeVoteHTTPResponse,
		)).Methods(http.MethodPost)
	return router
}

func decodeTalkListHTTPRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func decodeVoteHTTPRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var vr endpoints.VoteServiceRequest
	if err := json.NewDecoder(r.Body).Decode(&vr); err != nil {
		return nil, err
	}
	return vr, nil
}

func encodeTalkListHTTPResponse(ctx context.Context, w http.ResponseWriter, i interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(i.(*endpoints.TalkServiceResponse))
}

func encodeVoteHTTPResponse(ctx context.Context, w http.ResponseWriter, i interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(i.(*endpoints.VoteServiceResponse))
}
