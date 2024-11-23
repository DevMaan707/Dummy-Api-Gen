
# Dummy API Generator

`dummy-api-generator` is a Go package designed to dynamically generate dummy APIs based on predefined models. It parses struct definitions in Go files and creates RESTful APIs for testing or prototyping purposes. These APIs return dynamic responses based on the fields defined in the models.

## Features

- **Dynamic API Generation**: Automatically create RESTful APIs for every model with a `ResponseModel` suffix.
- **Custom Responses**: Generate responses dynamically based on the fields and types defined in the models.
- **Easy Integration**: Use with the popular Gin framework to quickly spin up a server.
- **Scalable Design**: Add new models without modifying the codebase—just define your models, and APIs are generated automatically.
- **Logging**: Informative logs to help you monitor route generation and server activity.

---

## Installation

1. Install the package using `go get`:

   ```bash
   go get github.com/DevMaan707/dummy-api-gen
   ```

2. Add it to your project:

   ```go
   import dummyapi "github.com/DevMaan707/dummy-api-gen/dummyApi"
   ```

---

## Usage

### Directory Structure

Ensure your project structure includes a directory for models. For example:

```
dummy-api-generator/
├── examples/
│   ├── models/                # Define your models here
│       ├── example.go         # Example models file
├── main.go                    # Entry point for your application
```

### Example Code

#### Step 1: Define Your Models

Define your models in a `.go` file within the `models` directory. Models must end with `ResponseModel` to be automatically parsed.

```go
// examples/models/example.go

package models

type ExampleResponseModel struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

type AnotherResponseModel struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}
```

#### Step 2: Set Up Your Main Application

Use the `dummyapi.GenerateAPIs` function to set up routes dynamically.

```go
// examples/main.go

package main

import (
	"log"

	dummyapi "github.com/DevMaan707/dummy-api-gen/dummyApi"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	modelsPath := "./examples/models"

	// Generate APIs based on models
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

#### Step 3: Run the Server

1. Start the server:

   ```bash
   go run examples/main.go
   ```

2. Open your browser or use a tool like Postman to test the APIs.

   - **Example Endpoint**: `GET http://localhost:8080/api/ExampleResponseModel`
   - **Response**:

     ```json
     {
         "id": 0,
         "name": "sample_text",
         "email": "sample_text"
     }
     ```

---

## How It Works

1. **Model Parsing**: The package scans the directory specified by `modelsPath` for structs ending in `ResponseModel`.
2. **Dynamic API Creation**: For each parsed model, it generates a `GET` endpoint that returns a JSON response based on the model's fields.
3. **Dynamic Responses**: Field types determine default response values:
   - `int`: `0`
   - `string`: `"sample_text"`
   - Others: `null`

---

## Features in Detail

1. **Automatic Parsing**: You don’t need to configure routes manually. Just define models, and APIs are generated automatically.
2. **Dynamic Response Values**: Each field in the model is reflected in the API response with default placeholder values.
3. **Structured Logs**: Logs are generated for every route creation, making it easy to debug or monitor.

---

## Limitations

- **Model Naming**: Models must end with `ResponseModel` to be parsed.
- **Field Types**: Only primitive field types (e.g., `int`, `string`) are currently supported.

---

## Contributing

Contributions are welcome! Feel free to submit issues or pull requests to improve this package.

---


## Author

Created by **Aymaan**. Follow [GitHub](https://github.com/DevMaan707) for more projects.

--- 

