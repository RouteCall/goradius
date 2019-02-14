package state

import (
	"github.com/routecall/goradius/config"
)

type GlobalState struct {
	Config  *config.Config
	Current *State
}

type State struct {
	ID string
}

// create new global state of application
func New() *GlobalState {
	g := new(GlobalState)
	g.Config = config.Init()
	g.Config.Log.Debugf("new global state: %+v\n", g)
	return g
}
