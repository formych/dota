package config

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql" //...
	"github.com/go-yaml/yaml"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // ...
	"github.com/sirupsen/logrus"
)

// DB enum
const (
	MySQL = iota
	PgSQL
)

// Global
var (
	DB          *sqlx.DB
	RedisClient *redis.Client
)

// Log ...
type Log struct {
	Level      string `yaml:"level"`
	LineLevel  string `yaml:"line_level"`
	AsyncCount int64  `yaml:"async_count"`
}

// 目前简单实现，后面再改为数组
// 分布式用到
// Key          string `yaml:"key"`
// Mode         string `yaml:"mode"`

// ConnConf ...
type ConnConf struct {
	Host         string `yaml:"host"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Port         int    `yaml:"port"`
	Database     string `yaml:"database"`
	ConnTimeout  int    `yaml:"conn_timeout"`
	ReadTimeout  int    `yaml:"read_timeout"`
	WriteTimeout int    `yaml:"write_timeout"`
}

// C 通用配置
var C = struct {
	Port        string    `yaml:"port"`
	WaitTimeout int       `yaml:"wait_timeout"`
	Log         Log       `yaml:"log"`
	MySQLConf   *ConnConf `yaml:"mysql"`
	PgSQLConf   *ConnConf `yaml:"pgsql"`
	RedisConf   *ConnConf `yaml:"redis"`
}{}

var c = flag.String("c", "", "config file path")

func init() {
	logrus.Info("Init server automatically")
	flag.Parse()
	data, err := ioutil.ReadFile(*c)
	if err != nil {
		logrus.Errorf("read config file failed, file:[%s], err:[%s]", *c, err.Error())
		os.Exit(-1)
	}
	err = yaml.Unmarshal([]byte(data), &C)
	if err != nil {
		logrus.Errorf("unmarshal config file failed, data:[%s], err:[%s]", data, err.Error())
		os.Exit(-1)
	}

	if C.MySQLConf != nil {
		logrus.Infof("mysql config:[%+v]", *C.MySQLConf)
		Connect(MySQL, C.MySQLConf)
	}
	if C.PgSQLConf != nil {
		logrus.Infof("pgsql config:[%+v]", *C.PgSQLConf)
		Connect(PgSQL, C.PgSQLConf)
	}
	if C.RedisConf != nil {
		logrus.Infof("redis config:[%+v]", *C.RedisConf)
		InitRedis(C.RedisConf)
	}
}

// Connect ...
func Connect(driverType int, option *ConnConf) {
	var (
		driverName, dsn string
		err             error
	)
	if driverType == MySQL {
		driverName, dsn = InitMysql(option)
	} else if driverType == PgSQL {
		driverName, dsn = InitPgSQL(option)
	} else {
		logrus.Errorf("unsport db driver, type:[%d]", driverType)
		return
	}

	DB, err = sqlx.Open(driverName, dsn)
	if err != nil {
		logrus.Errorf("Open db failed, err:[%s]", err.Error())
		os.Exit(-1)
	}
}

// InitMysql ...
func InitMysql(conn *ConnConf) (driverName, dsn string) {
	driverName = "mysql"
	timeoutArgs := fmt.Sprintf("timeout=%dms&readTimeout=%dms&writeTimeout=%dms", conn.ConnTimeout, conn.ReadTimeout, conn.WriteTimeout)
	dsn = conn.User + ":" + conn.Password + "@tcp(" + conn.Host + ":" + strconv.Itoa(conn.Port) + ")/" + conn.Database + "?charset=utf8&parseTime=True&loc=Local&" + timeoutArgs
	return
}

// InitPgSQL ...
func InitPgSQL(conn *ConnConf) (driverName, dsn string) {
	driverName = "postgres"
	dsn = "postgres://" + conn.User + ":" + conn.Password + "@" + conn.Host + ":" + strconv.Itoa(conn.Port) + "/" + conn.Database + "?sslmode=disable"
	return
}

// InitRedis ...
func InitRedis(conn *ConnConf) {
	conf := &redis.Options{
		Addr:         fmt.Sprintf("%s:%d", conn.Host, conn.Port),
		Password:     conn.Password,
		DialTimeout:  time.Duration(conn.ConnTimeout) * time.Millisecond,
		ReadTimeout:  time.Duration(conn.ReadTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(conn.WriteTimeout) * time.Millisecond,
	}
	RedisClient = redis.NewClient(conf)
}

// Close 后续优化
func Close() {
	if err := DB.Close(); err != nil {
		logrus.Errorf("clode db failed, err:[%s]", err.Error())
	}
	if err := RedisClient.Close(); err != nil {
		logrus.Errorf("clode redis failed, err:[%s]", err.Error())
	}
}

// 暂时没想到好的日志处理方式, 后面优化
// func initLogConf() {
// 	// default use info level
// 	switch C.Log.Level {
// 	case "debug":
// 		logrus.SetLevel(logrus.DebugLevel)
// 	case "warn":
// 		logrus.SetLevel(logrus.WarnLevel)
// 	case "error":
// 		logrus.SetLevel(logrus.ErrorLevel)
// 	case "fatal":
// 		logrus.SetLevel(logrus.FatalLevel)
// 	case "panic":
// 		logrus.SetLevel(logrus.PanicLevel)
// 	default:
// 		logrus.SetLevel(logrus.InfoLevel)
// 	}
// }
