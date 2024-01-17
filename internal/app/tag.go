package app

import (
	"time"
)

type tag struct {
	tagName   string
	body      string
	createdAt time.Time
}

func (t tag) TagName() string {
	return t.tagName
}

func (t tag) Body() string {
	return t.body
}

func (t tag) CreatedAt() string {
	return t.createdAt.Format("2006-01-02")
}
