package api

import (
	"net/http"
	"sync"
	"time"

	"github.com/formych/dota/dao"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	chipTypeMu   sync.RWMutex
	chipTypeList []*ChipTypeConf
)

// ChipList ...
func ChipList(c *gin.Context) {
	if chipTypeList == nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "成功", "data": "[]"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "成功", "data": chipTypeList})
}

// ChipAdd ...
func ChipAdd(c *gin.Context) {
	chipConf := &ChipTypeConf{}
	err := c.ShouldBind(chipConf)
	if err != nil {
		logrus.Errorf("Bind  data failed, err:[%s]", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "添加失败", "data": map[string]string{"error": err.Error()}})
		return
	}
	now := time.Now()
	chip := &dao.ChipType{
		Name:      chipConf.Name,
		Desc:      chipConf.Desc,
		CreatedAt: now,
		UpdateAt:  now,
		Status:    1,
	}
	_, err = dao.ChipTypeDAO.Add(chip)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "数据库错误", "data": map[string]string{"error": err.Error()}})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "添加成功", "data": chipConf})
}

// ChipTypeConf ...
type ChipTypeConf struct {
	ID   int64  `json:"id"`
	Name string `json:"name" form:"name" binding:"required"`
	Desc string `json:"desc" form:"desc" binding:"required"`
}

// ChipTypeConfs ...
var ChipTypeConfs = []*ChipTypeConf{}

func init() {
	LoadChipTypeConf()
}

// LoadChipTypeConf 加载
func LoadChipTypeConf() {
	chips, err := dao.ChipTypeDAO.GetAll()
	if err != nil {
		logrus.Errorf("load chip type failed, err:[%s]", err.Error())
		return
	}
	var tmpChipTypeList []*ChipTypeConf
	for _, v := range chips {
		tmpChipTypeList = append(tmpChipTypeList, &ChipTypeConf{ID: v.ID, Name: v.Name})
	}
	chipTypeMu.Lock()
	chipTypeList = tmpChipTypeList
	chipTypeMu.Unlock()
	return
}
