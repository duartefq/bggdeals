package main

import (
	"github.com/mmcdole/gofeed"
	"io/ioutil"
	"strings"
	"log"
)

type LastGUID struct {
	Title string
	GUID  string
}

func (p *LastGUID) save() error {
	return ioutil.WriteFile(p.Title, []byte(p.GUID), 0600)
}

func loadLastGUID(title string) (*LastGUID, error) {
	body, err := ioutil.ReadFile(title)
	if err != nil {
		return &LastGUID{Title: title, GUID: ""}, err
	}
	return &LastGUID{Title: title, GUID: string(body)}, nil
}

type Item struct {
	Title string
	Link string
}

type BGGFeed struct {
	url string
	filter string
	last_guid *LastGUID
	Handler handler
}

func (b *BGGFeed) GetItems() []*gofeed.Item {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(b.url)
	return feed.Items
}

func (b *BGGFeed) Crawl() {
	items := b.GetItems()

	count := 0

	for _, item := range items {
		if (item.GUID == b.last_guid.GUID) {
			break
		}
		if (!strings.Contains(item.Title, b.filter)) {
			continue
		}

		err := post(b.Handler, &Item{ Title: item.Title, Link: item.Link })

		if (err != nil) {
			log.Println("Error posting this thread", err)
		} else {
			log.Println("Post created: ", item.Title, item.Link)
			count += 1
		}
	}

	log.Println(count, "Processed")

	b.last_guid.GUID = items[0].GUID
	b.last_guid.save()
}
