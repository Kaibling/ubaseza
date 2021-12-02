package irc

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/kaibling/zucktrans/logging"
	"github.com/kaibling/zucktrans/transltr"
	"github.com/sirupsen/logrus"
	"gopkg.in/irc.v3"
)

type IRCClient struct {
	client     map[string]*irc.Client
	translator *transltr.Transltr
	url        string
}

func NewIRCClient(url string) *IRCClient {
	return &IRCClient{client: make(map[string]*irc.Client), url: url}
}

func (s *IRCClient) ConfigureClient(channelName string) {
	conn, err := net.Dial("tcp", s.url)
	if err != nil {
		log.Fatalln(err)
	}
	config := irc.ClientConfig{
		Nick: "justinfan342",
		Pass: "",
		User: "justinfan342",
		Handler: irc.HandlerFunc(func(c *irc.Client, m *irc.Message) {
			if m.Command == "001" {
				c.Write(fmt.Sprintf("JOIN #%s", channelName))
			} else if m.Command == "PRIVMSG" && c.FromChannel(m) {
				mas, _ := c.ReadMessage()
				rawText := mas.Params[1:]
				text := strings.Join(rawText, " ")
				prefixUser := mas.Prefix.Name
				lang, err := s.translator.GetLanguage(text, "LibreTranslate")
				if err != nil {
					logging.Logger.WithFields(logrus.Fields{
						"user":     prefixUser,
						"message":  text,
						"language": nil,
						"channel":  channelName,
					}).Error("language not found: %s", err.Error())
					return
				}
				transText, err := s.translator.Translate(text, "LibreTranslate")
				if err != nil {
					logging.Logger.WithFields(logrus.Fields{
						"user":     prefixUser,
						"message":  text,
						"language": lang,
						"channel":  channelName,
					}).Error("transaltion failed: %s", err.Error())
					return
				}
				logging.Logger.WithFields(logrus.Fields{
					"user":       prefixUser,
					"message":    text,
					"language":   lang,
					"translated": transText,
					"channel":    channelName,
				}).Info()

			} else {
				mas, _ := c.ReadMessage()
				logging.Logger.Debug(mas)
			}
		}),
	}

	client := irc.NewClient(conn, config)
	s.client[channelName] = client
}

func (s *IRCClient) AddTranslator(translator *transltr.Transltr) {
	s.translator = translator
}
func (s *IRCClient) ReadChannel(channelName string) {
	if _, ok := s.client[channelName]; !ok {
		logging.Logger.Warnf("no client for channel '%s'", channelName)
		return
	}
	err := s.client[channelName].Run()
	if err != nil {
		log.Fatalln(err)
	}
}
