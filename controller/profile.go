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


func CreateProfile(c *gin.Context) {
	var profile models.Profile
	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	now := time.Now().Format(time.RFC3339)
	profile.ID = uuid.New().String()
	profile.CreatedAt = now
	profile.UpdatedAt = now

	var err error

	
	item := map[string]types.AttributeValue{
		"id":            &types.AttributeValueMemberS{Value: profile.ID},
		"name":          &types.AttributeValueMemberM{Value: map[string]types.AttributeValue{"en": &types.AttributeValueMemberS{Value: profile.Name.EN}, "th": &types.AttributeValueMemberS{Value: profile.Name.TH}}},
		"job_position":  &types.AttributeValueMemberM{Value: map[string]types.AttributeValue{"en": &types.AttributeValueMemberS{Value: profile.JobPosition.EN}, "th": &types.AttributeValueMemberS{Value: profile.JobPosition.TH}}},
		"detail":        &types.AttributeValueMemberM{Value: map[string]types.AttributeValue{"en": &types.AttributeValueMemberS{Value: profile.Detail.EN}, "th": &types.AttributeValueMemberS{Value: profile.Detail.TH}}},
		"email":         &types.AttributeValueMemberS{Value: profile.Email},
		"address":       &types.AttributeValueMemberS{Value: profile.Address},
		"phone_number":  &types.AttributeValueMemberS{Value: profile.PhoneNumber},
		"created_at":    &types.AttributeValueMemberS{Value: profile.CreatedAt},
		"updated_at":    &types.AttributeValueMemberS{Value: profile.UpdatedAt},
	}

	fmt.Println("Item to insert into DynamoDB:", item)

	_, err = db.Client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("profile"),
		Item:      item,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}

func DeleteProfile(c *gin.Context) {
	id := c.Param("id")

	_, err := db.Client.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String("profile"),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "delete error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "profile deleted"})
}


func GetProfile(c *gin.Context) {
	result, err := db.Client.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String("profile"), 
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var profile []models.Profile
	for _, item := range result.Items {
		var pf models.Profile

		err := attributevalue.UnmarshalMap(item, &pf)
		if err != nil {
			fmt.Println("Unmarshal error:", err)
		}

		profile = append(profile, pf)
	}

	c.JSON(http.StatusOK, profile)
}