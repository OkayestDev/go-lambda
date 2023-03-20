package testControllers

import (
	"learninggo/src/controllers"
	"testing"
)

func TestHealth(test *testing.T) {
	ginContext := GetTestGinContext()
	controllers.Health(ginContext)
}
