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
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)


func CreateProject(c *gin.Context) {
	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	now := time.Now().Format(time.RFC3339)
	project.ID = uuid.New().String()
	project.CreatedAt = now
	project.UpdatedAt = now

	var err error

	
	item := map[string]types.AttributeValue{
		"id":         &types.AttributeValueMemberS{Value: project.ID},
		"name":       &types.AttributeValueMemberM{Value: map[string]types.AttributeValue{"en": &types.AttributeValueMemberS{Value: project.Name.EN}, "th": &types.AttributeValueMemberS{Value: project.Name.TH}}},
		"job_tools":  &types.AttributeValueMemberM{Value: map[string]types.AttributeValue{"en": &types.AttributeValueMemberS{Value: project.JobTools.EN}, "th": &types.AttributeValueMemberS{Value: project.JobTools.TH}}},
		"detail":     &types.AttributeValueMemberM{Value: map[string]types.AttributeValue{"en": &types.AttributeValueMemberS{Value: project.Detail.EN}, "th": &types.AttributeValueMemberS{Value: project.Detail.TH}}},
		"org_id": 	  &types.AttributeValueMemberS{Value: project.OrgID},
		"created_at": &types.AttributeValueMemberS{Value: project.CreatedAt},
		"updated_at": &types.AttributeValueMemberS{Value: project.UpdatedAt},
	}

	fmt.Println("Item to insert into DynamoDB:", item)

	_, err = db.Client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("project"),
		Item:      item,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, project)
}


func DeleteProject(c *gin.Context) {
	id := c.Param("id")

	_, err := db.Client.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String("project"),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "delete error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "project deleted"})
}