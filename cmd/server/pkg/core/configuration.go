package core

type configuration struct{}

// default configuration for server
//
//	DefaultConfig = &Configuration{
//		PORT: 8080,
//		DSN:  os.Getenv("DSN"),
//	}
var DefaultConfig *configuration

func init() {
	DefaultConfig = &configuration{}
}

// creates and returns a pointer to new zero value based configuration
func NewConfiguration() *configuration {
	return new(configuration)
}
