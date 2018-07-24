package src

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/infracloudio/vloadgenerator/src/types"
	log "github.com/sirupsen/logrus"
	vegeta "github.com/tsenart/vegeta/lib"
)

func GenerateJenkinsAttack(appConfig *types.AppConfig) {

	var targets []vegeta.Target
	var numberOfTargets int
	var metrics vegeta.Metrics
	var attacker *vegeta.Attacker

	numberOfTargets = appConfig.Rate * appConfig.Duration

	targetType := []func(){
		createJobRequest(appConfig.URL, &targets),
	}

	for index := 0; index < numberOfTargets; index++ {
		rand.Seed(time.Now().UnixNano())
		createTarget := targetType[rand.Intn(len(targetType))]
		createTarget()
	}

	log.WithFields(log.Fields{"Number of targets generated": len(targets)}).Debug()

	attacker = vegeta.NewAttacker()

	for res := range attacker.Attack(
		vegeta.NewStaticTargeter(targets...),
		uint64(appConfig.Rate),
		time.Duration(appConfig.Duration)*time.Second,
		"HSL attack") {

		log.Debug(res.Error)
		metrics.Add(res)
	}
	metrics.Close()
	log.WithFields(log.Fields{"99th percentile": metrics.Latencies.P99, "rate": metrics.Rate, "requests": metrics.Requests, "duration": metrics.Duration}).Info()
}

func createJobRequest(url string, targets *[]vegeta.Target) func() {

	return func() {
		body, err := ioutil.ReadFile("src/utils/jenkins-job-config.xml")
		check(err)

		jobName := "generated-job-name-" + generateUUID()

		// basicauth
		auth := basicAuth("admin", "admin")

		var header = make(http.Header)
		header.Add("content-type", "application/xml")
		header.Add("Authorization", "Basic "+auth)

		target := vegeta.Target{
			Method: http.MethodPost,
			URL:    url + "/createItem?name=" + jobName,
			Body:   body,
			Header: header,
		}
		addValue(targets, target)

	}
}

// URL to create a job

// http://localhost:8080/createItem?name=testjob2

// url to get config.xml of an existing job

// http://localhost:8080/job/testjob/config.xml

// config.xml

// <?xml version='1.1' encoding='UTF-8'?>
// <project>
//     <description>test job 2 created</description>
//     <keepDependencies>false</keepDependencies>
//     <properties>
//         <com.dabsquared.gitlabjenkins.connection.GitLabConnectionProperty plugin="gitlab-plugin@1.5.8">
//             <gitLabConnection></gitLabConnection>
//         </com.dabsquared.gitlabjenkins.connection.GitLabConnectionProperty>
//     </properties>
//     <scm class="hudson.scm.NullSCM"/>
//     <canRoam>true</canRoam>
//     <disabled>false</disabled>
//     <blockBuildWhenDownstreamBuilding>false</blockBuildWhenDownstreamBuilding>
//     <blockBuildWhenUpstreamBuilding>false</blockBuildWhenUpstreamBuilding>
//     <triggers/>
//     <concurrentBuild>false</concurrentBuild>
//     <builders/>
//     <publishers/>
//     <buildWrappers/>
// </project>
