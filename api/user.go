package api

import (
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/formych/dota/config"
	"github.com/formych/dota/dao"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

var (
	hmacSampleSecret = []byte("hello world")
	userToken        = "user:token"
)
var authTypeMap = map[int8]bool{
	1: true,
}

// CodeMap 返回码映射表
var CodeMap = map[int32]string{
	2000001: "token为空",
	2000002: "wrong token",
	2000003: "Invalid token1",
	2000004: "Invalid token2",
}

// UserAuth ...
type UserAuth struct {
	ID          int64     `form:"id" json:"id,omitempty"`
	AuthType    int8      `form:"auth_type" json:"auth_type"`
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
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "msg": "参数不合法", "data": map[string]string{"error": err.Error()}})
		return
	}
	if !authTypeMap[userAuth.AuthType] {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "msg": "参数不合法", "data": map[string]string{"error": "invalid auth_type"}})
		return
	}
	reg := regexp.MustCompile(`^1[3456789]\d{9}$`)
	if !reg.MatchString(userAuth.Identifier) {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "msg": "参数不合法", "data": map[string]string{"error": "手机号码不合法"}})
		return
	}

	dbuser, err := dao.UserAuthDAO.Get(userAuth.Identifier)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "msg": "网络错误", "data": map[string]string{"error": err.Error()}})
		return
	}
	if dbuser != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "msg": "用户已存在"})
		return
	}
	passHash, err := GenPassword(userAuth.Certificate)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "msg": "网络错误", "data": map[string]string{"error": err.Error()}})
		return
	}
	now := time.Now()
	newUser := &dao.UserAuth{
		AuthType:    userAuth.AuthType,
		Identifier:  userAuth.Identifier,
		Certificate: passHash,
		CreatedAt:   now,
		UpdateAt:    now,
		Status:      1,
	}
	id, err := dao.UserAuthDAO.Add(newUser)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "msg": "网络错误", "data": map[string]string{"error": err.Error()}})
		return
	}
	newUser.ID = id
	token, err := GetToken(newUser)
	if err != nil {
		token = ""
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "生成token失败", "data": map[string]string{"error": err.Error()}})
		return
	}
	_, err = config.RedisClient.HSet(userToken, userAuth.Identifier, 1).Result()
	if err != nil {
		logrus.Errorf("set redis failed err:[%s]", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "msg": "网络错误", "data": map[string]string{"err": err.Error()}})
		return
	}
	c.SetCookie("sid", token, 7*24*3600, "", "", false, true)
	data := map[string]interface{}{
		"user_info": newUser,
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "注册成功", "data": data})
}

// SignIn 登录
func SignIn(c *gin.Context) {
	u := &UserAuth{}
	err := c.ShouldBind(u)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "msg": "登录失败", "data": map[string]string{"error": err.Error()}})
		return
	}
	dbuser, err := dao.UserAuthDAO.Get(u.Identifier)
	if err != nil {
		logrus.Errorf("get user_auth data failed, err:[%s]", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "msg": "网络错误", "data": map[string]string{"error": err.Error()}})
		return
	}
	if dbuser == nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "msg": "用户不存在"})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(dbuser.Certificate), []byte(u.Certificate))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "msg": "密码不正确"})
		return
	}
	fmt.Printf("%+v", dbuser)
	token, err := GetToken(dbuser)
	if err != nil {
		token = ""
	}
	_, err = config.RedisClient.HSet(userToken, dbuser.Identifier, 1).Result()
	if err != nil {
		logrus.Errorf("set redis failed err:[%s]", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "msg": "网络错误", "data": map[string]string{"err": err.Error()}})
		return
	}
	data := map[string]interface{}{
		"user_info": dbuser,
	}
	c.SetCookie("sid", token, 7*24*3600, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "登录成功", "data": data})
}

// Authentication 认证
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, err := c.Cookie("sid")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 2000004, "msg": CodeMap[2000004], "data": err.Error()})
			c.Abort()
			return
		}
		claims, code := DecodeToken(tokenStr)
		if code != 0 {
			c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": CodeMap[code]})
			c.Abort()
			return
		}
		// 做数据验证
		// 暂时采取用re	dis控制单一token
		_, err = config.RedisClient.HGet(userToken, fmt.Sprintf("%v", claims["identifier"])).Result()
		if err != nil {
			logrus.Errorf("get redis failed err:[%s]", err.Error())
			c.JSON(http.StatusOK, gin.H{"code": 2000004, "msg": CodeMap[2000004]})
			c.Abort()
			return
		}
		identifier := claims["identifier"]
		c.Set("uid", claims["uid"])
		c.Set("identifier", identifier)
		c.Next()
	}
}

// GetToken 生成一个简单的token, 后续优化
func GetToken(u *dao.UserAuth) (tokenStr string, err error) {
	fmt.Println(*u)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":        u.ID,
		"identifier": u.Identifier,
		"nbf":        time.Date(2018, 01, 01, 00, 0, 0, 0, time.UTC).Unix(),
		"exp":        time.Now().Add(7 * 24 * time.Hour).Unix(),
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
	fmt.Println(claims)
	return nil, 2000003
}

// ResetPassword ...
func ResetPassword(c *gin.Context) {
	uid := c.GetInt64("uid")
	newPassword := c.PostForm("certificate")
	passHash, err := GenPassword(newPassword)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "msg": "网络错误", "data": map[string]string{"error": err.Error()}})
		return
	}
	err = dao.UserAuthDAO.UpdatePassword(uid, passHash)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "msg": "网络错误", "data": map[string]string{"error": err.Error()}})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "msg": "更新成功", "data": map[string]string{"error": err.Error()}})
}

// GenPassword ...
func GenPassword(certificate string) (passHash string, err error) {
	hashByte, err := bcrypt.GenerateFromPassword([]byte(certificate), bcrypt.DefaultCost)
	if err != nil {
		logrus.Errorf("hash user certificate failed, certificate:[%s], err:[%s]", certificate, err.Error())
		return
	}
	passHash = string(hashByte)
	return
}
