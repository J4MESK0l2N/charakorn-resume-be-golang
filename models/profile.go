package models

type Profile struct {
	ID        		string         `json:"id" dynamodbav:"id"`
	Name      		LocalizedField `json:"name" dynamodbav:"name"`
	JobPosition     LocalizedField `json:"job_position" dynamodbav:"job_position"`
	Detail  		LocalizedField `json:"detail" dynamodbav:"detail"`
	Address  		string 		   `json:"address" dynamodbav:"address"`
	Email   		string 		   `json:"email" dynamodbav:"email"`
	PhoneNumber   	string 		   `json:"phone_number" dynamodbav:"phone_number"`
	CreatedAt 		string         `json:"created_at" dynamodbav:"created_at"`
	UpdatedAt 		string         `json:"updated_at" dynamodbav:"updated_at"`
}
