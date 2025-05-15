package controllers

import (
	"booker/modules/scheduling"
	"booker/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type SchedulingController struct {
	Service *scheduling.SchedulingService
}

func NewSchedulingController(r *gin.Engine, service *scheduling.SchedulingService) {
	sc := &SchedulingController{Service: service}
	r.POST("/schedulings", sc.Create)
	r.GET("/schedulings", sc.List)
	r.GET("/schedulings/:id", sc.FindOne)
	r.PUT("/schedulings/:id", sc.Update)
	r.DELETE("/schedulings/:id", sc.Delete)
}

func (sc *SchedulingController) Create(c *gin.Context) {
	var dto scheduling.CreateSchedulingDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := sc.Service.Create(dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

func (sc *SchedulingController) List(c *gin.Context) {
	query := utils.QueryOptions{
		Page:     utils.ParseQueryInt(c, "page", 0),
		PageSize: utils.ParseQueryInt(c, "pageSize", 10),
	}
	resp, err := sc.Service.List(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (sc *SchedulingController) FindOne(c *gin.Context) {
	id := c.Param("id")
	uuidVal, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	scheduling, err := sc.Service.FindOne(uuidVal)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, scheduling)
}

func (sc *SchedulingController) Update(c *gin.Context) {
	id := c.Param("id")
	uuidVal, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	var dto scheduling.UpdateSchedulingDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := sc.Service.Update(uuidVal, dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (sc *SchedulingController) Delete(c *gin.Context) {
	id := c.Param("id")
	uuidVal, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	if err := sc.Service.Delete(uuidVal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
