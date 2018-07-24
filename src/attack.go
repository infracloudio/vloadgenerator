package src

import (
	"os"

	"github.com/infracloudio/vloadgenerator/src/types"
	log "github.com/sirupsen/logrus"
	vegeta "github.com/tsenart/vegeta/lib"
)

type AttackTargets struct {
	targets []vegeta.Target
}

func init() {

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func Attack(appConfig *types.AppConfig) {
	err := sanityCheck(appConfig)

	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	err = testConnectivity(appConfig.URL)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	// List of targets to be exercised , whether thats GET or POST
	//var targets []vegeta.Target

	if appConfig.Name == "hsl" {
		GenerateHSLAttack(appConfig)
	}

	if appConfig.Name == "webgoat" {
		log.Debug("WIP")
	}

	if appConfig.Name == "jenkins" {

	}
}
