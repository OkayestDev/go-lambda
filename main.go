// package main

// import (
// 	"golambda/controllers"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	router := gin.Default()
// 	router.GET("/ping", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "pong",
// 		})
// 	})

// 	router.GET("/health", controllers.Health)
// 	router.GET("/users", controllers.GetUser)
// 	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
//}

package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func main() {
	g := gin.Default()
	g.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	ginLambda = ginadapter.New(g)
	lambda.Start(Handler)
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, request)
}
