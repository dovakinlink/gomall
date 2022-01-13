package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Field = zap.Field

var (
	Logger  *zap.Logger
	String  = zap.String
	Any     = zap.Any
	Int     = zap.Int
	Float32 = zap.Float32
)

func InitLogger(logpath string, loglevel string) {
	// 日志分割
	hook := lumberjack.Logger{
		Filename:   logpath, // 日志文件路径，默认 os.TempDir()
		MaxSize:    100,     // 每个日志文件保存100M，默认 100M
		MaxBackups: 30,      // 保留30个备份，默认不限
		MaxAge:     7,       // 保留7天，默认不限
		Compress:   true,    // 是否压缩，默认不压缩
	}
	write := zapcore.AddSync(&hook)

	// 设置日志级别
	// debug 可以打印出 info debug warn
	// info  级别可以打印 warn info
	// warn  只能打印 warn
	// debug->info->warn->error
	var level zapcore.Level
	switch loglevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	case "warn":
		level = zap.WarnLevel
	default:
		level = zap.InfoLevel
	}

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "linenum",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseColorLevelEncoder, // 小写
		EncodeTime:     zapcore.ISO8601TimeEncoder,         // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder, // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)

	var writes = []zapcore.WriteSyncer{write}
	// 如果是开发环境，同时在控制台上也输出
	if level == zap.DebugLevel {
		writes = append(writes, zapcore.AddSync(os.Stdout))
	}
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(writes...),
		level,
	)
	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段，如：添加一个服务器名称
	filed := zap.Fields(zap.String("application", "gomall"))
	// 构造日志
	Logger = zap.New(core, caller, development, filed)
	Logger.Info("logger init success")
}
