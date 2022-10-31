package registrierung

// Registrierungsservice mit Auswertung der Formularfelder
// implement http.Handler interface
import (
	"log"
	"net/http"
	"strconv"
)

type RegistrierungsHandler struct{}

// implement http.Handler.ServeHTTP()
func (rh *RegistrierungsHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Println("Could not parse form data:", err)
		rw.WriteHeader(http.StatusNotAcceptable)
		return
	}

	// ---- Auswertung der Formularfelder
	registrierung := &Registrierung{}
	registrierung.Firstname = req.Form.Get("Firstname")
	registrierung.Lastname = req.Form.Get("Lastname")
	registrierung.Email = req.Form.Get("Email")
	registrierung.Company = req.Form.Get("Company")
	registrierung.CourseCode = req.Form.Get("CourseCode")
	registrierung.Date = req.Form.Get("Date")
	// -- parse bool value
	bTermsAccepted, err := strconv.ParseBool(req.Form.Get("TermsAccepted"))
	if err != nil {
		log.Println("Bool parse error:", err)
		rw.WriteHeader(http.StatusNotAcceptable)
		return
	}
	registrierung.TermsAccepted = bTermsAccepted

	// ---- persist in db ...
	log.Printf("new registration: %v", registrierung)
	rw.WriteHeader(http.StatusCreated)
}
