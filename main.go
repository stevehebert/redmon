package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/sfreiberg/gotwilio"
	"github.com/turnage/graw"
	"github.com/turnage/graw/reddit"
)

type reminderBot struct {
	bot reddit.Bot
}

type TwilioConfiguration struct {
	AccountSid string
	AuthToken  string
	FromPhone  string
	ToPhone    string
}

func getTwilioClient() (*gotwilio.Twilio, TwilioConfiguration) {
	content, err := ioutil.ReadFile("twilio.toml")
	if err != nil {
		fmt.Println("unable to read twilio.toml: ", err)
		os.Exit(1)
	}

	var twilioConfiguration TwilioConfiguration

	if _, err := toml.Decode(string(content), &twilioConfiguration); err != nil {
		fmt.Println("unable to translate twilio.toml file: ", err)
		os.Exit(1)
	}

	return gotwilio.NewTwilioClient(twilioConfiguration.AccountSid, twilioConfiguration.AuthToken), twilioConfiguration

}

func (r *reminderBot) Post(p *reddit.Post) error {
	fmt.Printf("watching post by [%s] -- [%s] @ \n[%s]\n\n", p.Author, p.Title, p.URL)
	if (strings.Contains(p.Title, "Tokyo") && strings.Contains(p.Title, "60")) || strings.Contains(p.Title, "Tokyo60") || true {
		fmt.Printf("notifying of match at: %s", p.URL)

		twilio, conf := getTwilioClient()

		twilio.SendSMS(conf.FromPhone, conf.ToPhone, fmt.Sprintf("Monitored message from [%s] - [%s]: %s", p.Author, p.Title, p.URL), "", "")

		<-time.After(10 * time.Second)
		return r.bot.SendMessage(
			p.Author,
			fmt.Sprintf("Notification: %s", p.Title),
			fmt.Sprintf("Tokyo60: %s", p.URL),
		)
	}
	return nil
}

func main() {
	if bot, err := reddit.NewBotFromAgentFile("redmond.agent", 0); err != nil {
		fmt.Println("Failed to create bot handle: ", err)
	} else {
		cfg := graw.Config{Subreddits: []string{"mechmarket"}}
		handler := &reminderBot{bot: bot}
		if _, wait, err := graw.Run(handler, bot, cfg); err != nil {
			fmt.Println("Failed to start graw run: ", err)
			os.Exit(1)
		} else {
			fmt.Println("graw run failed: ", wait())
			os.Exit(1)
		}
	}
}
