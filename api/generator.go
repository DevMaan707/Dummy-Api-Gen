package api

import (
	"encoding/json"
	"net/http"

	"github.com/DevMaan707/dummy-api-gen/adapters"
	"github.com/DevMaan707/dummy-api-gen/shared"
)

func GenerateAPIs(router adapters.Router, models []shared.ModelData) error {
	apiGroup := router.Group("/test")

	for _, model := range models {
		path := "/" + model.Name

		apiGroup.GET(path, func(w http.ResponseWriter, r *http.Request) {
			response := generateMockResponse(model.ResponseFields)
			writeJSONResponse(w, response)
		})

		if len(model.RequestFields) > 0 {
			apiGroup.POST(path, func(w http.ResponseWriter, r *http.Request) {
				var body map[string]interface{}
				if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
					http.Error(w, "Invalid request body", http.StatusBadRequest)
					return
				}
				for field, typ := range model.RequestFields {
					if !validateFieldType(body[field], typ) {
						http.Error(w, "Invalid field: "+field, http.StatusBadRequest)
						return
					}
				}

				response := generateMockResponse(model.ResponseFields)
				writeJSONResponse(w, response)
			})
		}
	}

	return nil
}

func generateMockResponse(fields map[string]string) map[string]interface{} {
	response := make(map[string]interface{})
	for name, typ := range fields {
		switch typ {
		case "string":
			response[name] = "example"
		case "int":
			response[name] = 123
		default:
			response[name] = nil
		}
	}
	return response
}

func validateFieldType(value interface{}, expectedType string) bool {
	switch expectedType {
	case "string":
		_, ok := value.(string)
		return ok
	case "int":
		_, ok := value.(float64)
		return ok
	default:
		return false
	}
}

func writeJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
