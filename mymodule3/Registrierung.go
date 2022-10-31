package registrierung

// das Domaenenmodell
type Registrierung struct {
	Firstname     string
	Lastname      string
	Email         string
	Company       string
	CourseCode    string
	Date          string
	TermsAccepted bool
	// TODO: address etc
}
