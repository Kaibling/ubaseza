package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/kaibling/zucktrans/Plugins/chat/irc"
	"github.com/kaibling/zucktrans/Plugins/translation"
	"github.com/kaibling/zucktrans/config"
	"github.com/kaibling/zucktrans/transltr"
	"github.com/kaibling/zucktrans/webserver"
)

func main() {
	config.LoadConfig()
	transltrService := transltr.NewTransltr()
	libTrans := translation.NewLibreTranslateService()
	transltrService.AddPlugin(libTrans)

	c := make(chan string)
	ic := irc.NewIRCClient("irc.chat.twitch.tv:6667", c)
	ic.AddTranslator(transltrService)
	ic.Run(config.Config.IRC.Channels)

	ws := webserver.NewWebServer(c)
	ws.Configure()
	ws.Start()

	block()

}

func block() {
	signal_chan := make(chan os.Signal, 1)
	signal.Notify(signal_chan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	for {
		s := <-signal_chan
		if s == syscall.SIGINT {
			fmt.Println("Interupt signal")
			break
		}

	}
	os.Exit(1)
}

//export ZT_IRC_PLUGIN=notUsed
//export ZT_IRC_CHANNELS=dishkaz,mira
//export ZT_IRC_LANG_PLUGIN=notUsed
