package controller

import (
	"github.com/RaymondSalim/API-server-template/server/models/request"
	"github.com/RaymondSalim/API-server-template/server/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin-swagger middleware
// swagger embed files

/*
	The controller contains the HTTP endpoints.
	The constructor takes the service as an argument.
	The controller can then call functions on the service in the endpoints.
	It takes care of translating JSON to request objects and response objects to JSON.

	The controller package does not contain interfaces for the controllers as they don’t need to be mocked for testing.
	It is only the dependencies of the controller that need to be mocked.
*/

type FooController struct {
	fooService service.FooService
}

func NewFooController(fooService service.FooService) FooController {
	return FooController{fooService: fooService}
}

// @BasePath /foo

// AddFoo godoc
// @Tags        Foo
// @Summary     Adds a new foo
// @Description Adds a new foo to the database
// @Accept      json
// @Produce     json
// @Param       FooRequest   body     request.AddFoo true "Request Body"
// @Param       X-Request-ID header   string         true "Request ID" default(174b9d6a-dafe-4f68-8e4b-6dcfbe7a804e)
// @Success     200          {object} response.FooResponse
// @Router      /foo/create [post]
func (fc FooController) AddFoo(c *gin.Context) {
	var fooRequest request.AddFoo
	err := c.ShouldBindJSON(&fooRequest)
	if err != nil {
		// handle error
	}

	fooResponse, err := fc.fooService.AddFoo(c, fooRequest)
	if err != nil {
		// handle error
	}
	c.JSON(http.StatusOK, fooResponse)
}

// GetFoo godoc
// @Tags        Foo
// @Summary     Get foo
// @Description Get foo from the database
// @Accept      json
// @Produce     json
// @Param       FooRequest   body     request.GetFoo true "Request Body"
// @Param       X-Request-ID header   string         true "865782e5-ccbf-4c5f-b967-f3df1fcd1f75"
// @Success     200          {object} response.FooResponse
// @Router      /foo/get [post]
func (fc FooController) GetFoo(c *gin.Context) {
	var fooRequest request.GetFoo
	err := c.ShouldBindJSON(&fooRequest)
	if err != nil {
		// handle error
	}

	fooResponse, err := fc.fooService.GetFoo(c, fooRequest.ID)
	if err != nil {
		// handle error
	}

	c.JSON(http.StatusOK, fooResponse)
}

// DeleteFoo godoc
// @Tags        Foo
// @Summary     Delete foo
// @Description Delete foo from the database
// @Accept      json
// @Produce     json
// @Param       FooRequest   body     request.DeleteFoo true "Request Body"
// @Param       X-Request-ID header   string            true "865782e5-ccbf-4c5f-b967-f3df1fcd1f75"
// @Success     200          {object} response.FooResponse
// @Router      /foo/delete [post]
func (fc FooController) DeleteFoo(c *gin.Context) {
	var fooRequest request.DeleteFoo
	err := c.ShouldBindJSON(&fooRequest)
	if err != nil {
		// handle error
	}

	err = fc.fooService.DeleteFoo(c, fooRequest.ID)
	if err != nil {
		// handle error
	}

	c.JSON(http.StatusOK, "")
}
