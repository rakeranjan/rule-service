package config

//Configuration :- struct for configurat
type Configuration struct {
	ProgramName string
}

var config *Configuration

// Config method is used to carry single instance of configuration
func Config() *Configuration {
	if config == nil {
		config = getConfig()
	}
	return config
}

func getConfig() *Configuration {
	config := Configuration{
		ProgramName: "rule-service",
	}
	return &config
}
