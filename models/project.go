package models

type Project struct {
	ID        string         `json:"id" dynamodbav:"id"`
	Name      LocalizedField `json:"name" dynamodbav:"name"`
	JobTools  LocalizedField `json:"job_tools" dynamodbav:"job_tools"`
	Detail    LocalizedField `json:"detail" dynamodbav:"detail"`
	OrgID     string 		 `json:"org_id" dynamodbav:"org_id"`
	CreatedAt string         `json:"created_at" dynamodbav:"created_at"`
	UpdatedAt string         `json:"updated_at" dynamodbav:"updated_at"`
}