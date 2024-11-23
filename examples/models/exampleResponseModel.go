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

type AnotherResponseModel struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
