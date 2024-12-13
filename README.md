# Faux API

## Overview
The Faux API is a Go-based framework designed for generating customizable dummy APIs. It is lightweight, modular, and supports multiple frameworks, including Gin and Fiber. This tool is perfect for rapid prototyping, API testing, and development environments where real backend logic is not yet implemented.

---

## Key Features
- **Customizable Responses**:
  - Define static responses for specific endpoints.
  - Implement dynamic responses based on incoming request data.
- **Framework Support**:
  - Compatible with both Gin and Fiber frameworks.
- **Modular Design**:
  - Pass endpoint configurations directly from your application code for maximum flexibility.
- **Default Responses**:
  - Automatically generates default mock responses for undefined configurations.

---

## Installation

1. Import the package in your Go project:
   ```go
   import "github.com/DevMaan707/dummy-api-gen/api"
   import "github.com/DevMaan707/dummy-api-gen/adapters"
   ```

3. Install dependencies:
   ```bash
   go mod tidy
   ```

---

## Usage

### Step 1: Parse Models
The `ParseModels` function reads your data models to generate APIs dynamically.

```go
models, err := api.ParseModels("./models")
if err != nil {
    log.Fatalf("Error parsing models: %v", err)
}
```

### Step 2: Define Custom Endpoint Configurations
Define the response logic for specific endpoints directly in your application:

```go
type EndpointConfig struct {
    StaticResponse  map[string]interface{}
    ConditionalLogic func(request map[string]interface{}) map[string]interface{}
}

customEndpointConfigs := map[string]*EndpointConfig{
    "/product": {
        StaticResponse: map[string]interface{}{
            "id": 1,
            "name": "Custom Product",
            "price": 99.99,
        },
    },
    "/user": {
        ConditionalLogic: func(request map[string]interface{}) map[string]interface{} {
            if userID, ok := request["user_id"].(string); ok && userID == "1" {
                return map[string]interface{}{
                    "id": 1,
                    "username": "John Doe",
                }
            }
            return nil
        },
    },
}
```

### Step 3: Generate APIs
Pass your router, models, and custom configurations to the API generator:

```go
err = api.GenerateAPIsWithConfig(router, models, customEndpointConfigs)
if err != nil {
    log.Fatalf("Error generating APIs: %v", err)
}
```

### Step 4: Run the Server
Start your server:

```go
app := gin.Default()
router := adapters.NewGinRouter(app)
if err := router.Run(":8080"); err != nil {
    log.Fatalf("Failed to start server: %v", err)
}
```

---

## Example
```bash
curl -X GET http://localhost:8080/test/product
```
Response:
```json
{
    "id": 1,
    "name": "Custom Product",
    "price": 99.99
}
```

---

## Contributing
1. Fork the repository.
2. Create a feature branch.
3. Commit your changes.
4. Open a pull request.

---

