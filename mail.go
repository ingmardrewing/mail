package mail

import (
	"fmt"
	"net/smtp"
)

type Mail struct {
	user     string
	pass     string
	from     string
	to       string
	subject  string
	mailbody string
	server   string
	port     string
}

func NewMail(user, pass, server, port, from, to, subject, mailbody string) *Mail {
	m := new(Mail)
	m.user = user
	m.pass = pass
	m.server = server
	m.port = port
	m.from = from
	m.to = to
	m.subject = subject
	m.mailbody = mailbody
	return m
}

func (m *Mail) Send() error {

	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		m.user,
		m.pass,
		m.server)

	msg := m.getMsg()

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	return smtp.SendMail(
		m.server+":"+m.port,
		auth,
		m.from,
		[]string{m.to},
		[]byte(msg),
	)
}

func (m *Mail) getMsg() string {
	tpl := `From: %s
To: %s
Content-Type: text/plain; charset="utf-8"
Subject: %s

%s
`
	return fmt.Sprintf(tpl, m.from, m.to, m.subject, m.mailbody)
}
