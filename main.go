package main

import (
	"log"
	"os"
	"regexp"
	"unicode/utf8"

	"github.com/slack-go/slack"
)

func main() {
	r := regexp.MustCompile(`^(.(\s|　))+`)

	api := slack.New(
		os.Getenv("SLACK_TOKEN"),
	)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {

		case *slack.MessageEvent:
			if ev.User != os.Getenv("BOT_ID") {
				SendMsg(ev.Channel, ev.Text, api, r)
			}
		default:

		}
	}
}

func SendMsg(channel, msgText string, api *slack.Client, r *regexp.Regexp) {
	text := ":_:\n密　を　ふ　せ　ぎ　ま　し　ょ　う\n:_:"

	if utf8.RuneCountInString(msgText) > 1 {
		if !r.MatchString(msgText) {
			text = ":_:\n:_:\nあ　な　た　の　メ　ッ　セ　ー　ジ　は　密　で　す\n:_:\nテ　キ　ス　ト　に　空　白　を　入　れ　ま　し　ょ　う\n:_:"
		}
	}

	_, _, err := api.PostMessage(channel,
		slack.MsgOptionAsUser(true),
		slack.MsgOptionText(text, false),
		slack.MsgOptionAttachments(slack.Attachment{}))

	if err != nil {
		log.Println(err)
	}
}
