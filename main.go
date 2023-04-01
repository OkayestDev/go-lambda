package main

import (
	"context"
	"golambda/src/repositories"

	"github.com/aws/aws-lambda-go/events"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
)

var ginLambda *ginadapter.GinLambda

func main() {
	// g := gin.Default()
	// g.GET("/health", controllers.Health)
	// ginLambda = ginadapter.New(g)
	// lambda.Start(Handler)

	user := repositories.UserRepo.Get("123", "123")

	println(user);
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, request)
}
