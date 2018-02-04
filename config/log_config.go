package config

import (
	l4g "github.com/XGoServer/threeLibs/alecthomas/log4go"
	"os"
	"path/filepath"
)

const LOG_ROTATE_SIZE = 10000
const LOG_FILENAME    = "my.log"

func ConfigureLog(s *LogConfigStruct) {
	l4g.Close()
	if s.EnableConsole {
		level := l4g.DEBUG
		if s.ConsoleLevel == "INFO" {
			level = l4g.INFO
		} else if s.ConsoleLevel == "WARN" {
			level = l4g.WARNING
		} else if s.ConsoleLevel == "ERROR" {
			level = l4g.ERROR
		}
		lw := l4g.NewConsoleLogWriter()
		lw.SetFormat("[%D %T] [%L] %M")
		l4g.AddFilter("stdout", level, lw)
	}

	if s.EnableFile {
		var fileFormat = s.FileFormat
		if fileFormat == "" {
			fileFormat = "[%D %T] [%L] %M"
		}
		level := l4g.DEBUG
		if s.FileLevel == "INFO" {
			level = l4g.INFO
		} else if s.FileLevel == "WARN" {
			level = l4g.WARNING
		} else if s.FileLevel == "ERROR" {
			level = l4g.ERROR
		}
		flw := l4g.NewFileLogWriter(getLogFileLocation(s.FileLocation), false)
		flw.SetFormat(fileFormat)
		flw.SetRotate(true)
		flw.SetRotateLines(LOG_ROTATE_SIZE)
		l4g.AddFilter("file", level, flw)
	}
}
func getLogFileLocation(fileLocation string) string {
	if fileLocation == "" {
		logDir, _ := findDir("logs")
		return logDir + LOG_FILENAME
	} else {
		return fileLocation + LOG_FILENAME
	}
}

func findDir(dir string) (string, bool) {
	fileName := "."
	found := false
	if _, err := os.Stat("./" + dir + "/"); err == nil {
		fileName, _ = filepath.Abs("./" + dir + "/")
		found = true
	} else if _, err := os.Stat("../" + dir + "/"); err == nil {
		fileName, _ = filepath.Abs("../" + dir + "/")
		found = true
	} else if _, err := os.Stat("../../" + dir + "/"); err == nil {
		fileName, _ = filepath.Abs("../../" + dir + "/")
		found = true
	}

	return fileName + "/", found
}

