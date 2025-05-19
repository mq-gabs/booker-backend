package schedulingprofile

import (
	"booker/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type SchedulingProfileController struct {
	Service *SchedulingProfileService
}

func SetRoutes(r *gin.Engine) {
	repo := NewSchedulingProfileMemoryRepository()
	service := NewSchedulingProfileService(repo)

	pc := &SchedulingProfileController{Service: service}

	g := r.Group("/scheduling-profiles")

	g.POST("/", pc.Create)
	g.GET("/", pc.List)
	g.GET("/:id", pc.FindOne)
	g.PUT("/:id", pc.Update)
	g.DELETE("/:id", pc.Delete)
}

func (pc *SchedulingProfileController) Create(c *gin.Context) {
	var dto CreateSchedulingProfileDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := pc.Service.Create(dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

func (pc *SchedulingProfileController) List(c *gin.Context) {
	query := utils.QueryOptions{
		Page:     utils.ParseQueryInt(c, "page", 0),
		PageSize: utils.ParseQueryInt(c, "pageSize", 10),
	}
	resp, err := pc.Service.List(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (pc *SchedulingProfileController) FindOne(c *gin.Context) {
	id := c.Param("id")
	uuidVal, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	profile, err := pc.Service.FindOne(uuidVal)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, profile)
}

func (pc *SchedulingProfileController) Update(c *gin.Context) {
	id := c.Param("id")
	uuidVal, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	var dto UpdateSchedulingProfileDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := pc.Service.Update(uuidVal, dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (pc *SchedulingProfileController) Delete(c *gin.Context) {
	id := c.Param("id")
	uuidVal, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	if err := pc.Service.Delete(uuidVal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
