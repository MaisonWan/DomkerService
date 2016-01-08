package email

import (
	"DomkerService/notification/config"
	"net/smtp"
	"strconv"
	"strings"
)

type Mail struct {
	username string
	password string
	host     string
	port     int
	writer   string
}

func (m *Mail) Init(c *config.Config) {
	m.username = c.GetEmailUsername()
	m.password = c.GetEmailPassword()
	m.host = c.GetEmailHost()
	m.writer = "盗梦客"
	m.port = 25
}

// 发送邮件
func (m *Mail) Send(to, subject, body, mailtype string) error {
	auth := smtp.PlainAuth("", m.username, m.password, m.host)
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + mailtype + "; charset=UTF-8"
	}

	msg := []byte("To:" + to + "\r\nFrom: \"" + m.writer + "\" <" + m.username + ">\r\nSubject:" + subject +
		"\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	send_host := m.host + ":" + strconv.Itoa(m.port)
	err := smtp.SendMail(send_host, auth, m.username, send_to, msg)
	return err
}
