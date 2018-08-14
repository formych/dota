package api

import (
	"net/http"
	"time"

	"github.com/formych/dota/dao"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// GuessRecord 竞猜信息
type GuessRecord struct {
	ID          int64     `json:"id omitempty" form:"id"`
	UID         int64     `json:"uid" form:"uid" binding:"required"`
	GID         int64     `json:"gid" form:"gid" binding:"required"`
	Amount      int64     `json:"Amount" form:"Amount" binding:"required"`
	GuessTypeID int64     `json:"guess_type_id" form:"guess_type_id" binding:"required"`
	ChipTypeID  int64     `json:"chip_type_id" form:"chip_type_id" binding:"required"`
	Result      int64     `json:"result" form:"result" binding:"required"`
	Earnings    int64     `json:"earnings" form:"earnings" binding:"required"`
	Odds        int64     `json:"odds" form:"odds" binding:"required"`
	CreatedAt   time.Time `json:"created_at" form:"created_at"`
	UpdateAt    time.Time `json:"updated_at" form:"updated_at"`
	Status      int8      `json:"status" form:"updated_at"`
}

// GuessRecordConf 参数
type GuessRecordConf struct {
	Num  int64 `json:"num" form:"num" binding:"required"`
	Page int64 `json:"page" form:"page" binding:"required"`
}

// GuessRecordAdd 增加竞猜
func GuessRecordAdd(c *gin.Context) {
	g := &GuessRecord{}
	err := c.ShouldBind(g)
	if err != nil {
		logrus.Errorf("Bind  data failed, err:[%s]", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "添加失败", "data": map[string]string{"error": err.Error()}})
		return
	}
	uid := c.GetInt64("uid")
	now := time.Now()
	guessRecord := &dao.GuessRecord{
		UID:       uid,
		GID:       g.GID,
		Amount:    g.Amount,
		Result:    g.Result,
		Earnings:  g.Earnings,
		Odds:      g.Odds,
		CreatedAt: now,
		UpdateAt:  now,
		Status:    1,
	}
	err = dao.GuessRecordDAO.Add(guessRecord)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "数据库错误", "data": map[string]string{"error": err.Error()}})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "添加成功", "data": guessRecord})
}

// GuessRecordDelete 删除竞猜记录
// func GuessRecordDelete(c *gin.Context) {
// 	uid := c.GetInt64("uid")
// 	gidStr := c.Param("gid")
// 	gid, err := strconv.ParseInt(gidStr, 10, 64)
// 	if err != nil {
// 		logrus.Errorf("parse gid failed, err:[%s]", err.Error())
// 		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "更新失败", "data": map[string]string{"error": err.Error()}})
// 		return
// 	}

// 	err = dao.GuessRecordDAO.Delete(gid, uid)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "数据库错误", "data": map[string]string{"error": err.Error()}})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "删除成功"})
// }

// GuessUpdate 更新竞猜
// func GuessUpdate(c *gin.Context) {
// 	g := &GuessRecord{}
// 	err := c.ShouldBind(g)
// 	if err != nil {
// 		logrus.Errorf("Bind  data failed, err:[%s]", err.Error())
// 		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "添加失败", "data": map[string]string{"error": err.Error()}})
// 		return
// 	}
// 	now := time.Now()
// 	GuessRecord := &dao.GuessRecord{
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
// 	err = dao.GuessRecordDAO.Update(GuessRecord)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "数据库错误", "data": map[string]string{"error": err.Error()}})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "更新成功", "data": GuessRecord})
// }

// GuessListSub ...
// func GuessListSub(c *gin.Context) {
// 	startStr := c.PostForm("start")
// 	limitStr := c.PostForm("limit")
// 	start, serr := strconv.ParseInt(startStr, 10, 64)
// 	limit, lerr := strconv.ParseInt(limitStr, 10, 64)
// 	if lerr != nil || serr != nil {
// 		logrus.Errorf("parse failed, lerr:[%s], serr:[%s]", lerr, serr)
// 		return
// 	}
// 	list, err := dao.GuessRecordDAO.GetSub(start, limit)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "数据库错误", "data": map[string]string{"error": err.Error()}})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "删除成功", "data": list})
// }

// GuessRecordList ...
func GuessRecordList(c *gin.Context) {
	list, err := dao.GuessRecordDAO.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "数据库错误", "data": map[string]string{"error": err.Error()}})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "删除成功", "data": list})
}
