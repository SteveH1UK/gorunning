package mongodb

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"

	"github.com/SteveH1UK/gorunning"
)

// AthleteDAO DAO for Atheletes
type AthleteDAO struct {
	ms *MongoSession
}

// AthleteDAOInterface - interface for DAO
type AthleteDAOInterface interface {
	CreateAthelete(a *root.NewAthelete) error
	FindAtheleteByFriendlyName(friendlyName string) (AtheleteModel, error)
	FindAllAtheletes() ([]AtheleteModel, error)
	EditAthelete(friendlyName string, a *root.NewAthelete) error
}

// NewAtheleDAO create new Athelete type
func NewAtheleDAO(mongoSession *MongoSession) AthleteDAOInterface {
	atheleteDAO := &AthleteDAO{mongoSession}
	// This returns an interface which holds a pointer to the atheleteDAO
	return atheleteDAO
}

// CreateAthelete - adds new athelete
func (dao *AthleteDAO) CreateAthelete(a *root.NewAthelete) error {
	iSession := dao.ms.getInitialSession()
	session := iSession.Copy()
	defer session.Close()
	c := populateAtheleteCollection(session, dao.ms.DbName)
	athelete := newAtheleteModel(a)
	err := c.Insert(&athelete)
	if err != nil {
		if mgo.IsDup(err) {
			err = root.ErrDBRecordExists
		}
	}
	return err
}

// EditAthelete - adds new athelete
func (dao *AthleteDAO) EditAthelete(friendlyName string, a *root.NewAthelete) error {

	_, err := dao.FindAtheleteByFriendlyName(friendlyName)
	if err != nil {
		fmt.Println("Error when finding exiting record ", friendlyName)
		return errors.New("Finding exist record " + err.Error())
	}

	iSession := dao.ms.getInitialSession()
	session := iSession.Copy()
	defer session.Close()
	c := populateAtheleteCollection(session, dao.ms.DbName)
	selector := bson.M{"friendly-url": friendlyName}
	var updator map[string]interface{}
	fmt.Println(" friendly name from [", friendlyName, "] to [", a.FriendyURL+"]")
	if strings.TrimSpace(friendlyName) == strings.TrimSpace(a.FriendyURL) {
		updator = bson.M{"$set": bson.M{"name": a.Name, "date-of-birth": a.DateOfBirth}}
	} else {
		fmt.Println("Updating friendly name from ", friendlyName, " to ", a.FriendyURL)
		updator = bson.M{"$set": bson.M{"friendly-url": a.FriendyURL, "name": a.Name, "date-of-birth": a.DateOfBirth}}
	}
	err = c.Update(selector, updator)

	return err
}

// FindAtheleteByFriendlyName - find atheletes name
func (dao *AthleteDAO) FindAtheleteByFriendlyName(friendlyName string) (AtheleteModel, error) {
	iSession := dao.ms.getInitialSession()
	session := iSession.Copy()
	defer session.Close()
	c := populateAtheleteCollection(session, dao.ms.DbName)

	var athelete AtheleteModel
	err := c.Find(bson.M{"friendly-url": friendlyName}).One(&athelete)

	if err != nil {
		switch err {
		case mgo.ErrNotFound:
			err = root.ErrDBErrNotFound
		}
	}

	return athelete, err
}

// FindAllAtheletes - find all atheletes
func (dao *AthleteDAO) FindAllAtheletes() ([]AtheleteModel, error) {
	iSession := dao.ms.getInitialSession()
	session := iSession.Copy()
	defer session.Close()
	c := populateAtheleteCollection(session, dao.ms.DbName)

	var atheletes []AtheleteModel
	err := c.Find(bson.M{}).All(&atheletes)

	return atheletes, err
}

func populateAtheleteCollection(session *mgo.Session, dbName string) *mgo.Collection {
	return session.DB(dbName).C("atheletes")
}

func createAtheleteIndices(session *mgo.Session, dbName string) {
	log.Println("Creating index for Atheletes")
	index := mgo.Index{
		Key:        []string{"friendly-url"},
		Unique:     true,
		DropDups:   false,
		Background: false,
		Sparse:     true,
	}
	c := populateAtheleteCollection(session, dbName)
	err := c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}

/*


func updateBook(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		isbn := pat.Param(r, "isbn")

		var book Book
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&book)
		if err != nil {
			ErrorWithJSON(w, "Incorrect body", http.StatusBadRequest)
			return
		}

		c := session.DB("store").C("books")

		err = c.Update(bson.M{"isbn": isbn}, &book)
		if err != nil {
			switch err {
			default:
				ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
				log.Println("Failed update book: ", err)
				return
			case mgo.ErrNotFound:
				ErrorWithJSON(w, "Book not found", http.StatusNotFound)
				return
			}
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
*/
