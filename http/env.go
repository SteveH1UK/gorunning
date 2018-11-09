package http

import (
	"github.com/SteveH1UK/gorunning/mongodb"
)

// Env - environmental dependencies for the http package
type Env struct {
	atheleteDAO mongodb.AthleteDAOInterface
}

// NewEnv - returns new environment
func NewEnv(atheleteDAO mongodb.AthleteDAOInterface) *Env {
	return &Env{atheleteDAO}
}
