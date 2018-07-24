package src

import (
	"fmt"
	"net"
	"net/url"
	"time"

	"github.com/infracloudio/vloadgenerator/src/types"
	log "github.com/sirupsen/logrus"
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

func testConnectivity(s string) error {
	var conn net.Conn
	log.Debug("Testing connectivity to app before starting the test.")

	u, err := url.Parse(s)
	if err != nil {
		log.Fatal(err)
	}

	c := make(chan string, 1)
	go func() {
		for {
			conn, err = net.Dial("tcp", net.JoinHostPort(u.Hostname(), u.Port()))
			if conn != nil {
				conn.Close()
				break
			}
			// wait 5 seconds before retrying.
			time.Sleep(5 * time.Second)
		}
		c <- "true"
		defer close(c)
	}()

	select {
	case _ = <-c:
		log.Debug("Connection established")
		return nil
	case <-time.After(2 * time.Minute):
		return fmt.Errorf("Could not establish connection")
	}
}
