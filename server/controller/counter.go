package controller

import (
	"encoding/json"
	"errors"
	"github.com/Novometrix/web-server-template/server/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
)

type CounterController struct {
	countService service.CounterService
}

func NewCounterController(countService service.CounterService) CounterController {
	return CounterController{countService: countService}
}

// @BasePath /counter

// GetLastCounter godoc
// @Tags        Counter
// @Summary     Get Last counter
// @Description Get Last Counter
// @Accept      json
// @Produce     html
// @Success     200 {object} response.FooResponse
// @Router      /counter/get [get]
func (cc CounterController) GetLastCounter(c *gin.Context) {
	count, err := cc.countService.GetLastCounter(c)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		log.Errorf("failed to get last counter with error: %v", err)
		c.String(http.StatusBadRequest, "not found")
		return
	}
	b, _ := json.MarshalIndent(count, "", "\t")
	c.String(http.StatusOK, "count: %+v", string(b))
}

// AddCounter godoc
// @Tags        Counter
// @Summary     Add counter
// @Description Add Counter
// @Accept      json
// @Produce     html
// @Success     200 {object} response.FooResponse
// @Router      /counter/add [post]
func (cc CounterController) AddCounter(c *gin.Context) {
	_ = cc.countService.PublishAddCounter(c)

	c.String(http.StatusOK, "ok")
}
