package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/formych/dota/config"
	"github.com/formych/dota/dao"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	hmacSampleSecret = []byte("hello world")
	userToken        = "user:token:%s"
)

// UserAuth ...
type UserAuth struct {
	ID          int64     `form:"id" json:"id"`
	AuthType    int8      `form:"auth_type" json:"auth_type" binding:"required"`
	Identifier  string    `form:"identifier" json:"identifier" binding:"required"`
	Certificate string    `form:"certificate" json:"certificate" binding:"required"`
	CreatedAt   time.Time `form:"created_at" json:"created_at"`
	UpdateAt    time.Time `form:"updated_at" json:"updated_at"`
	Status      int8      `form:"status" json:"status"`
}

// SignUp 注册
func SignUp(c *gin.Context) {
	var userAuth = &UserAuth{}
	if err := c.ShouldBind(userAuth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dbuser, err := dao.UserAuthDAO.Get(userAuth.Identifier)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "error": err.Error()})
		return
	}
	if dbuser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "用户已存在"})
		return
	}
	newUser := &dao.UserAuth{
		AuthType:    userAuth.AuthType,
		Identifier:  userAuth.Identifier,
		Certificate: userAuth.Certificate,
		CreatedAt:   time.Now(),
		UpdateAt:    time.Now(),
		Status:      1,
	}
	if _, err = dao.UserAuthDAO.Add(newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "数据库错误", "data": map[string]string{"error": err.Error()}})
		return
	}
	token, err := GetToken(userAuth)
	if err != nil {
		token = ""
	}

	data := map[string]interface{}{
		"token":      token,
		"identifier": newUser.Identifier,
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "注册成功", "data": data})
}

// SignIn 登录
func SignIn(c *gin.Context) {
	u := &UserAuth{}
	err := c.ShouldBind(u)
	if err != nil {
		logrus.Errorf("Bind user data failed, err:[%s]", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "登录失败", "data": map[string]string{"error": err.Error()}})
		return
	}
	dbuser, err := dao.UserAuthDAO.Get(u.Identifier)
	if err != nil {
		logrus.Errorf("get user data failed, err:[%s]", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "登录失败", "data": map[string]string{"error": err.Error()}})
		return
	}
	if dbuser == nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "用户不存在"})
		return
	}
	token, err := GetToken(u)
	if err != nil {
		token = ""
	}
	data := map[string]interface{}{
		"user":  dbuser,
		"token": token,
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "登录成功", "data": data})
}

// Authentication 认证
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.PostForm("token")
		claims, code := DecodeToken(tokenStr)
		if code != 0 {
			c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": CodeMap[code]})
			c.Abort()
			return
		}
		// 做数据验证
		// 暂时采取用redis控制单一token
		_, err := config.RedisClient.Get(fmt.Sprintf(userToken, claims["identifier"])).Result()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": CodeMap[2000004]})
			c.Abort()
			return
		}
		identifier := claims["identifier"]
		c.Set("identifier", identifier)
		c.Next()
	}
}

// GetToken 生成一个简单的token, 后续优化
func GetToken(u *UserAuth) (tokenStr string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   u.ID,
		"name": u.Identifier,
		"nbf":  time.Date(2018, 01, 01, 00, 0, 0, 0, time.UTC).Unix(),
		"exp":  time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenStr, err = token.SignedString(hmacSampleSecret)
	if err != nil {
		logrus.Errorf("signing token string failed, err:[%s]", err.Error())
	}
	return
}

// DecodeToken 解析token
func DecodeToken(tokenStr string) (claims jwt.MapClaims, code int32) {
	if tokenStr == "" {
		return nil, 2000001
	}
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSampleSecret, nil
	})
	if err != nil {
		logrus.Errorf("Invalid token string, err:[%s]", err.Error())
		return nil, 200002
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, 0
	}
	return nil, 2000003
}

// CodeMap 返回码映射表
var CodeMap = map[int32]string{
	2000001: "token为空",
	2000002: "wrong token",
	2000003: "Invalid token",
	2000004: "Invalid token",
}
