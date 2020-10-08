package main

import (
	"log"
	"flag"
)

type handler interface {
    PostThread(item *Item) error
}

func post(h handler, item *Item) error {
	return h.PostThread(item)
}

func main() {
	log.SetPrefix("bgg-deals: ")
	log.SetFlags(0)

	filter := flag.String("filter", "", "filter thread titles")
	agent := flag.String("agent", "", "reddit agent file path")
	guid_file := flag.String("guid_file", "", "guid_file path")

	flag.Parse()

	if (flag.NArg() != 1) {
		log.Fatal("Please provide a subreddit")
	}

	if *agent == "" {
		log.Fatal("Please provide an agent file")
	}

	if *guid_file == "" {
		log.Fatal("Please provide a guid_file to store last GUID")
	}

	sub := flag.Arg(0)

	last_guid, err := loadLastGUID(*guid_file)

	bot, err := LoadRedditBot(*agent, sub)

	if err != nil {
		log.Fatal("Failed to create bot handle: ", err)
	}

	bggFeed := BGGFeed{
		url: "https://www.boardgamegeek.com/rss/uforum/10",
		filter: *filter,
		last_guid: last_guid,
		Handler: bot,
	}

	bggFeed.Crawl()
}
