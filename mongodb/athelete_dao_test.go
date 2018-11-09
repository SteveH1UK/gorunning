// +build integration

package mongodb

import (
	"fmt"
	"log"
	"testing"

	"github.com/globalsign/mgo"

	"github.com/SteveH1UK/gorunning"
	"github.com/SteveH1UK/gorunning/config"
)

// TestCreateNewAthelete - test creating new atheletes
func TestCreateNewAthelete(t *testing.T) {
	fmt.Println("Test")

	cfg, err := config.Get()
	if err != nil {
		log.Println("Error getting config:", err)
		panic(1)
	}

	m, err := Init(cfg.TestMongoURL, cfg.TestDbName)
	if err != nil {
		log.Fatalln("unable to connect to mongodb ", err)
		panic(-1)
	}

	//

	// delete all items in collection
	var iSession *mgo.Session
	iSession = m.getInitialSession()
	session := iSession.Copy()
	defer session.Close()

	c := session.DB(cfg.TestDbName).C("atheletes")
	c.DropCollection()

	// create athelete
	athelete := root.NewAthelete{}
	athelete.FriendyURL = "freddy"
	athelete.Name = "Fred Flinstone"
	athelete.DateOfBirth = "1986/03/14"
	dao := NewAtheleDAO(m)
	dao.CreateAthelete(&athelete)

	//check count
	n, errCount := c.Count()
	if n != 1 {
		fmt.Println("Error Debugging Count", n, errCount)
		t.Error("Expected 1 but got %T", n)
	}

}
