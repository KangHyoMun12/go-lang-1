package logger

import (
	"fmt"

	seelog "github.com/cihub/seelog"
)

/*
Logger log
*/
var Logger seelog.LoggerInterface

func loadAppConfig() {
	appConfig := `
	<?xml version="1.0" encoding="UTF-8"?>
	<seelog minlevel="trace">
    <outputs formatid="common">
        <rollingfile type="size" filename="/Golang/Programs/StartWebWithGo/roll.log" maxsize="100000" maxrolls="5"/>
		<filter levels="critical">
            <file path="/Golang/Programs/StartWebWithGo/critical.log" formatid="critical"/>
            <smtp formatid="criticalemail" senderaddress="astaxie@gmail.com" sendername="ShortUrl API" hostname="smtp.gmail.com" hostport="587" username="mailusername" password="mailpassword">
                <recipient address="xiemengjun@gmail.com"/>
            </smtp>
        </filter>
    </outputs>
    <formats>
        <format id="common" format="%Date/%Time [%LEV] %Msg%n" />
	    <format id="critical" format="%File %FullPath %Func %Msg%n" />
	    <format id="criticalemail" format="Critical error on our server!\n    %Time %Date %RelFile %Func %Msg \nSent by Seelog"/>
    </formats>
</seelog>
	`
	logger, err := seelog.LoggerFromConfigAsBytes([]byte(appConfig))
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
