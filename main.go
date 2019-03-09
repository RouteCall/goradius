package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/routecall/goradius/state"
)

// the init function for set logs default values
func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
}

func main() {
	globalState := state.New()
	globalState.Config.Log.Infof("globalState: %+v\n", globalState)
}
