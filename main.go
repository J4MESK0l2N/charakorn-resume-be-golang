package main

import (
	"github.com/gin-gonic/gin"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"be-golang/database"
	"be-golang/routes"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"	
)

var ginLambda *ginadapter.GinLambda

func init() {
	db.Init()

	r := gin.Default()

	routes.OrganizationRoutes(r)
	routes.ProjectRoutes(r)
	routes.ProfileRoutes(r)
	routes.ToolRoutes(r)

	ginLambda = ginadapter.New(r)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.Proxy(request)
}

func main() {
	lambda.Start(handler)
}
