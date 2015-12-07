package xylog

//定义了业务调试日志相关接口

import (
	"fmt"
	"os"

	logs "github.com/beego/logs"
)

//日志级别枚举值
type LogLevel int

const (
	EmergencyLevel     LogLevel = logs.LevelEmergency     //0
	AlertLevel         LogLevel = logs.LevelAlert         //1
	CriticalLevel      LogLevel = logs.LevelCritical      //2
	ErrorLevel         LogLevel = logs.LevelError         //3
	WarningLevel       LogLevel = logs.LevelWarning       //4
	NoticeLevel        LogLevel = logs.LevelNotice        //5
	InformationalLevel LogLevel = logs.LevelInformational //6
	DebugLevel         LogLevel = logs.LevelDebug         //7
)

func (l LogLevel) String() (str string) {
	switch l {
	case EmergencyLevel:
		str = fmt.Sprintf("Emergency (%d)", int(l))
	case AlertLevel:
		str = fmt.Sprintf("Alert (%d)", int(l))
	case CriticalLevel:
		str = fmt.Sprintf("Critical (%d)", int(l))
	case ErrorLevel:
		str = fmt.Sprintf("Error (%d)", int(l))
	case WarningLevel:
		str = fmt.Sprintf("Warning (%d)", int(l))
	case NoticeLevel:
		str = fmt.Sprintf("Notice (%d)", int(l))
	case InformationalLevel:
		str = fmt.Sprintf("Informational (%d)", int(l))
	case DebugLevel:
		str = fmt.Sprintf("DebugLevel (%d)", int(l))
	default:
		str = fmt.Sprintf("Undefined (%d)", int(l))
	}
	return
}

//日志管理器
type XYLogger struct {
	config *LoggerConfig   //日志配置信息
	beelog *logs.BeeLogger //beego日志指针
}

var (
	def *XYLogger = NewLogger(defConfig, 1000)
)

func NewLogger(lc *LoggerConfig, chanlen int64) (l *XYLogger) {
	l = &XYLogger{
		config: lc,
		beelog: logs.NewLogger(chanlen),
	}

	return
}

func (l *XYLogger) Logger() *logs.BeeLogger {
	return l.beelog
}

func (l *XYLogger) Config() *LoggerConfig {
	return l.config
}

func (l *XYLogger) ApplyConfig(lc *LoggerConfig) {
	if lc == nil {
		lc = defConfig
	}
	l.beelog.SetLevel(int(l.config.Level))

	if l.config.Verbose {
		l.beelog.EnableFuncCallDepth(true)
		l.beelog.SetLogFuncCallDepth(4)
	} else {
		l.beelog.EnableFuncCallDepth(false)
		l.beelog.SetLogFuncCallDepth(0)
	}

	var strconfig string
	if l.config.Stdout {
		//strconfig = fmt.Sprintf(`{"level":%v}`, l.config.Level)
		l.beelog.SetLogger("console", strconfig)
	} else {
		if l.config.Filename == "" {
			if l.config.NodeId >= 0 {
				l.config.LogId = l.config.NodeId
			} else {
				//l.config.Filename = fmt.Sprintf("%s.%d.log", l.config.AppName, os.Getpid())
				l.config.LogId = os.Getpid()
			}
			l.config.Filename = fmt.Sprintf("%s.%d.log", l.config.AppName, l.config.LogId)
		}

		strconfig = fmt.Sprintf(`{"filename":"%s/%s","maxlines":%v,"maxsize":%v,"daily":%v,"maxdays":%v,"rotate":%v}`,
			l.config.Path,
			l.config.Filename,
			l.config.Maxlines,
			l.config.Maxsize,
			l.config.Daily,
			l.config.Maxdays,
			l.config.Rotate)
		l.beelog.SetLogger("file", strconfig)
	}
}

func (l *XYLogger) SetLogLevel(level LogLevel) {
	l.beelog.SetLevel((int)(level))
}

func (l *XYLogger) Log(log_level LogLevel, format string, v ...interface{}) {
	switch log_level {
	case logs.LevelCritical:
		l.beelog.Critical(format, v...)
	case logs.LevelDebug:
		l.beelog.Debug(format, v...)
	case logs.LevelError:
		l.beelog.Error(format, v...)
	case logs.LevelInfo:
		l.beelog.Info(format, v...)
	//case logs.LevelTrace:
	//	l.beelog.Trace(format, v...)
	case logs.LevelWarn:
		l.beelog.Warn(format, v...)
	}
}

func ApplyConfig(lc *LoggerConfig) {
	def.ApplyConfig(lc)
}

func EmergencyNoId(format string, v ...interface{}) {
	if def.config.Level < EmergencyLevel {
		return
	}

	def.Log(EmergencyLevel, format, v...)
}

func AlertNoId(format string, v ...interface{}) {
	if def.config.Level < AlertLevel {
		return
	}

	def.Log(AlertLevel, format, v...)
}

func CriticalNoId(format string, v ...interface{}) {
	if def.config.Level < CriticalLevel {
		return
	}

	def.Log(CriticalLevel, format, v...)
}

func DebugNoId(format string, v ...interface{}) {
	if def.config.Level < DebugLevel {
		return
	}

	def.Log(DebugLevel, format, v...)
}

func WarningNoId(format string, v ...interface{}) {
	if def.config.Level < WarningLevel {
		return
	}
	def.Log(WarningLevel, format, v...)
}

func ErrorNoId(format string, v ...interface{}) {

	if def.config.Level < ErrorLevel {
		return
	}

	def.Log(ErrorLevel, format, v...)

}

func NoticeNoId(format string, v ...interface{}) {

	if def.config.Level < NoticeLevel {
		return
	}

	def.Log(NoticeLevel, format, v...)
}

func InformationalNoId(format string, v ...interface{}) {

	if def.config.Level < InformationalLevel {
		return
	}

	def.Log(InformationalLevel, format, v...)
}

func Emergency(id interface{}, format string, v ...interface{}) {
	if IsNeedLog(id, EmergencyLevel) {
		format = fmt.Sprintf("[%v] %s", id, format)
		def.Log(EmergencyLevel, format, v...)
	}
}

func Alert(id interface{}, format string, v ...interface{}) {
	if IsNeedLog(id, AlertLevel) {
		format = fmt.Sprintf("[%v] %s", id, format)
		def.Log(AlertLevel, format, v...)
	}
}

func Debug(id interface{}, format string, v ...interface{}) {
	if IsNeedLog(id, DebugLevel) {
		format = fmt.Sprintf("[%v] %s", id, format)
		def.Log(DebugLevel, format, v...)
	}
}

func Warning(id interface{}, format string, v ...interface{}) {
	if IsNeedLog(id, WarningLevel) {
		format = fmt.Sprintf("[%v] %s", id, format)
		def.Log(WarningLevel, format, v...)
	}
}

func Error(id interface{}, format string, v ...interface{}) {
	if IsNeedLog(id, ErrorLevel) {
		format = fmt.Sprintf("[%v] %s", id, format)
		def.Log(ErrorLevel, format, v...)
	}
}

func Critical(id interface{}, format string, v ...interface{}) {
	if IsNeedLog(id, CriticalLevel) {
		format = fmt.Sprintf("[%v] %s", id, format)
		def.Log(CriticalLevel, format, v...)
	}
}

func Notice(id interface{}, format string, v ...interface{}) {
	if IsNeedLog(id, NoticeLevel) {
		format = fmt.Sprintf("[%v] %s", id, format)
		def.Log(NoticeLevel, format, v...)
	}
}

func Informational(id interface{}, format string, v ...interface{}) {
	if IsNeedLog(id, InformationalLevel) {
		format = fmt.Sprintf("[%v] %s", id, format)
		def.Log(InformationalLevel, format, v...)
	}
}

//设置全局的日志级别
func SetLogLevel(level LogLevel) {
	def.SetLogLevel(level)
}

func IsNeedLog(id interface{}, level LogLevel) bool {

	if !(DefIdManager.IsIdExist(id)) {
		if def.config.Level < level {
			return false
		}
	}

	return true
}

func Close() {
	def.beelog.Close()
}
