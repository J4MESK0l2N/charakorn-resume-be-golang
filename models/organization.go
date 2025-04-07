package models

type Organization struct {
	ID        		 string         `json:"id" dynamodbav:"id"`
	Name      		 LocalizedField `json:"name" dynamodbav:"name"`
	Image     		 string         `json:"image" dynamodbav:"image"`
	JobPosition      string         `json:"job_position" dynamodbav:"job_position"`
	JoinDate  		 string         `json:"join_date" dynamodbav:"join_date"`
	EndDate   		 string         `json:"end_date" dynamodbav:"end_date"`
	CreatedAt 		 string         `json:"created_at" dynamodbav:"created_at"`
	UpdatedAt 		 string         `json:"updated_at" dynamodbav:"updated_at"`
}
