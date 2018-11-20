package logger

import (
	"fmt"
	"io/ioutil"

	seelog "github.com/cihub/seelog"
)

/*
Logger log
*/
var Logger seelog.LoggerInterface

func loadAppConfig() {
	// appConfig, err := os.Open("log.xml")
	// if err != nil {
	// 	panic(err.Error())
	// }
	file, err := ioutil.ReadFile("log.xml")
	if err != nil {
		panic(err.Error())
	}
	logger, err := seelog.LoggerFromConfigAsBytes(file)
	// logger, err := seelog.LoggerFromConfigAsFile("./logger/log.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	UseLogger(logger)
}

func init() {
	DisableLog()
	loadAppConfig()
}

// DisableLog disables all library log output
func DisableLog() {
	Logger = seelog.Disabled
}

// UseLogger uses a specified seelog.LoggerInterface to output library log.
// Use this func if you are using Seelog logging system in your app.
func UseLogger(newLogger seelog.LoggerInterface) {
	Logger = newLogger
}
