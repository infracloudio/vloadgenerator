package src

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
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
		// accountPOSTRequest(appConfig.URL, &targets),
		customerPOSTRequest(appConfig.URL, &targets),
		// patientPOSTRequest(appConfig.URL, &targets),
		// generateGETRequests(appConfig.URL, &targets),
	}

	for index := 0; index < numberOfTargets; index++ {
		rand.Seed(time.Now().UnixNano())
		createTarget := targetType[rand.Intn(len(targetType))]
		createTarget()
	}

	log.WithFields(log.Fields{"Number of targets generated": len(targets)}).Info()

	attacker = vegeta.NewAttacker()
	var results vegeta.Results

	for res := range attacker.Attack(
		vegeta.NewStaticTargeter(targets...),
		uint64(appConfig.Rate),
		time.Duration(appConfig.Duration)*time.Second,
		"HSL attack") {

		metrics.Add(res)
		results.Add(res)
	}
	generatePlotReport(&results)
	generateTextReport(&metrics)
	metrics.Close()
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

func generateRandomCustomer() types.Customer {
	var customer types.Customer
	customer.CustomerID = rand.Intn(50000)
	customer.ClientID = rand.Intn(50000)
	customer.FirstName = "Agent-" + strconv.Itoa(rand.Intn(2000))
	customer.LastName = "K-" + strconv.Itoa(rand.Intn(2000))
	customer.DateOfBirth = "1963-10-27"
	customer.PhoneNumber = "123456"
	customer.SocialInsuranceNumber = "1234" + strconv.Itoa(rand.Intn(200))
	customer.Ssn = "1234ssn" + strconv.Itoa(rand.Intn(200))
	customer.Tin = "1234tin" + strconv.Itoa(rand.Intn(200))
	customer.Address = types.Address{
		Address1: "EC",
		City:     "Bangalore",
		State:    "Kar",
	}
	customer.Accounts = []types.Account{
		types.Account{
			AccountNumber: rand.Intn(200),
			Balance:       rand.Intn(20000),
			Interest:      rand.Intn(20),
			RoutingNumber: rand.Intn(900),
			Type:          "SAVING",
		},
	}
	return customer
}
