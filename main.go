package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/turnage/graw"
	"github.com/turnage/graw/reddit"
)

type reminderBot struct {
	bot reddit.Bot
}

func (r *reminderBot) Post(p *reddit.Post) error {
	if strings.Contains(p.Title, "Tokyo60") {
		fmt.Printf("notifying of match at: %s", p.URL)
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
		} else {
			fmt.Println("graw run failed: ", wait())
		}
	}
}
