package data

// the model / domain entity
type UserRegistration struct {
	RequestId string `form:"-"` // uuid of web request
	ID        string `form:"-" bson:"_id, omitempty"`
	Firstname string `form:"firstname" binding:"required"`
	Lastname  string `form:"lastname" binding:"required"`
	//Email                 string    `form:"Email" binding:"required"`
	//Firma                 string    `form:"Firma"`
	//Schulungscode         string    `form:"Schulungscode" binding:"required"`
	//Datum                 time.Time `form:"Datum" binding:"required" time_format:"2006-01-02"`
	TermsAccepted bool `form:"termsAccepted" binding:"required"`
	//Confirmed             bool
	////Adresse, etc nicht dargestellt
}
