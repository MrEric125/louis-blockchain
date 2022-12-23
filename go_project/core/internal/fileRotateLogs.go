package internal

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
	"louis/global"
	"os"
	"path"
	"time"
)

var FileRotatelogs = new(fileRotateLogs)

type fileRotateLogs struct{}

// GetWriteSyncer 获取 zapcore.WriteSyncer
func (r *fileRotateLogs) GetWriteSyncer(level string) (zapcore.WriteSyncer, error) {
	fileWriter, err := rotatelogs.New(
		path.Join(global.LOUIS_CONFIG.Zap.Director, "%Y-%m-%d", level+".log"),
		rotatelogs.WithClock(rotatelogs.Local),
		rotatelogs.WithMaxAge(time.Duration(global.LOUIS_CONFIG.Zap.MaxAge)*24*time.Hour), // 日志留存时间
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if global.LOUIS_CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}
