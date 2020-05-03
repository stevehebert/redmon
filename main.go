package main

import (
	"fmt"
	"os"
	"time"

	"github.com/turnage/graw/reddit"
)

type announcer struct{}

func (a *announcer) Post(post *reddit.Post) error {
	fmt.Printf("%s posted \"%s\"\n", post.Author, post.Title)
	return nil
}

func main() {
	bot, err := reddit.NewBotFromAgentFile("cred.file", time.Second*15)

	if err != nil {
		fmt.Println("failed to create bot handle:", err)
		os.Exit(1)
	}

	harvest, err := bot.Listing("/r/mechmarket", "")
	if err != nil {
		fmt.Println("failed to fetch /r/mechmarket: ", err)
		os.Exit(1)

	}

	for _, post := range harvest.Posts[:5] {
		fmt.Printf("[%s] posted [%s]\n", post.Author, post.Title)
	}

}
