package tools

import (
	log "github.com/cihub/seelog"
	"sync"
)

// 单例创建日志处理器

var (
	logger SingleLogger
	once   sync.Once
)

type SingleLogger struct {
	Logger log.LoggerInterface
}

func GetLoggerInstance() SingleLogger {
	once.Do(func() {
		logger = SingleLogger{
			Logger: SetupLogger(),
		}
	})
	return logger
}

func SetupLogger() log.LoggerInterface {
	defer log.Flush()
	// 加载日志配置文件
	logger, err := log.LoggerFromConfigAsFile("./conf/seelog.xml")
	if err != nil {
		log.Errorf("parse config.xml error: %v", err)
		return nil
	}
	// 替换日志
	log.ReplaceLogger(logger)
	return logger
}
