package api

import (
	"encoding/json"
	"net/http"

	"github.com/DevMaan707/faux-api/adapters"
	"github.com/DevMaan707/faux-api/shared"
)

func GenerateAPIsWithConfig(router adapters.Router, models []shared.ModelData, configs map[string]*EndpointConfig) error {
	apiGroup := router.Group("/test")

	for _, model := range models {
		path := "/" + model.Name
		apiGroup.GET(path, func(w http.ResponseWriter, r *http.Request) {
			var request map[string]interface{}
			response := generateMockResponseWithConfig(path, model.ResponseFields, request, configs)
			writeJSONResponse(w, response)
		})
		if len(model.RequestFields) > 0 {
			apiGroup.POST(path, func(w http.ResponseWriter, r *http.Request) {
				var body map[string]interface{}
				if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
					http.Error(w, "Invalid request body", http.StatusBadRequest)
					return
				}

				response := generateMockResponseWithConfig(path, model.ResponseFields, body, configs)
				writeJSONResponse(w, response)
			})
		}
	}

	return nil
}

func generateMockResponseWithConfig(endpoint string, fields map[string]string, request map[string]interface{}, configs map[string]*EndpointConfig) map[string]interface{} {
	response := make(map[string]interface{})
	if config, exists := configs[endpoint]; exists {
		if config.StaticResponse != nil {
			for key, value := range config.StaticResponse {
				response[key] = value
			}
			return response
		}
		if config.ConditionalLogic != nil {
			customResponse := config.ConditionalLogic(request)
			for key, value := range customResponse {
				response[key] = value
			}
			return response
		}
	}
	for name, typ := range fields {
		if _, exists := response[name]; !exists {
			switch typ {
			case "string":
				response[name] = "example"
			case "int":
				response[name] = 123
			default:
				response[name] = nil
			}
		}
	}

	return response
}

func writeJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
