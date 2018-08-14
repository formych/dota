package router

import (
	"github.com/formych/dota/api"

	"github.com/gin-gonic/gin"
)

// Run 启动服务器和路由注册
func Run(port string) {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	user := r.Group("/user")
	{
		user.POST("/signup", api.SignUp)
		user.POST("/signin", api.SignIn)
		user.POST("/reset", api.ResetPassword)
	}
	v1 := r.Group("/v1")
	v1.Use(api.Authentication())

	guess := v1.Group("/guess")
	{
		// guess.POST("/list", api.GuessList)
		// guess.POST("/info", api.GuessListSub)
		guess.POST("/add", api.GuessAdd)
		// guess.DELETE("/", api.GuessDelete)
	}

	chip := v1.Group("/chip")
	{
		chip.POST("/list", api.ChipList)
		chip.POST("/add", api.ChipAdd)
	}

	settle := v1.Group("/settle")
	{
		settle.POST("/add", api.SettleTypeAdd)
		settle.POST("/list", api.SettleList)
	}

	balance := v1.Group("/balance")
	{
		balance.POST("/list", api.BalanceDetailList)
		// balance.POST("/info", api.BalanceInfo)
	}
	team := v1.Group("/team")
	{
		team.GET("/list", api.TeamList)
	}
	r.Run(":" + port)
}
