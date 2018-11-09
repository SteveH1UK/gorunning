package mongodb

import (
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"

	"github.com/SteveH1UK/gorunning"
)

// AtheleteModel - internal model for one Athelete
type AtheleteModel struct {
	Id          bson.ObjectId `bson:"_id,omitempty"`
	FriendyURL  string        `bson:"friendly-url"`
	Name        string        `bson:"name"`
	DateOfBirth string        `bson:"date-of-birth"` // dd-MM-yyyy
	CreatedAt   time.Time     `bson:"created-at"`
}

func atheleteModelIndex() mgo.Index {
	return mgo.Index{
		Key:        []string{"friendyURL"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}

func newAtheleteModel(a *root.NewAthelete) *AtheleteModel {
	athelete := AtheleteModel{FriendyURL: a.FriendyURL, Name: a.Name, DateOfBirth: a.DateOfBirth, CreatedAt: time.Now()}
	return &athelete
}
