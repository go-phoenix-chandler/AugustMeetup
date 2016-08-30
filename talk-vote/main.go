package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	gklog "github.com/go-kit/kit/log"
	"github.com/go-phoenix-chandler/AugustMeetup/talk-vote/bindings"
	"github.com/go-phoenix-chandler/AugustMeetup/talk-vote/database"
	"golang.org/x/net/context"
)

// interrupt stops execution
func interrupt() error {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	return fmt.Errorf("%s", <-c)
}

func main() {
	var logger gklog.Logger

	logger = gklog.NewJSONLogger(os.Stdout)
	logger = gklog.NewContext(logger).With("ts", gklog.DefaultTimestampUTC)
	log.SetOutput(gklog.NewStdlibAdapter(logger))
	log.SetFlags(0)

	ctx := context.Background()
	errChan := make(chan error)

	go func() {
		errChan <- interrupt()
	}()

	db, err := database.NewDatabase()
	if err != nil {
		errChan <- err
	}
	db.Build()

	bindings.StartApplicationHTTPListener(logger, ctx, db, errChan)

	logger.Log("fatal", <-errChan)
}
