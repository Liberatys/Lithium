package configuration

import "strings"

func ParseGivenConfigurationFileContent(fileContent string, seperatorSymbol string) map[string]string {
	configMap := make(map[string]string)
	fileLines := strings.Split(fileContent, "\n")
	for _, line := range fileLines {
		if len(line) > 0 {
			lineSlice := strings.Split(line, seperatorSymbol)
			configMap[strings.TrimSpace(lineSlice[0])] = strings.TrimSpace(lineSlice[1])
		}
	}
	return configMap
}
