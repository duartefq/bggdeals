package cmd

import (
	"strings"
	"time"
	"github.com/turnage/graw/reddit"
)


type redditDealsBot struct {
	bot reddit.Bot
	subreddit string
}

type handler interface {
    PostThread(item *Item) error
}

func (r *redditDealsBot) PostThread(item *Item) error {
	<-time.After(15 * time.Second)
	title := strings.TrimPrefix(item.Title, "Thread: Hot Deals:: ") + " (BGG: Hot Deals)"
	return r.bot.PostLink(r.subreddit, title, item.Link)
}

func LoadRedditBot(agentFile string, subreddit string) (*redditDealsBot, error) {
	bot, err := reddit.NewBotFromAgentFile(agentFile, 0)

	if err != nil {
		return nil, err
	}

	return &redditDealsBot{bot: bot, subreddit: subreddit}, nil
}
