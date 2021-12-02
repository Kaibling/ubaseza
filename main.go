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
)

func main() {
	config.LoadConfig()
	transltrService := transltr.NewTransltr()
	libTrans := translation.NewLibreTranslateService()
	transltrService.AddPlugin(libTrans)
	ic := irc.NewIRCClient("irc.chat.twitch.tv:6667")
	ic.AddTranslator(transltrService)
	channels := []string{"ashleyroboto", "koreshzy"}
	for _, channel := range channels {
		ic.ConfigureClient(channel)
		go ic.ReadChannel(channel)
	}
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
			fmt.Println("Warikomi")
			break
		}

	}
	os.Exit(1)
}
