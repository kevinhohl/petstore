package model

import (
	"strconv"
	"strings"
)

type Pet struct {
	ID           int `json:"id"`
	Category     Category `json:"category"`
	Name         string   `json:"name"`
	PhotoUrlsRaw string `json:"-"`
	PhotoUrls    []string `json:"photoUrls"`
	TagsIDsRaw   string `json:"-"`
	TagsRaw      string `json:"-"`
	Tags         []Tag `json:"tags"`
	Status       string `json:"status"`
}

func (p Pet) UnRaw() Pet {
	p.PhotoUrls = strings.Split(p.PhotoUrlsRaw, ", ")
	var t Tag
	var tg []Tag
	tags := strings.Split(p.TagsRaw, ", ")
	tagIDs := strings.Split(p.TagsIDsRaw, ", ")
	for i, item := range tagIDs {
		t.ID, _ = strconv.Atoi(item)
		t.Name = tags[i]
		tg = append(tg, t)
	}
	p.Tags = tg
	return p
}

type Status struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var StatusToName = map[string]int {
	"available": 10,
	"pending": 20,
	"sold": 30,
}
