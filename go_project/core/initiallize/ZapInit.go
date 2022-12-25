package initiallize

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"louis/global"
	"louis/utils"
	"os"
	"path"
	"time"
)

var ZapObj = new(_zap)

var logger *zap.SugaredLogger

type _zap struct{}

func ZapInit() {

	ok, _ := utils.PathExists(global.LOUIS_CONFIG.Zap.Director)

	if !ok {
		fmt.Printf("create %v directory\n", global.LOUIS_CONFIG.Zap.Director)
		_ = os.Mkdir(global.LOUIS_CONFIG.Zap.Director, os.ModePerm)
	}
	cores := ZapObj.getZapCores()
	log := zap.New(zapcore.NewTee(cores...), zap.AddCaller())

	defer log.Sync()

	if global.LOUIS_CONFIG.Zap.ShowLine {
		log = log.WithOptions(zap.AddCaller())
	}
	logger = log.Sugar()

}
func (z *_zap) getEncoderCore(l zapcore.Level, level zap.LevelEnablerFunc) zapcore.Core {
	writer, err := z.getWriteSyncer(l.String()) // 使用file-rotatelogs进行日志分割
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return nil
	}
	return zapcore.NewCore(z.getEncoder(), writer, level)
}

// GetWriteSyncer 获取 zapcore.WriteSyncer
func (r *_zap) getWriteSyncer(level string) (zapcore.WriteSyncer, error) {
	var baseDir string
	if global.LOUIS_CONFIG.Zap.Director != "" {
		baseDir = global.LOUIS_CONFIG.Zap.Director
	} else {
		baseDir = ".log"
	}
	fileWriter, err := rotatelogs.New(
		path.Join(baseDir, "%Y-%m-%d", level+".log"),
		rotatelogs.WithClock(rotatelogs.Local),
		rotatelogs.WithMaxAge(time.Duration(global.LOUIS_CONFIG.Zap.MaxAge)*24*time.Hour), // 日志留存时间
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if global.LOUIS_CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}

func (z *_zap) getZapCores() []zapcore.Core {
	cores := make([]zapcore.Core, 0, 7)
	for level := global.LOUIS_CONFIG.Zap.TransportLevel(); level <= zapcore.FatalLevel; level++ {
		cores = append(cores, z.getEncoderCore(level, z.getLevelPriority(level)))
	}
	return cores
}

// GetLevelPriority 根据 zapcore.Level 获取 zap.LevelEnablerFunc
func (z *_zap) getLevelPriority(level zapcore.Level) zap.LevelEnablerFunc {
	switch level {
	case zapcore.DebugLevel:
		return func(level zapcore.Level) bool { // 调试级别
			return level == zap.DebugLevel
		}
	case zapcore.InfoLevel:
		return func(level zapcore.Level) bool { // 日志级别
			return level == zap.InfoLevel
		}
	case zapcore.WarnLevel:
		return func(level zapcore.Level) bool { // 警告级别
			return level == zap.WarnLevel
		}
	case zapcore.ErrorLevel:
		return func(level zapcore.Level) bool { // 错误级别
			return level == zap.ErrorLevel
		}
	case zapcore.DPanicLevel:
		return func(level zapcore.Level) bool { // dpanic级别
			return level == zap.DPanicLevel
		}
	case zapcore.PanicLevel:
		return func(level zapcore.Level) bool { // panic级别
			return level == zap.PanicLevel
		}
	case zapcore.FatalLevel:
		return func(level zapcore.Level) bool { // 终止级别
			return level == zap.FatalLevel
		}
	default:
		return func(level zapcore.Level) bool { // 调试级别
			return level == zap.DebugLevel
		}
	}
}
func (z *_zap) getEncoder() zapcore.Encoder {
	var encoderConfig zapcore.EncoderConfig

	if global.LOUIS_CONFIG.Zap.Format == "json" {
		encoderConfig = z.getEncoderConfig()

	} else {
		encoderConfig = zap.NewProductionEncoderConfig()
		encoderConfig.EncodeTime = z.CustomTimeEncoder
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	}
	//consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	//allCore = append(allCore, zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel))

	return zapcore.NewConsoleEncoder(encoderConfig)

}

func (z *_zap) CustomTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(global.LOUIS_CONFIG.Zap.Prefix + t.Format("2006/01/02 - 15:04:05.000"))
}

func (z *_zap) getEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  global.LOUIS_CONFIG.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    global.LOUIS_CONFIG.Zap.ZapEncodeLevel(),
		EncodeTime:     z.CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
}
