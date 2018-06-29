package util

import (
	"fmt"
	"os"
	"strconv"
	"time"

	log "github.com/cihub/seelog"
)

type LogStruct struct {
	logger *log.LoggerInterface
}

var hostName string
var pid int

var levelToSyslogSeverity = map[log.LogLevel]int{
	// Mapping to RFC 5424 where possible
	log.TraceLvl:    7,
	log.DebugLvl:    7,
	log.InfoLvl:     6,
	log.WarnLvl:     4,
	log.ErrorLvl:    3,
	log.CriticalLvl: 2,
	log.Off:         7,
}

func CreateSyslogHeaderFormatter(params string) log.FormatterFunc {
	facility := 20
	i, err := strconv.Atoi(params)
	if err == nil && i >= 0 && i <= 23 {
		facility = i
	}

	return func(message string, level log.LogLevel, context log.LogContextInterface) interface{} {
		return fmt.Sprintf("<%d>1 %s %s %s %d - -", facility*8+levelToSyslogSeverity[level],
			time.Now().Format("2006-01-02T15:04:05.000000Z07:00"),
			hostName, pid)
	}
}

// Config : log config
func (l *LogStruct) Config(host string) error {
	logConfig := `
	<seelog>
		<outputs formatid="main">
			<console/>
			<conn formatid="syslog" net="udp4" addr="` + host + `" />
			<rollingfile type="date" filename="thanos.log" datepattern="2006.01.02" maxrolls="30"/>
		</outputs>
		<formats>
			<format id="main" format="%Date(2006-01-02T15:04:05.000000000) [%Level] %Msg%n"/>
			<format id="syslog" format="%CustomSyslogHeader(20) %Msg%n"/>
		</formats>
	</seelog>
		`
	if host == "" {
		logConfig = `
		<seelog>
			<outputs formatid="main">
				<console/>
				<rollingfile type="date" filename="thanos.log" datepattern="2006.01.02" maxrolls="30"/>
			</outputs>
			<formats>
				<format id="main" format="%Date(2006-01-02T15:04:05.000000000) [%Level] %Msg%n"/>
				<format id="syslog" format="%CustomSyslogHeader(20) %Msg%n"/>
			</formats>
		</seelog>
		`
	}

	hostName, _ = os.Hostname()
	pid = os.Getpid()

	if err := log.RegisterCustomFormatter("CustomSyslogHeader", CreateSyslogHeaderFormatter); err != nil {
		fmt.Println("log register fail:", err)
		return err
	}
	if logger, err := log.LoggerFromConfigAsString(logConfig); err != nil {
		fmt.Println("log config fail:", err)
		return err
	} else {
		l.logger = &logger
		log.UseLogger(logger)
	}

	return nil
}

func (l *LogStruct) Flush() {
	log.Flush()
}

func (l *LogStruct) Debug(a ...interface{}) {
	log.Debug(a...)
}

func (l *LogStruct) Info(a ...interface{}) {
	log.Info(a...)
}

func (l *LogStruct) Error(a ...interface{}) {
	log.Error(a...)
}

func (l *LogStruct) Warn(a ...interface{}) {
	log.Warn(a...)
}

func (l *LogStruct) Critical(a ...interface{}) {
	log.Critical(a...)
}
