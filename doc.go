package medium_examples

// Authenticate will notify the program
// a user is attempting to connect.
// To get a userid please visit [https://example.com]
func Authenticate(userid string) error {
	//...
}

// ConnectToTCP will connect to the server
// following the protocol defined [RFC 000]
//
// [RFC 000]: https://example.com
func ConnectToTCP() error {

}

// AwesomeFactory will instantiate a new
// Awesome object. 
func AweSomeFactory() Awesome {
	return AweSome{ Ids : [10]string{} }
}

// Awesome is awesome
// To instantiate, refer to [medium_examples.AweSomeFactory] 
// Remember to [connect] first
//
// [connect]: medium_examples.ConnectToTCP
type Awesome struct {
	Ids []string
}
