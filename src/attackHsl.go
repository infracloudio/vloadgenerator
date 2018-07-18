package src

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/infracloudio/vloadgenerator/src/types"
	log "github.com/sirupsen/logrus"
	vegeta "github.com/tsenart/vegeta/lib"
)

func GenerateHSLAttack(appConfig *types.AppConfig) {

	var targets []vegeta.Target
	var numberOfTargets int
	var metrics vegeta.Metrics
	var attacker *vegeta.Attacker

	numberOfTargets = appConfig.Rate * appConfig.Duration

	targetType := []func(){
		accountPOSTRequest(appConfig.URL, &targets),
		customerPOSTRequest(appConfig.URL, &targets),
		patientPOSTRequest(appConfig.URL, &targets),
		generateGETRequests(appConfig.URL, &targets),
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

func accountPOSTRequest(url string, targets *[]vegeta.Target) func() {

	return func() {
		account := generateRandomAccount()
		log.WithFields(log.Fields{"Account: ": account}).Debug()
		body, err := json.Marshal(account)
		check(err)
		var header = make(http.Header)
		header.Add("content-type", "application/json")

		target := vegeta.Target{
			Method: http.MethodPost,
			URL:    url + "/account",
			Body:   body,
			Header: header,
		}

		addValue(targets, target)
	}
}

func customerPOSTRequest(url string, targets *[]vegeta.Target) func() {
	return func() {
		customer := generateRandomCustomer()
		log.WithFields(log.Fields{"Customer: ": customer}).Debug()
		body, err := json.Marshal(customer)
		check(err)
		var header = make(http.Header)
		header.Add("content-type", "application/json")

		target := vegeta.Target{
			Method: http.MethodPost,
			URL:    url + "/customers",
			Body:   body,
			Header: header,
		}

		addValue(targets, target)
	}

}

func patientPOSTRequest(url string, targets *[]vegeta.Target) func() {

	return func() {
		patient := generateRandomPatient()
		log.WithFields(log.Fields{"Patient: ": patient}).Debug()
		body, err := json.Marshal(patient)
		check(err)
		var header = make(http.Header)
		header.Add("content-type", "application/json")

		target := vegeta.Target{
			Method: http.MethodPost,
			URL:    url + "/patients",
			Body:   body,
			Header: header,
		}

		addValue(targets, target)
	}
}

func GenerateLoadData(count int, duration int, api string) {

	numberOfTargets := count * duration

	fmt.Println("Generating LoadData for number of requests", count, api)
	var targets []vegeta.Target
	// targets := make([]vegeta.Target, numberOfTargets)

	rand.Seed(time.Now().UnixNano())
	ftype := []func(){generateAccounts(api, &targets), generateGETRequests(api, &targets)}

	for index := 0; index < numberOfTargets; index++ {
		generator := ftype[rand.Intn(len(ftype))]
		generator()
	}

	log.WithFields(log.Fields{"Number of targets generated": len(targets)}).Debug()

	rate := uint64(count) // per second
	du := time.Duration(duration) * time.Second

	for index := 0; index < len(targets); index++ {
		fmt.Println(string(targets[index].Body), targets[index].URL)
	}

	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	for res := range attacker.Attack(vegeta.NewStaticTargeter(targets...), rate, du, "abc") {
		fmt.Println(res.Error)
		metrics.Add(res)
	}
	metrics.Close()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func generateAccounts(api string, targets *[]vegeta.Target) func() {

	return func() {
		account := generateRandomAccount()
		log.WithFields(log.Fields{"Account: ": account}).Debug()
		body, err := json.Marshal(account)
		check(err)
		var header = make(http.Header)
		header.Add("content-type", "application/json")

		target := vegeta.Target{
			Method: http.MethodPost,
			URL:    api + "/account",
			Body:   body,
			Header: header,
		}

		addValue(targets, target)
	}
}

func generateGETRequests(url string, targets *[]vegeta.Target) func() {

	return func() {
		appURLs := []string{
			"/customers",
			"/account",
			"/patients",
			"/saveSettings",
			"/loadSettings",
			"/customers/3",
			"/account/2",
			"/customers/1",
			"/error",
			"/debugEscaped?firstName=%22%22",
			"/account/1",
			"/account/3",
			"/search/user?foo=new%20java.lang.ProcessBuilder(%7B%27%2Fbin%2Fbash%27%2C%27-c%27%2C%27echo%203vilhax0r%3E%2Ftmp%2Fhacked%27%7D).start()",
			"/debug?customerId=ID-4242&clientId=1&firstName=%22%22&lastName=%22%22&dateOfBirth=10-11-17&ssn=%22%22&socialSecurityNum=%22%22&tin=%22%22&phoneNumber=%22%22",
			"/debugEscaped?firstName=%22%22",
			"/admin/login"}

		target := vegeta.Target{
			Method: http.MethodGet,
			URL:    url + appURLs[rand.Intn(len(appURLs))],
		}
		addValue(targets, target)
	}

}

func addValue(s *[]vegeta.Target, target vegeta.Target) {
	*s = append(*s, target)
	// fmt.Printf("In addValue: s is %v\n", s)
}

func generateRandomAccount() types.Account {
	var account types.Account
	accountType := []string{"SAVING", "CHECKING"}
	account.RoutingNumber = rand.Intn(50000)
	account.Balance = rand.Intn(50000)
	account.Interest = rand.Intn(15)
	account.Type = accountType[rand.Intn(len(accountType))]
	return account
}

func generateRandomPatient() types.Patient {
	var patient types.Patient
	patient.FirstName = "Agent"
	patient.LastName = "Jay"
	patient.DateOfBirth = "21-11-1973"
	patient.HeartRate = rand.Intn(120)
	patient.Height = rand.Intn(200)
	patient.Weight = rand.Intn(500)
	patient.PulseRate = rand.Intn(200)
	patient.BloodPressure = rand.Intn(200)
	patient.BodyTemparature = rand.Intn(100)
	patient.Medications = "combiflam,crocin,dolo,lanzol"

	return patient
}

func generateRandomCustomer() types.Customer {
	var customer types.Customer
	customer.FirstName = "Agent"
	customer.LastName = "Kay"
	customer.DateOfBirth = "13-07-1963"
	customer.PhoneNumber = "123456"
	customer.SocialInsuranceNumber = "1234"
	customer.Ssn = "1234ssn"
	customer.Tin = "1234tin"
	customer.Address = types.Address{
		Address1: "EC",
		City:     "Bangalore",
		State:    "Kar",
	}
	customer.Accounts = types.Accounts{
		Accounts: []types.Account{
			types.Account{
				AccountNumber: 312,
				Balance:       120000,
				Interest:      5,
				RoutingNumber: 456,
				Type:          "SAVING",
			},
		},
	}

	return customer
}
