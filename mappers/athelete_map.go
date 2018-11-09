package mappers

import (
	"time"

	"github.com/SteveH1UK/gorunning"
	"github.com/SteveH1UK/gorunning/mongodb"
)

// NewAtheleteFromModel - athelete mapper Mongo to external
func NewAtheleteFromModel(am mongodb.AtheleteModel, baseURL string) root.Athelete {
	href := baseURL + am.FriendyURL
	return root.Athelete{HRef: href, FriendyURL: am.FriendyURL, Name: am.Name, DateOfBirth: am.DateOfBirth, CreationDate: mapTime(am.CreatedAt)}
}

func mapTime(time time.Time) string {

	return time.String()
}
