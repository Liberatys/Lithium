package communication

import (
	"encoding/json"
)

type SerializedService struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	IP          string `json:"ip"`
	Port        string `json:"port"`
}

func Serialize(serializedService SerializedService) string {
	emp, _ := json.Marshal(serializedService)
	return string(emp)
}

func Decode(serializedService string) SerializedService {
	bytes := []byte(serializedService)
	var serializeService SerializedService
	json.Unmarshal(bytes, &serializeService)
	return serializeService
}
