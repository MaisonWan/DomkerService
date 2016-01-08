package watcher

import (
	"bufio"
	"bytes"
	"fmt"
	"html"
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type Note struct {
	Subject    string
	Content    string
	CreateTime time.Time
	Url        string
}

const baseUrl = "http://wap.douban.com/"

type Douban struct {
	url          string
	timeline     time.Time
	initTimeline bool
}

func (d *Douban) Init(user string) {
	d.url = baseUrl + "people/" + user + "/notes"
	d.initTimeline = false
}

func (d *Douban) UpdateTimeline(t time.Time) {
	d.timeline = t
}

func (d *Douban) CheckNewNote(timeline time.Time) []Note {
	return d.CheckNewNoteWithUrl(d.url, timeline)
}

func (d *Douban) CheckNewNoteWithUrl(url string, timeline time.Time) []Note {
	content, statusCode := getContent(url)
	if statusCode != 200 {
		return nil
	}
	notes := getNotes(content)
	newNotes := d.FilterNotes(notes, timeline)
	//for i, n := range newNotes {
	//	fmt.Println(i, n)
	//}
	return newNotes
}

func (d *Douban) TickCheckNewNotes(seconds time.Duration) {
	ticker := time.Tick(time.Second * seconds)
	for now := range ticker {
		fmt.Println("Now is ", now)
	}
}

// 读取url的文本内容
func getContent(url string) (content string, statusCode int) {
	resp, err := http.Get(url)
	if err != nil {
		statusCode = -100
		return
	}
	// 结束之前关闭掉流
	defer resp.Body.Close()
	data, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		statusCode = -200
		return
	}
	statusCode = resp.StatusCode
	content = string(data)
	return
}

var noteIndexItem = regexp.MustCompile(`<a href="/note/(.+)/">.*</a>
  <span>\s.*</span>`)

var noteHrefItem = regexp.MustCompile(`href=".*"`)

var datetimeItem = regexp.MustCompile(`\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}`)

var noteNameItem = regexp.MustCompile(`>.*<`)

var nodeDetailItem = regexp.MustCompile(`(?s)<div class="entry item">.*<br />\s*<span class="forbidden">`)

func getNotes(content string) []Note {
	matches := noteIndexItem.FindAllStringSubmatch(content, 100)
	var notes = []Note{}
	for _, item := range matches {
		note := getNote(item[0])
		notes = append(notes, note)
	}
	return notes
}

func getNote(note string) Note {
	var n Note
	hrefMatches := noteHrefItem.FindAllStringSubmatch(note, 1)
	// Url
	if len(hrefMatches) > 0 {
		href := strings.Replace(hrefMatches[0][0], "href=\"/", "", 1)
		href = strings.Replace(href, "\"", "", 1)
		n.Url = baseUrl + href
	}
	// 时间
	datetimeMatches := datetimeItem.FindAllStringSubmatch(note, 1)
	if len(datetimeMatches) > 0 {
		if t, err := time.Parse("2006-01-02 15:04:05", datetimeMatches[0][0]); err == nil {
			n.CreateTime = t
		} else {
			fmt.Println(err)
		}
	}
	// 名字
	nameMatches := noteNameItem.FindAllStringSubmatch(note, 1)
	if len(nameMatches) > 0 {
		name := nameMatches[0][0]
		n.Subject = name[1 : len(name)-1]
	}
	return n
}

// 得到时间过滤线
func (d *Douban) GetTimeLine(notes []Note) time.Time {
	if len(notes) > 1 {
		return notes[1].CreateTime
	}
	return time.Now()
}

// 根据时间线过滤
func (d *Douban) FilterNotes(notes []Note, timeline time.Time) []Note {
	var result = []Note{}
	for _, note := range notes {
		if timeline.Before(note.CreateTime) {
			result = append(result, note)
		}
	}
	return result
}

func unescaped(x string) interface{} {
	return template.HTML(x)
}

func (d *Douban) ConvertEmailBody(note Note) string {
	//temp := template.New("")
	//temp = temp.Funcs(template.FuncMap{"unescaped": unescaped})
	tmpl, _ := template.ParseFiles("notification/email/note_item.tpl")
	actorMap := make(map[string]string)
	actorMap["Datetime"] = note.CreateTime.Format("2006-01-02 15:04:05")
	actorMap["Url"] = note.Url
	actorMap["Subject"] = note.Subject
	actorMap["Content"] = note.Content

	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)

	tmpl.Execute(bw, actorMap)
	bw.Flush()
	return html.UnescapeString(b.String())
}

func (d *Douban) ConvertEmailContent(body string) string {
	//temp := template.New("")
	//temp = temp.Funcs(template.FuncMap{"unescaped": unescaped})
	t, _ := template.ParseFiles("notification/email/email_template.tpl")
	actorMap := make(map[string]string)
	actorMap["Body"] = body

	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)

	t.Execute(bw, actorMap)
	bw.Flush()
	return html.UnescapeString(b.String())
}

// 解析详细界面内容
func (d *Douban) GetNoteDetail(url string) string {
	content, _ := getContent(url)
	matches := nodeDetailItem.FindAllStringSubmatch(content, 1)
	for _, info := range matches {
		c := info[0]
		c = strings.Replace(c, `<span class="forbidden">`, "", 1)
		c = strings.Replace(c, `<div class="entry item">`, "", 1)
		c = strings.Replace(c, `<span class="info">`, "", 1)
		c = strings.Trim(c, "\n")
		c = strings.Replace(c, `<a href="`, "", 1)
		index := strings.Index(c, "</span>")
		return c[index+7:]
	}
	return ""
}
