package src

import (
	"fmt"

	"github.com/infracloudio/vloadgenerator/src/types"
)

func sanityCheck(appConfig *types.AppConfig) error {
	names := []string{"hsl", "webgoat", "jenkins"}

	if !contains(names, appConfig.Name) {
		return fmt.Errorf("Invalid name : Please provide one of hsl , webg , jenkins")
	}
	if appConfig.Rate <= 0 || appConfig.Duration <= 0 {
		return fmt.Errorf("Rate / Duration cannot be zero or negative")
	}

	return nil
}

// Contains tells whether a contains x.
func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
