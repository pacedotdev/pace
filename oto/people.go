package public

// Person is a human who uses Pace.
type Person struct {
	// ID is the ID of the Person.
	// example: "5f19afce3979fb39"
	ID string
	// Username is the Person's username within the org.
	// example: "mat"
	Username string
	// Name is the name of the Person.
	// example: "Mat Ryer"
	Name string
	// PhotoURL is the URL of a picture of this Person.
	// example: "https://pace.dev/public/people/matryer.png"
	PhotoURL string
}
