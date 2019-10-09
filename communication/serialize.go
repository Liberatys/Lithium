package communication

import (
	"encoding/json"
	"github.com/Liberatys/Sanctuary/load"
)

type SerializedService struct {
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	IP          string    `json:"ip"`
	Port        string    `json:"port"`
	Load        load.Load `json:"load"`
	LoadIndex   int64
}

func Serialize(serializedService SerializedService) string {
	emp, _ := json.Marshal(serializedService)
	return string(emp)
}

func Decode(serializedService string) SerializedService {
	bytes := []byte(serializedService)
	var serializeService SerializedService
	json.Unmarshal(bytes, &serializeService)
	if serializedService.Load != nil {
		serializedService.LoadIndex = ((serializedService.Load.Network) * 10) + (serializedService.Load.GPU)
	}
	return serializeService
}
