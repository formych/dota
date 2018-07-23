package router

import (
	"github.com/formych/dota/api"

	"github.com/gin-gonic/gin"
)

// Mu ...

// Run 启动服务器和路由注册
func Run(port string) {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	user := r.Group("/user")
	{
		user.POST("/signup", api.SignUp)
		user.POST("/signin", api.SignIn)
	}
	r.Run(":" + port)
}
