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
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)


func CreateTool(c *gin.Context) {
	var tool models.Tool
	if err := c.ShouldBindJSON(&tool); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	now := time.Now().Format(time.RFC3339)
	tool.ID = uuid.New().String()
	tool.CreatedAt = now
	tool.UpdatedAt = now

	var err error

	
	item := map[string]types.AttributeValue{
		"id":            &types.AttributeValueMemberS{Value: tool.ID},
		"name":          &types.AttributeValueMemberS{Value: tool.Name},
		"image":  		 &types.AttributeValueMemberS{Value: tool.Image},
		"detail":        &types.AttributeValueMemberS{Value: tool.Detail},
		"created_at":    &types.AttributeValueMemberS{Value: tool.CreatedAt},
		"updated_at":    &types.AttributeValueMemberS{Value: tool.UpdatedAt},
	}

	fmt.Println("Item to insert into DynamoDB:", item)

	_, err = db.Client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("tool"),
		Item:      item,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tool)
}

func DeleteTool(c *gin.Context) {
	id := c.Param("id")

	_, err := db.Client.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String("tool"),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "delete error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "tool deleted"})
}


func GetTools(c *gin.Context) {
	result, err := db.Client.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String("tool"), 
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var tool []models.Tool
	for _, item := range result.Items {
		var t models.Tool

		err := attributevalue.UnmarshalMap(item, &t)
		if err != nil {
			fmt.Println("Unmarshal error:", err)
		}

		tool = append(tool, t)
	}

	c.JSON(http.StatusOK, tool)
}