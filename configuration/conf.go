package configuration

type Configuration struct {
	FileTyp              string
	FilePath             string
	ConfigurationContent string
}

func NewConfiguration(path string) Configuration {
	configuration := Configuration{
		FilePath:             path,
		FileTyp:              "configuration",
		ConfigurationContent: "",
	}
	return configuration
}

//TODO: implement the reading as well as the parsing of the configuration

func (configuration *Configuration) ReadConfig() {

}
func (configuration *Configuration) ParseConfig() {

}
