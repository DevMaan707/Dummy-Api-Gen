package api

type EndpointConfig struct {
	StaticResponse   map[string]interface{}
	ConditionalLogic func(request map[string]interface{}) map[string]interface{}
}

var EndpointConfigs = map[string]*EndpointConfig{}
