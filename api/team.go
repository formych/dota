package api

import (
	"net/http"

	"github.com/formych/dota/dao"
	"github.com/gin-gonic/gin"
)

// TeamList _
func TeamList(c *gin.Context) {
	teams, err := dao.TeamDAO.GetAll()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "获取失败", "data": map[string]string{"error": err.Error()}})
		return
	}
	if len(teams) == 0 {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "获取成功", "data": "[]"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "获取成功", "data": teams})
}

// TeamAdd _
// func TeamAdd(c *gin.Context) {
// 	team := &dao.Team{}

// 	s, err := dao.TeamDAO.GetAll()
// 	if err != nil {
// 		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "获取失败", "data": map[string]string{"error": err.Error()}})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "获取成功", "data": teams})

// }
