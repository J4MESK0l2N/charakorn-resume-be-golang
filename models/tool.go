package models

type Tool struct {
	ID        string         `json:"id" dynamodbav:"id"`
	Name      string		 `json:"name" dynamodbav:"name"`
	Image     string 		 `json:"image" dynamodbav:"image"`
	Detail    string 		 `json:"detail" dynamodbav:"detail"`
	CreatedAt string         `json:"created_at" dynamodbav:"created_at"`
	UpdatedAt string         `json:"updated_at" dynamodbav:"updated_at"`
}