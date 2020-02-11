package infra

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"AY1st/util"

	"github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

const (
	databaseConnectionMaxLifeTimeEnv        = "DATABASE_CONNECTION_MAX_LIFETIME"
	databaseConnectionMaxLifeTimeDefaultSec = 1
)

// LoadMySQLWriterConfigEnv initializes MySQL writer config using Environment Variables.
func LoadMySQLWriterConfigEnv() *mysql.Config {

	addr := getAddr(os.Getenv("DATABASE_WRITE_HOST"), os.Getenv("DATABASE_WRITE_PORT"))

	conf := &mysql.Config{
		Net:                  "tcp",
		Addr:                 addr,
		DBName:               os.Getenv("DATABASE_NAME"),
		User:                 os.Getenv("DATABASE_USER"),
		Passwd:               os.Getenv("DATABASE_PASSWORD"),
		AllowNativePasswords: true,
		// TODO: DBのタイムゾーンを決める
		Loc: time.UTC,
		// Aurora の Cluster エンドポイントのフェイルオーバー時にRead-Replicaに接続してしまう場合の対策
		RejectReadOnly: true,
	}
	return conf
}

// LoadMySQLReaderConfigEnv initializes MySQL reader config using Environment Variables.
func LoadMySQLReaderConfigEnv() *mysql.Config {

	addr := getAddr(os.Getenv("DATABASE_READ_HOST"), os.Getenv("DATABASE_READ_PORT"))

	conf := &mysql.Config{
		Net:                  "tcp",
		Addr:                 addr,
		DBName:               os.Getenv("DATABASE_NAME"),
		User:                 os.Getenv("DATABASE_USER"),
		Passwd:               os.Getenv("DATABASE_PASSWORD"),
		AllowNativePasswords: true,
		// TODO: DBのタイムゾーンを決める
		Loc: time.UTC,
	}
	return conf
}

// InitMySQLEngine initialize xorm engine for mysql
func InitMySQLEngine(conf *mysql.Config) (*xorm.Engine, error) {
	engine, err := xorm.NewEngine("mysql", conf.FormatDSN())
	if err != nil {
		return nil, err
	}

	charset, ok := conf.Params["charset"]
	if !ok {
		charset = "utf8mb4"
	}

	// DBのタイムゾーンは日本時間
	engine.SetTZDatabase(util.TokyoTimeLocation)
	engine.SetTZLocation(util.TokyoTimeLocation)
	engine.Charset(charset)
	engine.SetMapper(core.GonicMapper{})
	engine.StoreEngine("InnoDb")
	showSQL := os.Getenv("SHOW_SQL")
	if showSQL == "0" || showSQL == "false" {
		engine.ShowSQL(false)
	} else {
		engine.ShowSQL(true)
	}

	var connMaxLifeTime int
	connMaxLifeTimeStr := os.Getenv(databaseConnectionMaxLifeTimeEnv)
	connMaxLifeTime, err = strconv.Atoi(connMaxLifeTimeStr)
	if err != nil {
		connMaxLifeTime = databaseConnectionMaxLifeTimeDefaultSec
		logger := util.GetLogger()
		logger.Infof("%v expects int value, but %v was given.", databaseConnectionMaxLifeTimeEnv, connMaxLifeTimeStr)
		logger.Infof("use default %v [sec] for %v", databaseConnectionMaxLifeTimeDefaultSec, databaseConnectionMaxLifeTimeEnv)
	}
	engine.SetConnMaxLifetime(time.Duration(connMaxLifeTime) * time.Second)

	logLevel, err := parseLogLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		return nil, err
	}
	engine.SetLogLevel(logLevel)

	return engine, nil
}

func getAddr(host, port string) string {
	var addr string
	if strings.Contains(host, ":") {
		addr = host
	} else {
		pInt := 3306
		if p, err := strconv.Atoi(port); err == nil {
			pInt = p
		}
		addr = fmt.Sprintf("%v:%v", host, pInt)
	}
	return addr
}

// parseLogLevel parses level string into xorm's LogLevel
func parseLogLevel(lvl string) (core.LogLevel, error) {
	switch strings.ToLower(lvl) {
	case "panic", "fatal", "error":
		return core.LOG_ERR, nil
	case "warn", "warning":
		return core.LOG_WARNING, nil
	case "info":
		return core.LOG_INFO, nil
	case "debug":
		return core.LOG_DEBUG, nil
	}
	return core.LOG_DEBUG, fmt.Errorf("cannot parse \"%v\" into go-xorm/core.LogLevel", lvl)
}

var escapeReplace = []struct {
	Key      string
	Replaced string
}{
	{"\\", "\\\\"},
	{`'`, `\'`},
	{"\\0", "\\\\0"},
	{"\n", "\\n"},
	{"\r", "\\r"},
	{`"`, `\"`},
	{"\x1a", "\\Z"},
}

// EscapeMySQLString prevents from SQL-injection.
func EscapeMySQLString(value string) string {
	for _, r := range escapeReplace {
		value = strings.Replace(value, r.Key, r.Replaced, -1)
	}

	return value
}
