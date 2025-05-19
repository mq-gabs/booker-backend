package user

import (
	"booker/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type UserController struct {
	Service *UserService
}

func SetRoutes(r *gin.Engine) {
	repo := NewUserMemoryRepository()
	service := NewUserService(repo)

	uc := &UserController{Service: service}

	g := r.Group("/users")

	g.POST("/", uc.Create)
	g.GET("/", uc.List)
	g.GET("/:id", uc.FindOne)
	g.PUT("/:id", uc.Update)
	g.DELETE("/:id", uc.Delete)
}

func (uc *UserController) Create(c *gin.Context) {
	var dto CreateUserDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := uc.Service.Create(dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

func (uc *UserController) List(c *gin.Context) {
	query := utils.QueryOptions{
		Page:     utils.ParseQueryInt(c, "page", 0),
		PageSize: utils.ParseQueryInt(c, "pageSize", 10),
	}
	resp, err := uc.Service.List(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (uc *UserController) FindOne(c *gin.Context) {
	id := c.Param("id")
	uuidVal, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	user, err := uc.Service.FindOne(uuidVal)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (uc *UserController) Update(c *gin.Context) {
	id := c.Param("id")
	uuidVal, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	var dto UpdateUserDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := uc.Service.Update(uuidVal, dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (uc *UserController) Delete(c *gin.Context) {
	id := c.Param("id")
	uuidVal, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	if err := uc.Service.Delete(uuidVal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
