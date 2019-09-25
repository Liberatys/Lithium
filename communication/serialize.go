package communication

import (
	"encoding/json"
	"github.com/Liberatys/Sanctuary/service"
)

type SerializedService struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	IP          string `json:"ip"`
	Port        string `json:"port"`
}

func Serialize(service *service.Service) string {
	serService := SerializedService{
		Name:        service.Name,
		Type:        service.Type,
		Description: service.Description,
		IP:          service.IP,
		Port:        service.Port,
	}
	emp, _ := json.Marshal(serService)
	return string(emp)
}

func Decode(serializedService string) SerializedService {
	bytes := []byte(serializedService)
	var serializeService SerializedService
	json.Unmarshal(bytes, &serializeService)
	return serializeService
}
