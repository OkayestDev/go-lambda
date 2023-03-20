package testControllers

import (
	"learninggo/controllers"
	"testing"
)

func TestHealth(test *testing.T) {
	ginContext := GetTestGinContext()
	controllers.Health(ginContext)
}
