package logs

import (
	"path/filepath"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

func getLogWriter(logPath string, maxSize, maxAge, maxBackups int, compress bool) zapcore.WriteSyncer {
	fileName := filepath.Join(logPath, "campus-forum-system.log") // Join将任意数量的路径元素联接到单个路径中， 使用特定于 OS 的分隔符将它们分开。
	lumberJackerLogger := &lumberjack.Logger{
		Filename:   fileName,   // 日志文件的位置
		MaxSize:    maxSize,    // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxAge:     maxAge,     // 保留旧文件的最大天数
		MaxBackups: maxBackups, // 保留旧文件的最大个数
		Compress:   compress,   // 是否压缩/归档旧文件
	}
	return zapcore.AddSync(lumberJackerLogger)
}

// 将JSON Encoder更改为普通的Log Encoder
func getEncoder() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder   // 修改时间编码器
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 在日志文件中使用大写字母记录日志级别
	return zapcore.NewConsoleEncoder(encodeConfig)
}

func InitLogger(logPath string, maxSize, maxAge, maxBackups int, compress bool) {
	writeSyncer := getLogWriter(logPath, maxSize, maxAge, maxBackups, compress)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	Logger = zap.New(core, zap.AddCaller()).Sugar()
}

func GetLogger() *zap.SugaredLogger {
	return Logger
}
