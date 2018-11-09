package mongodb

import (
	"log"

	"github.com/globalsign/mgo"
)

// MongoSession - contains a pointer mgo Session
type MongoSession struct {
	DbName  string
	session *mgo.Session
}

var initialSession *MongoSession

// Init - gets an initial connection to the Mongo database
func Init(mongoURL string, dbName string) (*MongoSession, error) {

	mgoSession, err := mgo.Dial(mongoURL)
	if err != nil {
		return nil, err
	}

	mgoSession.SetMode(mgo.Monotonic, true)
	initialSession = &MongoSession{dbName, mgoSession}

	createIndicesIfRequired(initialSession.session, dbName)
	return initialSession, nil
}

// Shutdown - shuts down the initial MongoDB sessio
func (m *MongoSession) Shutdown() {
	if initialSession.session != nil {
		initialSession.session.Close()
	}
}

func (m *MongoSession) getInitialSession() *mgo.Session {
	if initialSession.session == nil {
		log.Println("Initial session is unexpectidy empty!")
		panic(-1)
	}
	return initialSession.session
}

func createIndicesIfRequired(iSession *mgo.Session, dbName string) {
	session := iSession.Copy()
	defer session.Close()

	createAtheleteIndices(session, dbName)
}
