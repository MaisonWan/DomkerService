// DomkerService project main.go
package main

import (
	"DomkerService/notification/config"
	"DomkerService/notification/email"
	"DomkerService/notification/watcher"
	"fmt"
	"strconv"
	"time"
)

var timeline time.Time

func tickCheckNewNotes(seconds time.Duration, d watcher.Douban, e email.Mail, c *config.Config) {
	ticker := time.Tick(time.Second * seconds)
	checkNewNotes(d, e, c)
	for _ = range ticker {
		checkNewNotes(d, e, c)
		//fmt.Println("Now is ", now)
	}
}

func checkNewNotes(d watcher.Douban, e email.Mail, c *config.Config) {
	notes := d.CheckNewNote(timeline)
	if timeline.IsZero() {
		timeline = d.GetTimeLine(notes)
		notes = d.FilterNotes(notes, timeline)
	}
	body := ""
	for _, note := range notes {
		note.Content = d.GetNoteDetail(note.Url)
		body += d.ConvertEmailBody(note)
	}
	if len(notes) > 0 {
		timeline = notes[0].CreateTime
		title := "[豆瓣]发现" + c.GetDoubanUser() + "有" + strconv.Itoa(len(notes)) + "个新日记"
		content := d.ConvertEmailContent(body)
		err := e.Send(c.GetEmailReceiver(), title, content, "html")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(time.Now(), title)
	} else {
		fmt.Println("checkNewNotes no result.", time.Now())
	}
}

func main() {
	c := config.GetConfig()
	e := email.Mail{}
	e.Init(c)

	douban := watcher.Douban{}
	douban.Init(c.GetDoubanUser())
	interval := time.Duration(c.GetEmailInterval()) * time.Minute
	tickCheckNewNotes(interval, douban, e, c)

	fmt.Println("send email end")
}
