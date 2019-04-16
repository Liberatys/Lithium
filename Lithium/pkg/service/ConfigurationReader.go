package service

import (
	"io/ioutil"
)

func ReadConfigurationFile(FileLocation string) string {
	data, err := ioutil.ReadFile(FileLocation) // just pass the file name
	if err != nil {
		return "None"
	}
	dataSequence := string(data[:])
	return dataSequence
}
