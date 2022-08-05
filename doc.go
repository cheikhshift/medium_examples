package medium_examples

// Authenticate will notify the program
// a user is attempting to connect.
// To get a userid please visit [https://example.com]
func Authenticate(userid string) error {
	//...
}

// AwesomeFactory will instantiate a new
// Awesome object. 
func AweSomeFactory() Awesome {
	return AweSome{ Ids : [10]string{} }
}

// Awesome is awesome
// To instantiate, refer to [medium_examples.AweSomeFactory] 
type Awesome struct {
	Ids []string
}
