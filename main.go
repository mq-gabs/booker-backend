package main

import (
	"booker/modules/scheduling"
	"booker/modules/schedulingprofile"
	"booker/modules/user"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	setRoutes(r)

	r.Run(":5000")
}

func setRoutes(r *gin.Engine) {
	user.SetRoutes(r)
	scheduling.SetRoutes(r)
	schedulingprofile.SetRoutes(r)
}
