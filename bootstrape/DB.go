package bootstrape

import (
	"context"
	"gin-mall/app/dao"
	"gin-mall/global"
	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

/**
 * @author: biao
 * @date: 2025/9/1 20:18
 * @code: 彼方尚有荣光在
 * @description: 初始化数据库
 */

func InitializeDB() *gorm.DB {
	switch global.App.Config.Mysql.DB {
	case "mysql":
		return initMySqlGorm()
	default:
		return initMySqlGorm()
	}
}

func initMySqlGorm() *gorm.DB {
	dbConfig := global.App.Config.Mysql

	if dbConfig.DB == "" {
		return nil
	}

	// 实现数据库的读写分离
	// 读数据库
	pathRead := strings.Join([]string{dbConfig.DbUser, ":", dbConfig.DbPassword, "@tcp(", dbConfig.DbHost, ":", dbConfig.DbPort, ")/", dbConfig.DbName,
		"?charset=utf8mb4&parseTime=True&loc=Local"}, "")
	// 写数据库
	pathWrite := strings.Join([]string{dbConfig.DbUser, ":", dbConfig.DbPassword, "@tcp(", dbConfig.DbHost, ":", dbConfig.DbPort, ")/", dbConfig.DbName,
		"?charset=utf8mb4&parseTime=True&loc=Local"}, "")

	mysqlConfig := mysql.Config{
		DSN:                       pathRead,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,  //禁用 datetime 精度， MySQL 5.6 之前不支持
		DontSupportRenameIndex:    true,  //重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, //// 根据版本自动配置
	}

	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		Logger: getGormLogger(),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		global.App.Log.Error("mysql connect failed; err: ", zap.Any("err", err))
		return nil
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Second * 30)

	// 主从配置
	_ = db.Use(dbresolver.Register(
		dbresolver.Config{
			Sources:  []gorm.Dialector{mysql.Open(pathWrite)},                      //写操作
			Replicas: []gorm.Dialector{mysql.Open(pathRead), mysql.Open(pathRead)}, //读操作
			Policy:   dbresolver.RandomPolicy{},
		}))
	dao.Migration(db)
	return db
}

// 自定义 gorm Writer
func getGormLogWriter() logger.Writer {
	var writer io.Writer

	// 是否启用日志文件
	if global.App.Config.Mysql.EnableFileLogWriter {
		// 自定义 Writer
		writer = &lumberjack.Logger{
			Filename:   global.App.Config.Log.RootDir + "/" + global.App.Config.Mysql.LogFilename,
			MaxSize:    global.App.Config.Log.MaxSize,
			MaxBackups: global.App.Config.Log.MaxBackups,
			MaxAge:     global.App.Config.Log.MaxAge,
			Compress:   global.App.Config.Log.Compress,
		}
	} else {
		// 默认 Writer
		writer = os.Stdout
	}
	return log.New(writer, "\r\n", log.LstdFlags)
}

func getGormLogger() logger.Interface {
	var logMode logger.LogLevel

	switch global.App.Config.Mysql.LogMode {
	case "silent":
		logMode = logger.Silent
	case "error":
		logMode = logger.Error
	case "warn":
		logMode = logger.Warn
	case "info":
		logMode = logger.Info
	default:
		logMode = logger.Info
	}

	return logger.New(getGormLogWriter(), logger.Config{
		SlowThreshold:             200 * time.Millisecond,                       // 慢 SQL 阈值
		LogLevel:                  logMode,                                      // 日志级别
		IgnoreRecordNotFoundError: false,                                        // 忽略ErrRecordNotFound（记录未找到）错误
		Colorful:                  !global.App.Config.Mysql.EnableFileLogWriter, // 禁用彩色打印
	})

}

// 创建数据库的客户端
func NewDBClient(ctx context.Context) *gorm.DB {
	db := global.App.DB
	return db.WithContext(ctx)
}
