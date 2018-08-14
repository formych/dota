package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/formych/dota/dao"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// GuessListConf 参数
type GuessListConf struct {
	Num  int64 `json:"num" form:"num" binding:"required"`
	Page int64 `json:"page" form:"page" binding:"required"`
}

// GuessInfo 竞猜信息表
type GuessInfo struct {
	ID           int64     `json:"id,omitempty" form:"id" db:"id"`
	UID          int64     `json:"uid" form:"uid" db:"uid"`
	Name         string    `json:"name" form:"name" db:"name" `
	SettleTypeID int64     `json:"settle_type_id" form:"settle_type_id" db:"settle_type_id" `
	GuessTypeID  int64     `json:"guess_type_id" form:"guess_type_id" db:"chip_type_id" `
	ChipTypeID   int64     `json:"chip_type_id" form:"chip_type_id" db:"chip_type_id" `
	Result       int64     `json:"result" form:"result" db:"result" `
	FundPool     int64     `json:"fund_pool" form:"fund_pool" db:"fund_pool"`
	Info         Info      `json:"Info" form:"Info" db:"Info" `
	StartTime    string    `json:"start_time" form:"start_time" db:"start_time" `
	EndTime      string    `json:"end_time" form:"end_time" db:"end_time" `
	CreatedAt    time.Time `json:"created_at" form:"created_at" db:"created_at"`
	UpdateAt     time.Time `json:"updated_at" form:"updated_at" db:"updated_at"`
	Status       int8      `json:"status" form:"updated_at" db:"status"`
}

// Option ...
type Option struct {
	Name  string `json:"name" form:"name" db:"name"`
	Value int    `json:"value" form:"value" db:"name"`
}

// Info ...
type Info struct {
	Teams   []int64
	Options []*Option
}

// GuessAdd 增加竞猜
func GuessAdd(c *gin.Context) {
	g := &GuessInfo{}
	err := c.ShouldBind(g)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "msg": "添加失败", "data": map[string]string{"error": err.Error()}})
		return
	}

	start, err1 := time.ParseInLocation("2006-01-02T15:04", g.StartTime, time.Local)
	end, err2 := time.ParseInLocation("2006-01-02T15:04", g.EndTime, time.Local)
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "msg": "时间格式错误", "data": map[string]string{"error": err1.Error()}})
		return
	}
	if start.After(end) {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "msg": "添加失败", "data": map[string]string{"error": "开始时间不能小于截止时间"}})
		return
	}
	// 后续对info立面的数据做校验
	infoBytes, err := json.Marshal(g.Info)
	now := time.Now()
	guessInfo := &dao.GuessInfo{
		UID:          int64(c.GetFloat64("uid")),
		Name:         g.Name,
		SettleTypeID: g.SettleTypeID,
		GuessTypeID:  g.GuessTypeID,
		ChipTypeID:   g.ChipTypeID,
		CreatedAt:    now,
		UpdateAt:     now,
		Info:         string(infoBytes),
		StartTime:    start,
		EndTime:      end,
		FundPool:     0,
		Status:       1,
	}
	fmt.Printf("%+v", guessInfo)
	_, err = dao.GuessInfoDAO.Add(guessInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "数据库错误", "data": map[string]string{"error": err.Error()}})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "添加成功", "data": guessInfo})
}

// GuessDelete 删除竞猜
func GuessDelete(c *gin.Context) {
	uid := int64(666666)
	gidStr := c.Param("gid")
	gid, err := strconv.ParseInt(gidStr, 10, 64)
	if err != nil {
		logrus.Errorf("parse gid failed, err:[%s]", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "更新失败", "data": map[string]string{"error": err.Error()}})
		return
	}

	err = dao.GuessInfoDAO.Delete(gid, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "数据库错误", "data": map[string]string{"error": err.Error()}})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "删除成功"})
}

// GuessUpdate 更新竞猜
// func GuessUpdate(c *gin.Context) {
// 	g := &dao.GuessInfo{}
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

// // GuessList ...
// func GuessList(c *gin.Context) {
// 	list, err := dao.GuessInfoDAO.GetAll()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "数据库错误", "data": map[string]string{"error": err.Error()}})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "删除成功", "data": list})
// }
