package mysql

import (
	"wiki_global/src/config"

	"xorm.io/core"

	"github.com/arthurkiller/rollingwriter"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	Engine *xorm.Engine
	writer rollingwriter.RollingWriter
)

// Init 初始化
func Init(driverName, mysqlDSN string, maxIdleConns, maxOpenConns int) (err error) {
	Engine, err = xorm.NewEngine(driverName, mysqlDSN)
	if err != nil {
		return err
	}

	if err := Engine.Ping(); err != nil {
		return err
	}

	Engine.SetMaxIdleConns(maxIdleConns)
	Engine.SetMaxOpenConns(maxOpenConns)
	//Engine.DatabaseTZ = time.Local
	//Engine.TZLocation = time.Local

	config := rollingwriter.Config{
		LogPath:                config.Conf().Log.Path,      //日志路径
		TimeTagFormat:          "060102150405",              //时间格式串
		FileName:               "mysql",                     //日志文件名
		MaxRemain:              5,                           //配置日志最大存留数
		RollingPolicy:          rollingwriter.VolumeRolling, //配置滚动策略 norolling timerolling volumerolling
		RollingTimePattern:     "* * * * * *",               //配置时间滚动策略
		RollingVolumeSize:      "1M",                        //配置截断文件下限大小
		WriterMode:             "none",
		BufferWriterThershould: 8 * 1024 * 1024,
		Compress:               true,
	}

	writer, err := rollingwriter.NewWriterFromConfig(&config)
	if err != nil {
		return err
	}

	var logger *xorm.SimpleLogger = xorm.NewSimpleLogger(writer)

	Engine.SetLogger(logger)
	Engine.ShowSQL(true)
	Engine.Logger().SetLevel(core.LOG_DEBUG)

	return
}
