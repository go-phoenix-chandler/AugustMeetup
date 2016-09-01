package database

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDatabaseManipulation(t *testing.T) {
	db := setup()
	defer db.Close()
	Convey("Given a new database connection", t, func() {

		Convey("When all talks are queried for", func() {
			talks, err := db.Talks()
			if err != nil {
				t.Error(err)
			}

			Convey("The value should equal 4", func() {
				So(talks, ShouldNotBeNil)
				So(len(talks), ShouldEqual, 4)
			})

		})

		Convey("When one talk is queried for", func() {
			talk, err := db.Talk(3)
			if err != nil {
				t.Error(err)
			}
			Convey("The value should equal 3", func() {
				So(talk.ID, ShouldEqual, 3)
			})

		})

		Convey("When a vote is cast", func() {
			talk, err := db.Talk(4)
			if err != nil {
				t.Error(err)
			}
			votes := talk.Votes
			if err := db.Vote(4); err != nil {
				t.Error(err)
			}
			talk, err = db.Talk(4)
			if err != nil {
				t.Error(err)
			}
			newVotes := talk.Votes
			Convey("The value should increment", func() {
				So(votes, ShouldBeLessThan, newVotes)
			})

		})

	})

}
