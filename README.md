# Dummy API Generator

The **Dummy API Generator** is a tool for automatically generating RESTful APIs based on predefined Go models. It supports both `GET` and `POST` endpoints, enabling quick development and testing of APIs without manually writing routes or logic.

---

## Features

- **Dynamic API Generation**: Automatically create APIs based on `RequestModel` and `ResponseModel` structs.
- **GET and POST Requests**:
  - `GET` requests return dummy responses derived from `ResponseModel`.
  - `POST` requests validate incoming request bodies against `RequestModel` fields and return dummy responses.
- **Automatic Validation**:
  - Validates `POST` requests to ensure the body contains the required fields with correct types.
  - Returns appropriate error messages for missing or invalid fields.
- **Flexible Model Parsing**:
  - Supports multiple models, allowing dynamic generation of multiple endpoints.
- **Easy Integration**: Works seamlessly with the Gin framework.

---

## Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/DevMaan707/dummy-api-gen.git
   ```
2. **Install dependencies**:
   ```bash
   go mod tidy
   ```

---

## Usage

### 1. Define Your Models

Define your models in a separate package (e.g., `models`). Models must follow these conventions:
- Models ending with `RequestModel` define fields for `POST` request validation.
- Models ending with `ResponseModel` define fields for the dummy response.

#### Example Models

```go
package models

type ExampleRequestModel struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

type ExampleResponseModel struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}
```

### 2. Generate APIs

Use the `GenerateAPIs` function in your main file to dynamically create routes.

#### Example Main File

```go
package main

import (
	"log"

	dummyapi "github.com/DevMaan707/dummy-api-gen/dummyApi"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Path to models directory
	modelsPath := "./examples/models"

	// Generate APIs
	err := dummyapi.GenerateAPIs(router, modelsPath)
	if err != nil {
		log.Fatalf("Error generating APIs: %v", err)
	}

	log.Println("Server is running on http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
```

---

## API Endpoints

### **GET Request**
- **URL**: `/api/<ModelName>`
- **Response**: Dummy response based on the `ResponseModel`.

#### Example
For the `ExampleResponseModel`:

```json
{
    "id": 0,
    "name": "sample_text",
    "email": "sample_text"
}
```

---

### **POST Request**
- **URL**: `/test/<ModelName>`
- **Request Body**: Fields must match the `RequestModel`.
- **Response**: Dummy response based on the `ResponseModel`.

#### Example

##### Request Body:
```json
{
    "name": "John Doe",
    "email": "john.doe@example.com"
}
```

##### Response:
```json
{
    "id": 0,
    "name": "sample_text",
    "email": "sample_text"
}
```

##### Validation Errors:
- **Missing Field**:
  ```json
  {
      "error": "Missing field: email"
  }
  ```
- **Invalid Type**:
  ```json
  {
      "error": "Invalid type for field: name"
  }
  ```

---

## Contributing

Contributions are welcome! Please fork the repository, create a new branch, and submit a pull request.

---
