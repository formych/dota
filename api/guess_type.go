package api

import (
	"net/http"
	"time"

	"github.com/formych/dota/dao"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// GuessTypeAdd 增加结算类型
func GuessTypeAdd(c *gin.Context) {
	g := &dao.GuessType{}
	err := c.ShouldBind(g)
	if err != nil {
		logrus.Errorf("Bind  data failed, err:[%s]", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "添加失败", "data": map[string]string{"error": err.Error()}})
		return
	}
	now := time.Now()
	settleInfo := &dao.GuessType{
		Name:      g.Name,
		Desc:      g.Desc,
		CreatedAt: now,
		UpdateAt:  now,
		Status:    1,
	}
	err = dao.GuessTypeDAO.Add(settleInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "数据库错误", "data": map[string]string{"error": err.Error()}})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "添加成功", "data": settleInfo})
}

// // GuessDelete 删除竞猜
// func GuessDelete(c *gin.Context) {
// 	uid := int64(666666)
// 	gidStr := c.Param("gid")
// 	gid, err := strconv.ParseInt(gidStr, 10, 64)
// 	if err != nil {
// 		logrus.Errorf("parse gid failed, err:[%s]", err.Error())
// 		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "更新失败", "data": map[string]string{"error": err.Error()}})
// 		return
// 	}

// 	err = dao.GuessInfoDAO.Delete(gid, uid)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "数据库错误", "data": map[string]string{"error": err.Error()}})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "删除成功"})
// }

// // GuessUpdate 更新竞猜
// func GuessUpdate(c *gin.Context) {
// 	g := &GuessInfo{}
// 	err := c.ShouldBind(g)
// 	if err != nil {
// 		logrus.Errorf("Bind  data failed, err:[%s]", err.Error())
// 		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "添加失败", "data": map[string]string{"error": err.Error()}})
// 		return
// 	}
// 	now := time.Now()
// 	guessInfo := &dao.GuessInfo{
// 		ID:           g.ID,
// 		UID:          g.UID,
// 		SettleTypeID: g.SettleTypeID,
// 		GuessTypeID:  g.GuessTypeID,
// 		ChipTypeID:   g.ChipTypeID,
// 		Result:       g.Result,
// 		FundPool:     g.FundPool,
// 		Desc:         g.Desc,
// 		UpdateAt:     now,
// 		Status:       1,
// 	}
// 	err = dao.GuessInfoDAO.Update(guessInfo)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "数据库错误", "data": map[string]string{"error": err.Error()}})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "更新成功", "data": guessInfo})
// }

// // GuessListSub ...
// func GuessListSub(c *gin.Context) {
// 	startStr := c.PostForm("start")
// 	limitStr := c.PostForm("limit")
// 	start, serr := strconv.ParseInt(startStr, 10, 64)
// 	limit, lerr := strconv.ParseInt(limitStr, 10, 64)
// 	if lerr != nil || serr != nil {
// 		logrus.Errorf("parse failed, lerr:[%s], serr:[%s]", lerr, serr)
// 		return
// 	}
// 	list, err := dao.GuessInfoDAO.GetSub(start, limit)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "数据库错误", "data": map[string]string{"error": err.Error()}})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "删除成功", "data": list})
// }

// GuessTypeList ...
func GuessTypeList(c *gin.Context) {
	list, err := dao.SettleTypeDAO.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "数据库错误", "data": map[string]string{"error": err.Error()}})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "获取成功", "data": list})
}
