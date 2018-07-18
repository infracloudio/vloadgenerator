package types

type Accounts struct {
	Accounts []Account
}

type Account struct {
	Type          string `json:"type"`
	RoutingNumber int    `json:"routingNumber"`
	AccountNumber int    `json:"accountNumber"`
	Balance       int    `json:"balance"`
	Interest      int    `json:"interest"`
}

type Patient struct {
	FirstName       string `json:"patientFirstName"`
	LastName        string `json:"patientLastName"`
	DateOfBirth     string `json:"dateOfBirth"`
	Weight          int    `json:"patientWeight"`
	Height          int    `json:"patientHeight"`
	Medications     string `json:"medications"`
	BodyTemparature int    `json:"body_temp_deg_c"`
	HeartRate       int    `json:"heartRate"`
	PulseRate       int    `json:"pulse_rate"`
	BloodPressure   int    `json:"bpDiastolic"`
}

type Customer struct {
	FirstName             string   `json:"firstName"`
	LastName              string   `json:"lastName"`
	DateOfBirth           string   `json:"dateOfBirth"`
	Ssn                   string   `json:"ssn"`
	SocialInsuranceNumber string   `json:"socialInsurancenum"`
	Tin                   string   `json:"tin"`
	PhoneNumber           string   `json:"phoneNumber"`
	Address               Address  `json:"address"`
	Accounts              Accounts `json:"accounts"`
}

type Address struct {
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	City     string `json:"city"`
	State    string `json:"state"`
	ZipCode  string `json:"zipCode"`
}

// AppConfig creates a struct for the cli app.
type AppConfig struct {

	// Name of the app under attack.
	Name string
	// URL of the App
	URL string
	// Rate of requests per second
	Rate int
	// Duration of the attack
	Duration int
}
