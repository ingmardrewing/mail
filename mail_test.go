package mail

import "testing"

func TestGetMsg(t *testing.T) {
	m := NewMail("", "", "", "", "", "", "", "")
	m.from = "test@example.com"
	m.to = "test2@example.com"
	m.subject = "test subject"
	m.mailbody = "test content"

	expected := `From: test@example.com
To: test2@example.com
Content-Type: text/plain; charset="utf-8"
Subject: test subject

test content
`

	actual := m.getMsg()

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}
