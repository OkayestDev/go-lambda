package testControllers

import (
	"golambda/src/controllers"
	"testing"
)

func TestHealth(test *testing.T) {
	ginContext := GetTestGinContext()
	controllers.Health(ginContext)
}
