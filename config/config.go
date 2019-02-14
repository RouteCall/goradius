package config

import (
	"os"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	Log  *log.Logger
	Envs Envs
}

type Envs map[string]string

// new environment variables
func (c *Config) NewEnvs() {
	// supported environment variables
	e := Envs{
		"REDIS_URL":   "",
		"RADIUS_ADDR": "",
	}
	for k := range e {
		e[k] = os.Getenv(k)
	}

	c.Envs = e
}

// initialize all application configuration
func Init() *Config {
	c := new(Config)

	//log
	c.Log = log.New()
	c.Log.Formatter = &log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	}
	c.Log.Level = log.DebugLevel

	c.NewEnvs()

	c.Log.Debugf("init config: %+v\n", c)
	return c
}
