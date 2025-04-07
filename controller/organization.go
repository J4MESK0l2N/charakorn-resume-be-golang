package handlers

import (
	"context"
	"net/http"
	"time"
	"fmt"
	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
	"be-golang/models"
	"be-golang/database"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type OrganizationWithProjects struct {
	models.Organization
	Projects []models.Project `json:"projects"`
}


func CreateOrganization(c *gin.Context) {
	var org models.Organization
	if err := c.ShouldBindJSON(&org); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	now := time.Now().Format(time.RFC3339)
	org.ID = uuid.New().String()
	org.CreatedAt = now
	org.UpdatedAt = now

	var err error

	
	item := map[string]types.AttributeValue{
		"id":         &types.AttributeValueMemberS{Value: org.ID},
		"name":       &types.AttributeValueMemberM{Value: map[string]types.AttributeValue{"en": &types.AttributeValueMemberS{Value: org.Name.EN}, "th": &types.AttributeValueMemberS{Value: org.Name.TH}}},
		"image":      &types.AttributeValueMemberS{Value: org.Image},
		"join_date":  &types.AttributeValueMemberS{Value: org.JoinDate},
		"job_position": &types.AttributeValueMemberS{Value: org.JobPosition},
		"end_date":   &types.AttributeValueMemberS{Value: org.EndDate},
		"created_at": &types.AttributeValueMemberS{Value: org.CreatedAt},
		"updated_at": &types.AttributeValueMemberS{Value: org.UpdatedAt},
	}

	fmt.Println("Item to insert into DynamoDB:", item)

	_, err = db.Client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("organization"),
		Item:      item,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, org)
}

func GetOrganization(c *gin.Context) {
	id := c.Param("id")
	res, err := db.Client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("organization"),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil || res.Item == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "organization not found"})
		return
	}

	var org models.Organization
	err = attributevalue.UnmarshalMap(res.Item, &org)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unmarshal error"})
		return
	}

	c.JSON(http.StatusOK, org)
}

func UpdateOrganization(c *gin.Context) {
	id := c.Param("id")

	var org models.Organization
	if err := c.ShouldBindJSON(&org); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	org.ID = id
	org.UpdatedAt = time.Now().Format(time.RFC3339)

	item, err := attributevalue.MarshalMap(org)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "marshal error"})
		return
	}

	_, err = db.Client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("organization"),
		Item:      item,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "dynamodb update error"})
		return
	}

	c.JSON(http.StatusOK, org)
}

func DeleteOrganization(c *gin.Context) {
	id := c.Param("id")

	_, err := db.Client.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String("organization"),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "delete error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "organization deleted"})
}

func GetAllOrganizations(c *gin.Context) {
	orgResult, err := db.Client.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String("organization"),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	projectResult, err := db.Client.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String("project"),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	projectMap := make(map[string][]models.Project)
	for _, item := range projectResult.Items {
		var project models.Project
		err := attributevalue.UnmarshalMap(item, &project)
		fmt.Println(projectMap[project.OrgID])
		if err == nil {
			projectMap[project.OrgID] = append(projectMap[project.OrgID], project)
		}
	}

	var results []OrganizationWithProjects

	

	for _, item := range orgResult.Items {
		var org models.Organization
		err := attributevalue.UnmarshalMap(item, &org)

		fmt.Println(projectMap)
		
		if err == nil {
			orgWithProjects := OrganizationWithProjects{
				Organization: org,
				Projects:     projectMap[org.ID], 
			}
			results = append(results, orgWithProjects)
		}
	}

	c.JSON(http.StatusOK, results)
}




